package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	head := 0
	tail := 0
	longestSubstring := 0

	currentChars := make(map[byte]int, len(s))
	currentSubstrLen := 0
	for head < len(s) {
		// if we found it already then it means that either it was found and removed or that we have already one in the current window
		if _, found := currentChars[s[head]]; found {
			currentChars[s[head]]++
			currentSubstrLen++
			// if we have 2 occurrencies of this in the same window
			if currentChars[s[head]] > 1 {
				// move tail forward until we find the repeated element
				foundDuplicate := false
				for tail < head && !foundDuplicate {
					if currentChars[s[tail]] == 2 {
						foundDuplicate = true
					}
					// remove the tail character from the map
					currentChars[s[tail]]--
					currentSubstrLen--
					tail++
				}
			}
		} else {
			// normal case
			currentChars[s[head]] = 1
			// save the current max
			currentSubstrLen++
		}
		if currentSubstrLen > longestSubstring {
			longestSubstring = currentSubstrLen
		}
		head++
	}
	// if we never moved the tail then the longest substring is the string itself
	if tail == 0 {
		return len(s)
	}
	return longestSubstring
}

func main() {
	n := lengthOfLongestSubstring("aab")
	fmt.Println(n)
}
