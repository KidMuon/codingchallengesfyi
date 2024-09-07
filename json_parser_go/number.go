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

	potential_number_string := strings.Split(contents, ",")[0]
	potential_number_string = strings.Split(potential_number_string, "]")[0]
	potential_number_string = strings.Split(potential_number_string, "}")[0]
	potential_number_string = skip_whitespace(potential_number_string)

	_, err = strconv.ParseFloat(potential_number_string, 64)
	if err != nil {
		return contents, errors.New("invalid number")
	}

	if strings.Contains(potential_number_string, "0x") {
		return contents, errors.New("hex numbers not supported")
	}

	if strings.Contains(potential_number_string, "0b") {
		return contents, errors.New("binary numbers not supported")
	}

	if strings.Contains(potential_number_string, "0o") {
		return contents, errors.New("octal numbers not supported")
	}

	if has_leading_zero(potential_number_string) {
		return contents, errors.New("numbers cannot have leading zeros")
	}

	contents, err = skip_next_expected(contents, potential_number_string)
	if err != nil {
		return contents, errors.New("invalid number")
	}

	return contents, nil
}

func has_leading_zero(number_string string) bool {
	if len(number_string) == 1 {
		return false
	}

	if strings.Contains(number_string, "0") {
		if strings.Index(number_string, "0") == strings.Index(number_string, "0.") {
			return false
		}
		if strings.Index(number_string, "0") == strings.Index(number_string, "0e") {
			return false
		}
		if strings.Index(number_string, "0") == strings.Index(number_string, "0E") {
			return false
		}
		before, _, _ := strings.Cut(number_string, "0")
		if strings.ContainsAny(before, "123456789.") {
			return false
		}
		return true
	}

	return false
}
