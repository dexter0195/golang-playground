package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		currentElem2 := nums2[i]
		targetPos := m + i
		for j := m + i - 1; j >= 0; j-- {
			currentElem1 := nums1[j]
			if currentElem2 < currentElem1 {
				nums1[j+1] = currentElem1
				targetPos = j
			}
		}
		nums1[targetPos] = currentElem2
	}
	fmt.Println(nums1)
}

func main() {
	a1 := make([]int, 6)
	a1[0] = 1
	a1[1] = 2
	a1[2] = 3
	a2 := []int{
		2,
		5,
		6,
	}
	merge(a1, 3, a2, 3)
}
