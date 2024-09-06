package main

import (
	"errors"
	"fmt"
	"strings"
)

func string_find(contents string) (string, error) {
	var err error
	contents, err = skip_next_expected(contents, "\"")
	if err != nil {
		return contents, errors.New("invalid value")
	}

	contents, err = process_string(contents)
	if err != nil {
		return contents, err
	}

	contents, err = skip_next_expected(contents, "\"")
	if err != nil {
		return contents, errors.New("invalid value")
	}

	return contents, nil
}

func process_string(contents string) (string, error) {
	var simplified_string string = contents

	//eliminate escaped quotes until you find a quote
	for {
		quote_index := strings.Index(simplified_string, "\"")
		if quote_index == -1 {
			return contents, errors.New("string has no closing quote")
		}

		quote_escaped := false
		for i := quote_index - 1; i >= 0; i-- {
			if simplified_string[i] == '\\' {
				quote_escaped = !quote_escaped
				continue
			}
			break
		}

		if quote_escaped {
			simplified_string = simplified_string[:quote_index-1] + simplified_string[quote_index+1:]
			continue
		}

		break
	}
	//if everything else passes send back the contents without escaped_quotes
	contents_return := simplified_string

	simplified_string, _, found := strings.Cut(simplified_string, "\"")
	if !found {
		return contents, errors.New("string has no closing quote")
	}

	illegal_string_characters := []rune{'\t', '\b', '\f', '\r', '\n'}

	for _, r := range illegal_string_characters {
		if strings.ContainsRune(simplified_string, r) {
			return contents, errors.New(fmt.Sprintf("illegal character %s found", string(r)))
		}
	}

	legal_escapes := []string{"\\\\", "\\/", "\\t", "\\b", "\\f", "\\r", "\\n", "\\u"}
	for _, s := range legal_escapes {
		simplified_string = strings.ReplaceAll(simplified_string, s, "")
	}
	if strings.Contains(simplified_string, "\\") {
		return contents, errors.New("illegal backslash escape")
	}

	return contents_return, nil
}
