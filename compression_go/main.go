package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Incorrect number of arguements. Expecting filename only")
		os.Exit(1)
	}
	fileContents, err := importFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Println(countOccurences(fileContents))
	os.Exit(0)
}

func importFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	res, err := io.ReadAll(f)
	if err != nil && err != io.EOF {
		return "", err
	}

	return string(res), nil
}

type occurenceCount map[string]int

func countOccurences(contents string) occurenceCount {
	occurences := make(occurenceCount)
	for _, r := range contents {
		if _, ok := occurences[string(r)]; ok {
			occurences[string(r)]++
		} else {
			occurences[string(r)] = 1
		}
	}
	return occurences
}
