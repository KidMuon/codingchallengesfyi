package main

import (
	"errors"
)

func key_find(contents string) (string, error) {
	var err error

	if !test_if_next(contents, "\"") {
		return contents, errors.New("invalid key")
	}
	contents, err = skip_next_expected(contents, "\"")
	if err != nil {
		return contents, errors.New("invalid key")
	}
	contents, err = process_string(contents)
	if err != nil {
		return contents, err
	}
	contents, err = skip_next_expected(contents, "\"")
	if err != nil {
		return contents, errors.New("invalid key")
	}

	if !test_if_next(contents, ":") {
		return contents, errors.New("invalid key")
	}
	contents, err = skip_next_expected(contents, ":")
	if err != nil {
		return contents, errors.New("invalid key")
	}

	contents, err = value_find(contents)
	if err != nil {
		return contents, err
	}

	if test_if_next(contents, ",") {
		contents, err = skip_next_expected(contents, ",")
		if err != nil {
			return contents, err
		}
		return key_find(contents)
	}

	return contents, nil
}
