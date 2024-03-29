package main

import "fmt"

func removeDuplicates(nums []int) int {
	values := make(map[int]bool, len(nums))
	count := 0
	lastDigitIndex := len(nums) - 1
	for i := 0; i <= lastDigitIndex; i++ {
		if _, ok := values[nums[i]]; ok {
			for j := i; j < lastDigitIndex; j++ {
				nums[j] = nums[j+1]
			}
			nums[lastDigitIndex] = 0
			lastDigitIndex--
			i--
		} else {
			values[nums[i]] = true
			count++
		}
	}
	fmt.Println(nums)
	return count
}

func main() {
	count := removeDuplicates([]int{0, 1, 2, 2, 3, 0, 4, 2})
	fmt.Println(count)
}
