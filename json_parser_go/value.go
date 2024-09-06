package main

import (
	"errors"
)

func value_find(contents string) (string, error) {
	var err error
	err = errors.New("unknown value type")
	switch {
	case test_if_next(contents, "\""):
		contents, err = string_find(contents)
		break
	case test_if_next(contents, "true") || test_if_next(contents, "false"):
		contents, err = bool_find(contents)
		break
	case test_if_next(contents, "null"):
		contents, err = null_find(contents)
		break
	case test_if_number_next(contents):
		contents, err = number_find(contents)
		break
	case test_if_next(contents, "{"):
		contents, err = object_find(contents)
		break
	case test_if_next(contents, "["):
		contents, err = array_find(contents)
		break
	}
	if err != nil {
		return "", err
	}
	return contents, nil
}

func string_find(contents string) (string, error) {
	var err error
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

func bool_find(contents string) (string, error) {
	var err error
	if test_if_next(contents, "true") {
		contents, err = skip_next_expected(contents, "true")
		if err != nil {
			return "", errors.New("invalid boolean")
		}
		return contents, nil
	}
	contents, err = skip_next_expected(contents, "false")
	if err != nil {
		return "", errors.New("invalid boolean")
	}
	return contents, nil
}

func null_find(contents string) (string, error) {
	var err error
	contents, err = skip_next_expected(contents, "null")
	if err != nil {
		return "", errors.New("invalid null")
	}
	return contents, nil
}
