package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	arStr := strings.Split(str, "")
	var digit, escape bool
	var digValue int
	var pastChar, newstr, printChar, prevstr string
	escapeChar := "\\"

	for i, char := range arStr {
		dig, err := isDigit(char)
		if err != nil {
			return "", ErrInvalidString
		}
		if dig {
			if digit == true || escape == true || i == 0 {
				return "", ErrInvalidString
			}
			digit = true
			digValue, err = strconv.Atoi(char)
		} else {
			digit = false
		}

		if char == escapeChar {
			if escape == true {
				return "", ErrInvalidString
			}
			escape = true
		}

		if digit {
			if digValue == 1 {
				printChar = ""
			} else if digValue == 0 {
				newstr = prevstr
				continue
			} else {
				printChar = strings.Repeat(pastChar, digValue-1)
			}
		} else if escape == true {
			if char == escapeChar {
				continue
			} else {
				printChar = escapeChar + char
				pastChar = printChar
				escape = false
			}
		} else {
			printChar = char
		}
		pastChar = char

		prevstr = newstr
		newstr = newstr + printChar
	}

	return newstr, nil
}

func isDigit(str string) (bool, error) {
	return regexp.MatchString(`\d`, str)
}
