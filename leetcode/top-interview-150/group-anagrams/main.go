package main

import (
	"fmt"
	"sort"
)

type myRune []rune

func (r myRune) Len() int {
	return len(r)
}

func (r myRune) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r myRune) Less(i, j int) bool {
	return r[i] < r[j]
}

func StringToRuneSlice(s string) []rune {
	var chars []rune
	for _, r := range s {
		chars = append(chars, r)
	}
	return chars
}

func groupAnagrams(strs []string) [][]string {
	// iterate over all the list of strings
	// we keep a map storing all the indices indexed by anagram
	// to find if we need to add a word to an anagram key we sort all the letters and we compare that with the key

	anagrams := make(map[string][]int)
	var wordList [][]string

	for index, word := range strs {
		var sortedWord myRune
		sortedWord = StringToRuneSlice(word)
		sort.Sort(sortedWord)
		anagrams[string(sortedWord)] = append(anagrams[string(sortedWord)], index)
	}

	for _, v := range anagrams {
		var anagramList []string
		for _, strsIndex := range v {
			anagramList = append(anagramList, strs[strsIndex])
		}
		wordList = append(wordList, anagramList)
	}

	return wordList

}

func main() {
	sorted := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(sorted)
}
