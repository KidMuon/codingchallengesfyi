package main

import (
	"errors"
	"fmt"
	"io"
	"os"
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
	filecontent := getFileContent(filename)
	//fmt.Println("File content\n", filecontent)

	err := testJSONfile(filecontent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s invalid JSON: %v\n", filename, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s valid JSON\n", filename)
	os.Exit(0)
}

func getFileContent(filename string) string {
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

	return filecontent
}

func testJSONfile(filecontent string) error {
	var err error
	var remainder string

	if test_if_next(filecontent, "{") {
		remainder, err = object_find(filecontent)
	} else {
		remainder, err = array_find(filecontent)
	}
	fmt.Println(remainder)
	remainder = skip_whitespace(remainder)
	if err == nil && len(remainder) != 0 {
		err = errors.New("extra information in file")
	}
	return err
}
