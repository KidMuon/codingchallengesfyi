package main

import (
	"errors"
)

func array_find(contents string) (string, error) {
	var err error
	contents, err = skip_next_expected(contents, "[")
	if err != nil {
		return "", errors.New("invalid array")
	}

	if !test_if_next(contents, "]") {
		contents, err = value_find(contents)
		if err != nil {
			return "", err
		}
	}

	contents, err = skip_next_expected(contents, "]")
	if err != nil {
		return "", errors.New("invalid array")
	}

	return contents, nil
}
