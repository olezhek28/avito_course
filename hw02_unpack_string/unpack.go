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
		return inStr, nil
	}

	var prevSymbol rune
	var prevCh rune

	var res strings.Builder
	for _, ch := range inStr {
		if unicode.IsDigit(ch) {
			if unicode.IsDigit(prevCh) {
				return "", ErrInvalidString
			}
			if prevCh == 0 || prevSymbol == 0 {
				return "", ErrInvalidString
			}
		} else {
			if unicode.IsDigit(prevCh) {
				count, err := strconv.Atoi(string(prevCh))
				if err != nil {
					return "", ErrInvalidString
				}

				res.WriteString(strings.Repeat(string(prevSymbol), count))
			} else if prevCh != 0 {
				res.WriteRune(prevSymbol)
			}

			prevSymbol = ch
		}

		prevCh = ch
	}

	end, err := getEndStr(inStr, prevCh, prevSymbol)
	if err != nil {
		return "", err
	}

	res.WriteString(end)
	return res.String(), nil
}

func getEndStr(inStr string, prevCh rune, prevSymbol rune) (string, error) {
	lastRune, _ := utf8.DecodeLastRuneInString(inStr)
	if lastRune == utf8.RuneError {
		return "", ErrInvalidString
	}

	if unicode.IsDigit(lastRune) {
		count, err := strconv.Atoi(string(prevCh))
		if err != nil {
			return "", ErrInvalidString
		}

		return strings.Repeat(string(prevSymbol), count), nil
	}

	return string(lastRune), nil
}
