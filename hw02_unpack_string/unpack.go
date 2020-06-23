package hw02_unpack_string // nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const escapeChar = '\\'

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedStr string) (string, error) {
	var err error
	var prevChar rune
	multiplier, multiplierIsSet, escapeCharActive := 1, false, false
	buf := strings.Builder{}

	for i, char := range packedStr {
		switch {
		case char == escapeChar && !escapeCharActive:
			escapeCharActive = true
		case unicode.IsDigit(char) && !escapeCharActive:
			if i == 0 || multiplierIsSet {
				return "", ErrInvalidString
			}

			multiplier, err = strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}

			multiplierIsSet, escapeCharActive = true, false
		default:
			writeToBuf(&buf, prevChar, multiplier)

			prevChar = char
			multiplier, multiplierIsSet, escapeCharActive = 1, false, false
		}
	}

	writeToBuf(&buf, prevChar, multiplier)

	return buf.String(), nil
}

func writeToBuf(buf *strings.Builder, char rune, multiplier int) {
	if char > 0 {
		buf.WriteString(strings.Repeat(string(char), multiplier))
	}
}
