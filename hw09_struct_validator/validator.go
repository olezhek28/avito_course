package hw09structvalidator

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	validateTag     = "validate"
	separatorAnd    = "|"
	separatorOr     = ","
	lenValidator    = "len:"
	inValidator     = "in:"
	regexpValidator = "regexp:"
	minValidator    = "min:"
	maxValidator    = "max:"
)

var (
	ErrNotStruct  = errors.New("type is not a struct")
	ErrInvalidTag = errors.New("invalid check in validate tag")

	ErrIn         = errors.New("%s does not contain: %v")
	ErrStrLen     = errors.New("%s length exceeds limit: %v")
	ErrStrRegexp  = errors.New("%s doesn't match regexp: %v")
	ErrEmptySlice = errors.New("slice is empty")
	ErrIntMin     = errors.New("%s value is less than: %v")
	ErrIntMax     = errors.New("%s value is great than: %v")
)

type ValidationError struct {
	Field string
	Value interface{}
	Err   error
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%v: \"%v\" %v", v.Field, v.Value, v.Err)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	buffer := bytes.Buffer{}
	for _, val := range v {
		buffer.WriteString(val.Error())
		buffer.WriteString("\n")
	}

	return fmt.Sprint(buffer.String())
}

func (v ValidationErrors) wrongFields() string {
	buffer := bytes.Buffer{}
	for _, val := range v {
		switch typeField := val.Value.(type) {
		case string:
			buffer.WriteString(typeField)
		case int64:
			buffer.WriteString(fmt.Sprint(typeField))
		default:
			panic("unknown type") // only tests
		}
		buffer.WriteString("\n")
	}

	return fmt.Sprint(buffer.String())
}

func Validate(v interface{}) error {
	in := reflect.ValueOf(v)
	if in.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	refType := reflect.TypeOf(v)
	var errs ValidationErrors

	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		if tag, ok := field.Tag.Lookup(validateTag); ok && tag != "" {
			validateErr, err := validate(field.Name, in.Field(i), field.Type, tag)
			if err != nil {
				return fmt.Errorf("failed to validate field %v: %w", field.Name, err)
			}

			errs = append(errs, validateErr...)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func validate(fieldName string, val reflect.Value, filedType reflect.Type, tag string) (ValidationErrors, error) {
	var validateErrs ValidationErrors
	var err error

	//nolint:exhaustive
	switch filedType.Kind() {
	case reflect.String, reflect.Int:
		validateErrs, err = validateField(fieldName, val, tag)
	case reflect.Slice:
		validateErrs, err = validateSlice(fieldName, val, tag)
	default:
		return nil, fmt.Errorf("unknown type")
	}

	if err != nil {
		return nil, err
	}

	return validateErrs, nil
}

func validateField(fieldName string, val reflect.Value, tag string) (ValidationErrors, error) {
	var validateErrs ValidationErrors
	var err error

	for _, rules := range strings.Split(tag, separatorAnd) {
		var vErr ValidationErrors
		switch {
		case val.Kind() == reflect.String && strings.HasPrefix(rules, lenValidator):
			rule := strings.TrimPrefix(rules, lenValidator)
			vErr, err = validateStringForLen(fieldName, val, rule)
		case (val.Kind() == reflect.String || val.Kind() == reflect.Int) && strings.HasPrefix(rules, inValidator):
			rule := strings.TrimPrefix(rules, inValidator)
			vErr = validateForIn(fieldName, val, rule)
		case val.Kind() == reflect.String && strings.HasPrefix(rules, regexpValidator):
			rule := strings.TrimPrefix(rules, regexpValidator)
			vErr = validateStringForRegexp(fieldName, val, rule)
		case val.Kind() == reflect.Int && (strings.HasPrefix(rules, minValidator) || strings.HasPrefix(rules, maxValidator)):
			vErr, err = validateIntForMinMax(fieldName, val, rules)
		default:
			return nil, ErrInvalidTag
		}

		validateErrs = append(validateErrs, vErr...)
	}

	if err != nil {
		return nil, err
	}

	return validateErrs, nil
}

func validateSlice(fieldName string, val reflect.Value, tag string) (ValidationErrors, error) {
	var validateErrs ValidationErrors
	var err error

	if val.Len() == 0 {
		validateErrs = append(validateErrs, ValidationError{
			Field: fieldName,
			Value: fmt.Sprint(val.Interface()),
			Err:   ErrEmptySlice,
		})
	}

	for i := 0; i < val.Len(); i++ {
		var vErr ValidationErrors

		//nolint:exhaustive
		switch val.Index(0).Kind() {
		case reflect.String, reflect.Int:
			vErr, err = validateField(fieldName, val.Index(i), tag)
		default:
			return nil, fmt.Errorf("unknown type")
		}
		if err != nil {
			return nil, err
		}

		validateErrs = append(validateErrs, vErr...)
	}

	return validateErrs, nil
}

func validateStringForLen(fieldName string, val reflect.Value, rules string) (ValidationErrors, error) {
	var validateErrs ValidationErrors

	for _, rule := range strings.Split(rules, separatorOr) {
		targetLen, err := strconv.Atoi(rule)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string %v to int: %w", rule, err)
		}
		valStr := val.String()
		if targetLen == len(valStr) {
			return nil, nil
		}
	}

	validateErrs = append(validateErrs, ValidationError{
		Field: fieldName,
		Value: val.String(),
		Err:   fmt.Errorf(ErrStrLen.Error(), fieldName, rules),
	})

	return validateErrs, nil
}

func validateForIn(fieldName string, val reflect.Value, rules string) ValidationErrors {
	var validateErrs ValidationErrors

	for _, rule := range strings.Split(rules, separatorOr) {
		if rule == fmt.Sprint(val.Interface()) {
			return nil
		}
	}

	validateErrs = append(validateErrs, ValidationError{
		Field: fieldName,
		Value: fmt.Sprint(val.Interface()),
		Err:   fmt.Errorf(ErrIn.Error(), fieldName, rules),
	})

	return validateErrs
}

func validateStringForRegexp(fieldName string, val reflect.Value, rule string) ValidationErrors {
	var validateErrs ValidationErrors

	regexp := regexp.MustCompile(rule)
	if regexp.MatchString(val.String()) {
		return nil
	}

	validateErrs = append(validateErrs, ValidationError{
		Field: fieldName,
		Value: fmt.Sprint(val.Interface()),
		Err:   fmt.Errorf(ErrStrRegexp.Error(), fieldName, rule),
	})

	return validateErrs
}

func validateIntForMinMax(fieldName string, val reflect.Value, rule string) (ValidationErrors, error) {
	var validateErrs ValidationErrors

	switch {
	case strings.HasPrefix(rule, minValidator):
		minLimit, err := strconv.Atoi(strings.TrimPrefix(rule, minValidator))
		if err != nil {
			return nil, fmt.Errorf("failed to convert string %v to int: %w", rule, err)
		}

		if val.Int() < int64(minLimit) {
			validateErrs = append(validateErrs, ValidationError{
				Field: fieldName,
				Value: fmt.Sprint(val.Interface()),
				Err:   fmt.Errorf(ErrIntMin.Error(), fieldName, minLimit),
			})
		}
	case strings.HasPrefix(rule, maxValidator):
		maxLimit, err := strconv.Atoi(strings.TrimPrefix(rule, maxValidator))
		if err != nil {
			return nil, fmt.Errorf("failed to convert string %v to int: %w", rule, err)
		}

		if val.Int() > int64(maxLimit) {
			validateErrs = append(validateErrs, ValidationError{
				Field: fieldName,
				Value: fmt.Sprint(val.Interface()),
				Err:   fmt.Errorf(ErrIntMax.Error(), fieldName, maxLimit),
			})
		}
	}

	return validateErrs, nil
}
