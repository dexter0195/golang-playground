package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 1
	alreadyEncountered := false
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			nums[j+1] = nums[i]
			alreadyEncountered = false
			j++
			count++
		} else {
			if !alreadyEncountered {
				nums[j+1] = nums[i]
				alreadyEncountered = true
				j++
				count++
			} else {
				continue
			}
		}
	}
	return count
}

func main() {
	count := removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	fmt.Println(count)
}

// 0 1 2 2 2 2 3 3 3 3 4
//       j   i
// alredy t
