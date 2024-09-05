package main

import "errors"

func key_find(contents string) (string, error) {
	var err error

	if !test_if_next(contents, "\"") {
		return "", errors.New("invalid key")
	}
	contents, err = skip_next_expected(contents, "\"")
	if err != nil {
		return "", errors.New("invalid key")
	}
	contents, err = skip_next_expected(contents, "\"")
	if err != nil {
		return "", errors.New("invalid key")
	}

	if !test_if_next(contents, ":") {
		return "", errors.New("invalid key")
	}
	contents, err = skip_next_expected(contents, ":")
	if err != nil {
		return "", errors.New("invalid key")
	}

	contents, err = value_find(contents)
	if err != nil {
		return "", err
	}

	if test_if_next(contents, ",") {
		contents, _ = skip_next_expected(contents, ",")
		return key_find(contents)
	}

	return contents, nil
}
