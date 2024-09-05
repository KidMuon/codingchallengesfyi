package main

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"unicode"
	"unicode/utf8"
)

func main() {
	//read in the arguments from the command line
	//assume the first one is a flag and the second is a file
	//we can build in testing the arguments later

	cmdArgs := os.Args

	//parse the flag

	flags, files, err := parseFlagsAndFiles(cmdArgs[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	//pass that option struct to a function that returns the function
	//function will later order the functions that need to be run
	//to have a defined and consistent order of output.

	//run those functions against a file from fileArg
	//open the file
	//
	byteBuffer := make([]byte, 1024)
	if len(files) > 0 {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "file error: %v\n", err)
				os.Exit(1)
			}

			output := toolOutput{
				flagOutputs: make([]int, len(flags)),
				filename:    file,
			}
			for {
				byteRead, err := f.Read(byteBuffer)
				if err != nil && err != io.EOF {
					fmt.Fprintf(os.Stderr, "file error: %v\n", err)
					os.Exit(1)
				}
				readBuffer := byteBuffer[:byteRead]
				for i, flag := range flags {
					output.flagOutputs[i] += flag.process(readBuffer)
				}
				if err == io.EOF {
					break
				}
			}
			output.print()
		}
	} else {
		output := toolOutput{
			flagOutputs: make([]int, len(flags)),
			filename:    "",
		}

		for {
			byteRead, err := os.Stdin.Read(byteBuffer)
			if err != nil && err != io.EOF {
				fmt.Fprintf(os.Stderr, "stdin error: %v\n", err)
				os.Exit(1)
			}
			readBuffer := byteBuffer[:byteRead]
			for i, flag := range flags {
				output.flagOutputs[i] += flag.process(readBuffer)
			}
			if err == io.EOF {
				break
			}
		}

		output.print()
	}
	os.Exit(0)
}

type toolOutput struct {
	flagOutputs []int
	filename    string
}

func (output toolOutput) print() {
	outputString := ""
	for _, flagValue := range output.flagOutputs {
		outputString += fmt.Sprintf("%d ", flagValue)
	}
	outputString += fmt.Sprintf("%s", output.filename)
	fmt.Println(outputString)
}

type commandLineFlag struct {
	flagLetter string
	process    func([]byte) int
	priority   int
	uniqueID   int
	value      int
}

func (c *commandLineFlag) reset() {
	c.value = 0
}

func getValidFlags() map[string]commandLineFlag {
	cw_inSpace := false
	cached_command_w := func(a []byte) int {
		count, inSpace := command_w(a, cw_inSpace)
		cw_inSpace = inSpace
		return count
	}

	cm_leftovers := []byte{}
	command_m_with_leftovers := func(a []byte) int {
		count, leftovers := command_m(a, cm_leftovers)
		cm_leftovers = leftovers
		return count
	}

	validFlags := map[string]commandLineFlag{
		"c": {
			flagLetter: "c",
			process:    command_c,
			priority:   4,
			uniqueID:   3,
		},
		"l": {
			flagLetter: "l",
			process:    command_l,
			priority:   1,
			uniqueID:   4,
		},
		"w": {
			flagLetter: "w",
			process:    cached_command_w,
			priority:   2,
			uniqueID:   5,
		},
		"m": {
			flagLetter: "m",
			process:    command_m_with_leftovers,
			priority:   3,
			uniqueID:   6,
		},
	}
	return validFlags
}

func command_c(a []byte) int {
	return len(a)
}

func command_l(a []byte) int {
	count := 0
	for _, bite := range a {
		if bite == byte('\n') {
			count++
		}
	}
	return count
}

func command_w(a []byte, inSpaces bool) (int, bool) {
	count := 0
	s := string(a)
	for _, r := range s {
		if unicode.IsSpace(r) {
			if inSpaces {
				continue
			}
			count++
			inSpaces = true
		} else {
			inSpaces = false
		}
	}
	return count, inSpaces
}

func command_m(a []byte, leftovers []byte) (int, []byte) {
	a = append(leftovers, a...)
	new_leftovers := []byte{}

	//check if the end of a is valid
	for i := 0; i < len(a); i++ {
		r, size := utf8.DecodeLastRune(a[:len(a)-i])
		if r == utf8.RuneError && size == 1 {
			continue
		} else {
			new_leftovers = a[len(a)-i:]
			a = a[:len(a)-i]
			break
		}
	}

	return utf8.RuneCount(a), new_leftovers
}

func parseFlagsAndFiles(s []string) (flags []commandLineFlag, files []string, err error) {
	for _, arg := range s {
		if arg[0] == byte('-') {
			foundFlags, err := findFlags(arg[1:])
			if err != nil {
				return []commandLineFlag{}, []string{}, err
			}
			flags = append(flags, foundFlags...)
		} else {
			files = append(files, arg)
		}
	}
	if len(flags) == 0 {
		flags, _, _ = parseFlagsAndFiles([]string{"-clw"})
	}
	return flags, files, nil
}

func findFlags(arg string) (flags []commandLineFlag, err error) {
	if arg[0] == byte('-') {
		return findLongFlags(arg[1:])
	}

	uniqueFlags := map[int]struct{}{}
	for _, char := range arg {
		flag, ok := getValidFlags()[string(char)]
		if !ok {
			return []commandLineFlag{}, errors.New("unknown flag")
		}
		if _, ok := uniqueFlags[flag.uniqueID]; ok {
			continue
		}
		uniqueFlags[flag.uniqueID] = struct{}{}
		flags = append(flags, flag)
	}

	slices.SortFunc(flags, func(a, b commandLineFlag) int {
		return cmp.Compare(a.priority, b.priority)
	})

	return flags, nil
}

func findLongFlags(arg string) ([]commandLineFlag, error) {
	return []commandLineFlag{}, errors.New("long flags not yet implemented")
}
