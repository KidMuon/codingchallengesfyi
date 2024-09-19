package main

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
