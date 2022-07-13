package hw09structvalidator

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserGroup string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string    `json:"id" validate:"len:15"`
		Group  UserGroup `validate:"in:admin,user,guest"`
		Name   string
		Age    int      `validate:"min:14|max:100"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Phones []string `validate:"len:11"`
	}

	App struct {
		Version string `validate:"len:7"`
	}

	AppNoValCheck struct {
		Version string `validate:"errCheck:7"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404"`
		Body string `json:"omitempty"`
	}

	TestStringComplex struct {
		StringSlice []string `validate:"len:4,7|in:43rf,regeerf"`
	}

	TestIntComplex struct {
		IntSlice []int `validate:"min:10|max:100"`
	}
)

//nolint:funlen
func TestValidate(t *testing.T) {
	tests := []struct {
		name         string
		in           interface{}
		expectedErrs []error
		thisIsWrong  []string
		thisIsRight  []string
	}{
		{
			name:         "not struct",
			in:           "string, not struct",
			expectedErrs: []error{ErrNotStruct},
		},
		{
			name:         "str len is right",
			in:           App{Version: "gdrhscg"},
			expectedErrs: nil,
			thisIsWrong:  []string{},
			thisIsRight:  []string{"gdrhscg"},
		},
		{
			name:         "str len is not right",
			in:           App{Version: "3de"},
			expectedErrs: []error{ErrStrLen},
			thisIsWrong:  []string{"3de"},
			thisIsRight:  []string{},
		},
		{
			name:         "tag is not correct",
			in:           AppNoValCheck{Version: "blabla"},
			expectedErrs: []error{ErrInvalidTag},
		},
		{
			name:         "str slice with incorrect fields",
			in:           TestStringComplex{StringSlice: []string{"regeerf", "43rf", "grte43"}},
			expectedErrs: []error{ErrStrLen, ErrIn},
			thisIsWrong:  []string{"grte43"},
			thisIsRight:  []string{"43rf", "regeerf"},
		},
		{
			name: "complex checking (negative)",
			in: User{
				ID:     "123456789012345",
				Group:  "unknown",
				Name:   "petya",
				Age:    2,
				Email:  "lomov.ru",
				Phones: []string{},
			},
			expectedErrs: []error{
				ErrIntMin,
				ErrStrRegexp,
				ErrIn,
				ErrStrLen,
			},
			thisIsWrong: []string{"lomov.ru", "unknown", "2", "[]"},
			thisIsRight: []string{"123456789012345"},
		},
		{
			name: "complex checking (negative, slice is not empty)",
			in: User{
				ID:     "123456789012345",
				Group:  "unknown",
				Name:   "petya",
				Age:    2,
				Email:  "lomov.ru",
				Phones: []string{"12"},
			},
			expectedErrs: []error{
				ErrIntMin,
				ErrStrRegexp,
				ErrIn,
				ErrStrLen,
			},
			thisIsWrong: []string{"lomov.ru", "unknown", "2", "12"},
			thisIsRight: []string{"123456789012345"},
		},
		{
			name: "complex checking (positive)",
			in: User{
				ID:     "123456789012345",
				Group:  "admin",
				Name:   "",
				Age:    20,
				Email:  "lomov@mail.ru",
				Phones: []string{"89109236732"},
			},
			expectedErrs: nil,
		},
		{
			name: "int is included",
			in: Response{
				Code: 200,
				Body: "",
			},
			expectedErrs: nil,
			thisIsWrong:  []string{},
			thisIsRight:  []string{"200"},
		},
		{
			name: "int is not included",
			in: Response{
				Code: 503,
				Body: "",
			},
			expectedErrs: []error{ErrIn},
			thisIsWrong:  []string{"503"},
			thisIsRight:  []string{},
		},
		{
			name: "int slice with incorrect fields",
			in: TestIntComplex{
				IntSlice: []int{12, 0, 1, 22, 61, 431, 100, 32, -1},
			},
			expectedErrs: []error{ErrIntMin, ErrIntMin, ErrIntMin, ErrIntMax},
			thisIsWrong:  []string{"0", "1", "431", "-1"},
			thisIsRight:  []string{"12", "22", "61", "100", "32"},
		},
		{
			name: "without tags",
			in: Token{
				Header:    []byte{'d', 's', 'e', 'f'},
				Payload:   []byte{'f', 'w'},
				Signature: []byte{'w', 'd', 'e', 'q'},
			},
			expectedErrs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("subtest: %v", tt.name), func(t *testing.T) {
			tt := tt
			t.Parallel()
			errFromValidate := Validate(tt.in)
			if tt.expectedErrs == nil {
				require.NoErrorf(t, errFromValidate, "need: no error, got: ", errFromValidate)
			}
			var pValidationErrors ValidationErrors
			isValidationErrors := errors.As(errFromValidate, &pValidationErrors)
			for _, oneExpErr := range tt.expectedErrs {
				switch {
				case errors.Is(oneExpErr, ErrNotStruct):
					require.ErrorIs(t, errFromValidate, ErrNotStruct)
				case errors.Is(oneExpErr, ErrInvalidTag):
					require.ErrorIs(t, errFromValidate, ErrInvalidTag)
				case isValidationErrors:
					pExpectedErr := oneExpErr
					require.ErrorAs(t, errFromValidate, &pExpectedErr)
					require.Equal(t, len(tt.expectedErrs), len(pValidationErrors))
				}
			}
			for _, v := range tt.thisIsWrong {
				require.True(t, strings.Contains(pValidationErrors.wrongFields(), v))
			}
			for _, v := range tt.thisIsRight {
				require.False(t, strings.Contains(pValidationErrors.wrongFields(), v))
			}
		})
	}
}
