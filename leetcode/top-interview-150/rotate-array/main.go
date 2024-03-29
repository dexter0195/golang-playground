package main

import "fmt"

// this solution works but it's not accepted because it takes too much time for long inputs
// func rotateArray(nums []int, rotatingFactor int) []int {
// 	for r := 0; r <= rotatingFactor; r++ {
// 		// rotate the array
// 		last := nums[len(nums)-1]
// 		// rotate all the elements in the middle
// 		for i := len(nums) - 1; i > 0; i-- {
// 			nums[i] = nums[i-1]
// 		}
// 		nums[0] = last
// 	}
// 	return nums
// }

func rotateArray(nums []int, k int) []int {
	numsRotated := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newPosition := (i + k) % len(nums)
		numsRotated[newPosition] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = numsRotated[i]
	}
	return nums
}

func main() {
	fmt.Println(rotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}

// 1 2 3 4 5 6 7 8  k=3
//
