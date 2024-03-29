package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleanedString := re.ReplaceAllString(strings.ToLower(s), "")
	if len(cleanedString) == 0 {
		return true
	}

	i := 0
	j := len(cleanedString) - 1
	for i < len(cleanedString) && j > 0 {
		a := cleanedString[i]
		b := cleanedString[j]
		if a != b {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(" "))
}
