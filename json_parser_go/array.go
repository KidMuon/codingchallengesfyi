package main

import (
	"errors"
)

func array_find(contents string) (string, error) {
	var err error
	contents, err = skip_next_expected(contents, "[")
	if err != nil {
		return contents, errors.New("invalid array")
	}

	var values_in_array, values_evaluated int

	for !test_if_next(contents, "]") {
		if test_if_next(contents, ",") {
			contents, err = skip_next_expected(contents, ",")
			if err != nil {
				return contents, errors.New("invalid array")
			}
			values_evaluated++
		}
		if values_evaluated > values_in_array {
			return contents, errors.New("missing value in array")
		}
		contents, err = value_find(contents)
		if err != nil {
			return contents, err
		}
		values_in_array++
	}

	contents, err = skip_next_expected(contents, "]")
	if err != nil {
		return contents, errors.New("invalid array")
	}

	return contents, nil
}
