package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inStr string) (string, error) {
	if len(inStr) == 0 {
		return "", nil
	}

	var res strings.Builder
	prevDigit := -1
	var symbol rune
	for _, ch := range inStr {
		if !unicode.IsDigit(ch) {
			if prevDigit >= 0 {
				res.WriteString(strings.Repeat(string(symbol), prevDigit))
			} else if symbol > 0 {
				res.WriteRune(symbol)
			}

			symbol = ch
			prevDigit = -1
			continue
		}

		if prevDigit >= 0 || symbol == 0 {
			return "", ErrInvalidString
		}

		var err error
		prevDigit, err = strconv.Atoi(string(ch))
		if err != nil {
			return "", ErrInvalidString
		}
	}

	last, err := getEndStr(inStr, symbol, prevDigit)
	if err != nil {
		return "", ErrInvalidString
	}

	res.WriteString(last)
	return res.String(), nil
}

func getEndStr(inStr string, symbol rune, prevDigit int) (string, error) {
	lastRune, _ := utf8.DecodeLastRuneInString(inStr)
	if lastRune == utf8.RuneError {
		return "", ErrInvalidString
	}

	if unicode.IsDigit(lastRune) {
		return strings.Repeat(string(symbol), prevDigit), nil
	}

	return string(lastRune), nil
}
