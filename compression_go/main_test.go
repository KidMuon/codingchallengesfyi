package main

import (
	"testing"
)

func Test_count_e_inTestFile1(t *testing.T) {
	fileContents, _ := importFile("testfiles/1.txt")
	counts := countOccurences(fileContents)
	if counts["e"] != 21 {
		t.Fatalf(`counts["e"] = %v, expected 21`, counts["e"])
	}
}

func Test_importFile(t *testing.T) {
	fileContents, err := importFile("testfiles/hello.txt")
	if err != nil || fileContents != "Hello Test\n" {
		t.Fatalf(`fileContents, err = "%v", %v; expected "Hello Test", nil;`, fileContents, err)
	}
}
