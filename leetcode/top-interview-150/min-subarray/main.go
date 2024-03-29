package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 && target != 0 {
		return 0
	}
	beginningWindow := 0
	endWindow := 0
	minLen := len(nums)
	found := false
	sum := nums[beginningWindow]

	// we scan all the array until the window is completely at the end of the array
	for beginningWindow <= len(nums)-1 && endWindow <= len(nums)-1 {
		// if the current window is > target then we try to reduce the window, so we should increment just the beginwindow pointer
		if sum > target {
			if currLen := endWindow - beginningWindow + 1; currLen <= minLen {
				minLen = currLen
				found = true
			}
			sum -= nums[beginningWindow]
			beginningWindow++
			// else if the window is == target we already are at the min window for that iteration, so we should increment both pointers at the same time
		} else if sum == target {
			if currLen := endWindow - beginningWindow + 1; currLen <= minLen {
				minLen = currLen
				found = true
			}
			sum -= nums[beginningWindow]
			beginningWindow++
			endWindow++
			if endWindow < len(nums) {
				sum += nums[endWindow]
			}
			// else if window is < target we should make the window bigger by incrementing the endWindow
		} else {
			endWindow++
			if endWindow < len(nums) {
				sum += nums[endWindow]
			}
		}
	}
	if !found {
		return 0
	} else {
		return minLen
	}

}

func main() {
	subarr := minSubArrayLen(15, []int{2, 14})
	fmt.Println(subarr)
}
