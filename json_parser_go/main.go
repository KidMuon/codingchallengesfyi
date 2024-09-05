package main

import (
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

//object
//key
//string
//boolean
//number
//null
//array
