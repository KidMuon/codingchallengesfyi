package main

import (
	"errors"
	"strconv"
	"strings"
)

func test_if_number_next(contents string) bool {
	contents = skip_whitespace(contents)

	potential_number_string := strings.Split(contents, ",")[0]
	potential_number_string = strings.Split(potential_number_string, "]")[0]
	potential_number_string = strings.Split(potential_number_string, "}")[0]
	potential_number_string = skip_whitespace(potential_number_string)

	_, err := strconv.ParseFloat(potential_number_string, 64)
	if err != nil {
		return false
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
