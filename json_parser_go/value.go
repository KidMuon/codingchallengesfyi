package main

import (
	"errors"
)

func value_find(contents string) (string, error) {
	contents, err := string_find(contents)
	if err != nil {
		return "", err
	}
	return contents, nil
}

func string_find(contents string) (string, error) {
	var err error
	if !test_if_next(contents, "\"") {
		return "", errors.New("invalid value")
	}
	contents, err = skip_next_expected(contents, "\"")
	if err != nil {
		return "", errors.New("invalid value")
	}
	contents, err = skip_next_expected(contents, "\"")
	if err != nil {
		return "", errors.New("invalid value")
	}

	return contents, nil
}
