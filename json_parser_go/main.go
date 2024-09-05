package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "too many arguments passed. %d passed, only 1 expected.", len(os.Args))
		os.Exit(1)
	} else if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "no file name passed")
		os.Exit(1)
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file. %v\n", err)
		os.Exit(1)
	}

	var filebytes []byte
	byteBuffer := make([]byte, 1024)

	for {
		byteRead, err := f.Read(byteBuffer)

		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "file error: %v\n", err)
			os.Exit(1)
		}

		filebytes = append(filebytes, byteBuffer[:byteRead]...)

		if err == io.EOF {
			break
		}
	}

	filecontent := string(filebytes)

	fmt.Println("File content\n", filecontent)

	_, err = object_find(filecontent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid json: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Valid JSON!")
	os.Exit(0)
}

func skip_whitespace(contents string) string {
	skip_start_index := 0
	for _, r := range contents {
		if unicode.IsSpace(r) {
			skip_start_index++
		} else {
			break
		}
	}

	returning_contents := contents[skip_start_index:]
	return returning_contents
}

func object_find(contents string) (string, error) {
	contents = skip_whitespace(contents)
	if len(contents) == 0 || contents[0] != '{' {
		return "", errors.New("invalid object")
	}
	contents = contents[1:]
	//look for keys
	contents = skip_whitespace(contents)
	if len(contents) == 0 || contents[0] != '}' {
		return "", errors.New("invalid object")
	}
	contents = contents[1:]

	return contents, nil
}

//object
//key
//string
//boolean
//number
//null
//array
