package main

import "fmt"

// this solution uses memoization. It works however is not very efficient in terms of timing
// func isReachable(jumps []int, pos int, reachability map[int]bool) bool {
// 	if reachable, ok := reachability[pos]; ok {
// 		if reachable {
// 			return true
// 		}
// 	} else {
// 		//iterate over
// 		for i := 1; i <= jumps[pos] && pos+i <= len(jumps); i++ {
// 			if isReachable(jumps, pos+i, reachability) {
// 				reachability[pos] = true
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func canJump(jumps []int) bool {
// 	r := make(map[int]bool, len(jumps))
// 	r[len(jumps)-1] = true // we can always reach the last element from itself
// 	if len(jumps) == 1 {
// 		return true
// 	}
// 	foundPath := false
// 	for i := len(jumps) - 2; i >= 0; i-- {
// 		foundPath = isReachable(jumps, i, r)
// 	}
// 	return foundPath
// }

// let's solve it using the Kardane's Algorithm
// imagine as if you were pusing a ball towards the array, and see if you are able to get the ball to the last element
func canJump(jumps []int) bool {
	reach := 0

	for i, v := range jumps {
		if reach < i {
			return false
		}
		if i+v > reach {
			// push the ball forward
			reach += v
		}
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{2, 0, 1, 0, 0, 1}))
}
