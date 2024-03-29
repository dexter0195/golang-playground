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

// isBetterJump returns true if from the current position you can reach a better reach
//
// previousJumpPos : position of the last jump;
// currentPos : the current position in the array;
// jump : the jump capability at this position;
// currentReach : the current maximum position that we can arrive to;
func isBetterJump(previousJumpPos, currentPos, jump, currentReach int) bool {
	// it is a better jump if from the current position minus the size of the jumps that we had
	return currentPos+jump > currentReach
}

// let's solve it using the Kardane's Algorithm
// imagine as if you were pusing a ball towards the array, and see if you are able to get the ball to the last element
func jump(jumps []int) int {
	numJumps := 0
	reach := 0
	maxReach := 0
	previousJumpPos := 0
	maxJumpPos := 0

	for currentPos, currentJump := range jumps {
		if reach >= len(jumps) || currentPos == len(jumps)-1 {
			break
		}
		if isBetterJump(previousJumpPos, currentPos, currentJump, reach) {
			posReach := currentPos + currentJump
			if posReach > maxReach {
				maxReach = posReach
				maxJumpPos = currentPos
			}
		}
		// if we reached the end of the the previous jump
		if currentPos == reach {
			// set the jump to the maximum jump you found until now
			reach = maxReach
			// reset the maxReach to reach
			maxReach = reach
			// save the jump position
			previousJumpPos = maxJumpPos
			// increment the jump counter
			// if currentPos > 0 {
			numJumps++
			// }
		}
	}
	return numJumps
}

func main() {
	jumps := jump([]int{1, 1, 1, 1, 1})
	// jumps := jump([]int{2, 1})
	// jumps := jump([]int{2, 3, 1, 1, 4})
	fmt.Println(jumps)
}
