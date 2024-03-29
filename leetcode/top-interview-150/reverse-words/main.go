package main

import "fmt"

func reverseWords(s string) string {
	beginCurrentWord := len(s) - 1
	reversedString := ""

	if len(s) == 0 {
		return ""
	}

	// skip trailing spaces
	for i := beginCurrentWord; i >= 0; i-- { // O()
		if s[i] != ' ' {
			break
		}
		beginCurrentWord--
	}
	endCurrentWord := beginCurrentWord

	// move begincurrentWord until we find a space or we reach the beginning of the string
	for i := beginCurrentWord; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		beginCurrentWord--
	}

	// copy all the characters in the new string
	for i := beginCurrentWord + 1; i <= endCurrentWord; i++ {
		reversedString += string(s[i])
	}

	endCurrentWord = beginCurrentWord

	reversedSubstr := reverseWords(s[:endCurrentWord+1])

	if reversedSubstr == "" {
		return reversedString
	} else {
		return reversedString + " " + reversedSubstr
	}
}

func main() {
	rev := reverseWords("  hello world  ")
	fmt.Println(rev)
}
