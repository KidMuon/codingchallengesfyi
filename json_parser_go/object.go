package main

import "errors"

func object_find(contents string) (string, error) {
	var err error

	contents = skip_whitespace(contents)
	if !test_if_next(contents, "{") {
		return contents, errors.New("invalid object")
	}

	contents, err = skip_next_expected(contents, "{")
	if err != nil {
		return contents, errors.New("invalid object")
	}

	if !test_if_next(contents, "}") {
		contents, err = key_find(contents)
		if err != nil {
			return contents, err
		}
	}

	contents, err = skip_next_expected(contents, "}")
	if err != nil {
		return contents, errors.New("invalid object")
	}

	return contents, nil
}
