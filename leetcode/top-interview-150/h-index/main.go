package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	maxH := len(citations)
	for i := 0; i < len(citations); i++ {
		currH := citations[i]
		// if we have more remaining elements than current citations  // 1, 3, 6, 11, 12
		if currH < maxH {
			// then we save the maxH and we continue
			maxH--
		} else {
			break
		}
	}
	return maxH
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
