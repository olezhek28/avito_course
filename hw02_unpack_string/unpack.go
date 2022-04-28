package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return str, nil
	}

	var lastSymbol rune
	var prevCh rune

	var res string
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			if unicode.IsDigit(prevCh) {
				return "", ErrInvalidString
			}
			if prevCh == 0 || lastSymbol == 0 {
				return "", ErrInvalidString
			}
		} else {
			if unicode.IsDigit(prevCh) {
				count, err := strconv.ParseInt(string(prevCh), 10, 64)
				if err != nil {
					return "", ErrInvalidString
				}

				res += strings.Repeat(string(lastSymbol), int(count))
			} else if prevCh != 0 {
				res += string(lastSymbol)
			}

			lastSymbol = ch
		}

		prevCh = ch
	}

	if unicode.IsDigit(rune(str[len(str)-1])) {
		count, err := strconv.ParseInt(string(prevCh), 10, 64)
		if err != nil {
			return "", ErrInvalidString
		}

		res += strings.Repeat(string(lastSymbol), int(count))
	} else {
		res += string(str[len(str)-1])
	}

	return res, nil
}
