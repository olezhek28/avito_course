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
	ErrIntMin     = errors.New("has wrong value of field, needs min=")
	ErrIntMax     = errors.New("has wrong value of field, needs max=")
)

type ValidationError struct {
	Field string
	Value interface{}
	Err   error
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("поле %v: \"%v\" %v", v.Field, v.Value, v.Err)
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
			panic("you should add more types")
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

func validate(filedName string, val reflect.Value, filedType reflect.Type, tag string) (ValidationErrors, error) {
	var validateErrs ValidationErrors
	var err error

	//nolint:exhaustive
	switch filedType.Kind() {
	case reflect.String, reflect.Int:
		validateErrs, err = validateField(filedName, val, tag)
	case reflect.Slice:
		validateErrs, err = validateSlice(filedName, val, tag)
	default:
		panic("you should add more types")
	}

	if err != nil {
		return nil, err
	}

	return validateErrs, nil
}

func validateField(filedName string, val reflect.Value, tag string) (ValidationErrors, error) {
	var validateErrs ValidationErrors
	var err error

	for _, rules := range strings.Split(tag, separatorAnd) {
		var vErr ValidationErrors
		switch {
		case val.Kind() == reflect.String && strings.HasPrefix(rules, lenValidator):
			seqOfLenChecks := strings.TrimPrefix(rules, lenValidator)
			vErr, err = validateStringForLen(filedName, val, seqOfLenChecks)
		case (val.Kind() == reflect.String || val.Kind() == reflect.Int) && strings.HasPrefix(rules, inValidator):
			seqOfInChecks := strings.TrimPrefix(rules, inValidator)
			vErr, err = validateForIn(filedName, val, seqOfInChecks)
		case val.Kind() == reflect.String && strings.HasPrefix(rules, regexpValidator):
			regexpChecks := strings.TrimPrefix(rules, regexpValidator)
			vErr, err = validateStringForRegexp(filedName, val, regexpChecks)
		case val.Kind() == reflect.Int && (strings.HasPrefix(rules, minValidator) || strings.HasPrefix(rules, maxValidator)):
			vErr, err = validateIntForMinMax(filedName, val, rules)
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

func validateSlice(filedName string, val reflect.Value, tag string) (ValidationErrors, error) {
	var validateErrs ValidationErrors
	var err error

	if val.Len() == 0 {
		validateErrs = append(validateErrs, ValidationError{
			Field: filedName,
			Value: fmt.Sprint(val.Interface()),
			Err:   ErrEmptySlice,
		})
	}

	for i := 0; i < val.Len(); i++ {
		var vErr ValidationErrors

		//nolint:exhaustive
		switch val.Index(0).Kind() {
		case reflect.String, reflect.Int:
			vErr, err = validateField(filedName, val.Index(i), tag)
		default:
			panic("you should add more types")
		}
		if err != nil {
			return nil, err
		}

		validateErrs = append(validateErrs, vErr...)
	}

	return validateErrs, nil
}

func validateStringForLen(filedName string, val reflect.Value, rules string) (ValidationErrors, error) {
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
		Field: filedName,
		Value: val.String(),
		Err:   fmt.Errorf(ErrStrLen.Error(), filedName, rules),
	})

	return validateErrs, nil
}

func validateForIn(filedName string, val reflect.Value, rules string) (ValidationErrors, error) {
	var validateErrs ValidationErrors

	for _, rule := range strings.Split(rules, separatorOr) {
		if rule == fmt.Sprint(val.Interface()) {
			return nil, nil
		}
	}

	validateErrs = append(validateErrs, ValidationError{
		Field: filedName,
		Value: fmt.Sprint(val.Interface()),
		Err:   fmt.Errorf(ErrIn.Error(), filedName, rules),
	})

	return validateErrs, nil
}

func validateStringForRegexp(filedName string, val reflect.Value, rule string) (ValidationErrors, error) {
	var validateErrs ValidationErrors

	regexp := regexp.MustCompile(rule)
	if regexp.MatchString(val.String()) {
		return nil, nil
	}

	validateErrs = append(validateErrs, ValidationError{
		Field: filedName,
		Value: val.String(),
		Err:   fmt.Errorf(ErrStrRegexp.Error(), filedName, rule),
	})

	return validateErrs, nil
}

func validateIntForMinMax(filedName string, val reflect.Value, rule string) (ValidationErrors, error) {
	var validateErrs ValidationErrors

	var needMin int
	var needMax int
	var checkIsMin bool
	var checkIsMax bool
	var maxOrMinErr error
	var err error
	switch {
	case strings.HasPrefix(rule, minValidator):
		needMin, err = strconv.Atoi(strings.TrimPrefix(rule, minValidator))
		checkIsMin = true
		maxOrMinErr = fmt.Errorf("%w%v", ErrIntMin, needMin)
	case strings.HasPrefix(rule, maxValidator):
		needMax, err = strconv.Atoi(strings.TrimPrefix(rule, maxValidator))
		checkIsMax = true
		maxOrMinErr = fmt.Errorf("%w%v", ErrIntMax, needMax)
	}
	if err != nil {
		return nil, fmt.Errorf("in conversion of string %v to int has got an error: %w", rule, err)
	}
	if !(checkIsMin && int64(needMin) <= val.Int() || checkIsMax && int64(needMax) >= val.Int()) {
		validateErrs = append(validateErrs, ValidationError{
			Field: filedName,
			Value: val.Int(),
			Err:   maxOrMinErr,
		})
	}

	return validateErrs, nil
}
