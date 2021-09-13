package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(t string) (string, error) {
	i := 0
	input := []rune(t)
	var Result strings.Builder
	for i < len(input) {
		symbol := input[i]
		if unicode.IsDigit(symbol) {
			return "", ErrInvalidString
		}
		i++
		if i >= len(input) || !unicode.IsDigit(input[i]) {
			Result.WriteString(string(symbol))
		} else {
			count, _ := (strconv.Atoi(string(input[i])))
			Result.WriteString(strings.Repeat(string(symbol), count))
			i++
		}
	}
	return Result.String(), nil
}
