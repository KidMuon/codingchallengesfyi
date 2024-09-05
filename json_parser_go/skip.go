package main

import (
	"errors"
	"strings"
)

func skip_whitespace(contents string) string {
	return strings.TrimSpace(contents)
}

func test_if_next(contents string, test_str string) bool {
	contents = skip_whitespace(contents)
	before, _, found := strings.Cut(contents, test_str)
	return len(before) == 0 && found
}

func skip_next_expected(contents string, expected string) (string, error) {
	_, after, found := strings.Cut(contents, expected)
	if !found {
		return "", errors.New("expected not found")
	}
	return after, nil
}
