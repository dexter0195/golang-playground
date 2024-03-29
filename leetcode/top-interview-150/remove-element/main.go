package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	lastDigitIndex := len(nums) - 1
	for i := 0; i <= lastDigitIndex; i++ {
		if nums[i] == val {
			for j := i; j < lastDigitIndex; j++ {
				nums[j] = nums[j+1]
			}
			nums[lastDigitIndex] = 0
			lastDigitIndex--
			i--
		} else {
			count++
		}
	}
	fmt.Println(nums)
	return count
}

func main() {
	count := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(count)
}
