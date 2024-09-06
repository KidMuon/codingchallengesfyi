package main

import (
	"errors"
	"strconv"
	"strings"
)

func test_if_number_next(contents string) bool {
	contents = skip_whitespace(contents)
	var decimal bool = false
	var exponent bool = false

	for i, r := range contents {
		if strings.ContainsRune("0", r) {
			if decimal || rune(contents[i+1]) == '.' {
				continue
			} else if exponent || strings.ContainsRune("Ee", rune(contents[i+1])) {
				continue
			} else if i != 0 && contents[i-1] != '-' {
				continue
			} else if i == 0 && !strings.ContainsRune("0123456789Ee.", rune(contents[i+1])) {
				continue
			} else {
				return false
			}
		} else if strings.ContainsRune("123456789", r) {
			continue
		} else if strings.ContainsRune(".", r) {
			if !decimal {
				decimal = true
			} else {
				return false
			}
		} else if strings.ContainsRune("Ee", r) {
			if !exponent {
				exponent = true
			} else {
				return false
			}
		} else if strings.ContainsRune("-", r) {
			if i == 0 || strings.ContainsRune("Ee", rune(contents[i-1])) {
				continue
			} else {
				return false
			}
		} else if strings.ContainsRune("+", r) {
			if i != 0 && strings.ContainsRune("Ee", rune(contents[i-1])) {
				continue
			} else {
				return false
			}
		} else {
			if i == 0 {
				return false
			}
			break
		}
	}

	return true
}

func number_find(contents string) (string, error) {
	var err error
	contents = skip_whitespace(contents)

	var number_in_value string
	for _, r := range contents {
		if strings.ContainsRune("+-.Ee0123456789", r) {
			number_in_value += string(r)
		} else {
			break
		}
	}

	contents, err = skip_next_expected(contents, number_in_value)
	if err != nil {
		return contents, errors.New("invalid number")
	}

	_, err = strconv.ParseFloat(number_in_value, 64)
	if err != nil {
		return contents, errors.New("invalid number")
	}

	return contents, nil
}
