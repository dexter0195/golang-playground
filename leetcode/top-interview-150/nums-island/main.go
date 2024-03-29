package main

import "fmt"

func numIslands(grid [][]byte) int {
	countIslands := 0
	nRows := len(grid)
	nColumns := len(grid[0])
	visitedMatrix := make([][]bool, len(grid))
	for i := range visitedMatrix {
		visitedMatrix[i] = make([]bool, nColumns)
	}

	// visit all the elements in the matrix
	for rIndex := range grid {
		for cIndex := range grid[rIndex] {
			// if we didn't already visit and we are not on the water
			if !visitedMatrix[rIndex][cIndex] && grid[rIndex][cIndex] == '1' {
				// visit all the lands that you can from that point with DFS
				visitDFS(grid, visitedMatrix, rIndex, cIndex, nRows, nColumns)
				countIslands++
			}
		}
	}
	return countIslands
}

func visitDFS(grid [][]byte, visitedMatrix [][]bool, r, c, nRows, nColumns int) {
	// base case
	// the current element grid[r][c] is out of bound
	if r > nRows-1 || c > nColumns-1 {
		return
	}
	visitedMatrix[r][c] = true
	if r-1 >= 0 {
		// visit on the up
		if grid[r-1][c] == '1' {
			if !visitedMatrix[r-1][c] {
				visitDFS(grid, visitedMatrix, r-1, c, nRows, nColumns)
			}
		}
	}
	if r+1 <= nRows-1 {
		// visit on the bottom
		if grid[r+1][c] == '1' {
			if !visitedMatrix[r+1][c] {
				visitDFS(grid, visitedMatrix, r+1, c, nRows, nColumns)
			}
		}
	}
	if c+1 <= nColumns-1 {
		// visit on the right
		if grid[r][c+1] == '1' {
			if !visitedMatrix[r][c+1] {
				visitDFS(grid, visitedMatrix, r, c+1, nRows, nColumns)
			}
		}
	}
	if c-1 >= 0 {
		// visit on the left
		if grid[r][c-1] == '1' {
			if !visitedMatrix[r][c-1] {
				visitDFS(grid, visitedMatrix, r, c-1, nRows, nColumns)
			}
		}
	}
}

// there is also another version that is more compact but that takes longer to execute
// func visitDFS(grid [][]byte, visitedMatrix [][]bool, r, c, nRows, nColumns int) {
// 	// base case
// 	// the current element grid[r][c] is out of bound
// 	if r > nRows-1 || c > nColumns-1 || r < 0 || c < 0 || visitedMatrix[r][c] || grid[r][c]) != 1 {
// 		return
// 	}
// 	visitedMatrix[r][c] = true
// 	visitDFS(grid, visitedMatrix, r-1, c, nRows, nColumns)
// 	visitDFS(grid, visitedMatrix, r+1, c, nRows, nColumns)
// 	visitDFS(grid, visitedMatrix, r, c+1, nRows, nColumns)
// 	visitDFS(grid, visitedMatrix, r, c-1, nRows, nColumns)
// }

func main() {
	fmt.Println(numIslands([][]byte{ // broken, as the exercie was requesting bytes
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}

// ['1','0','1','1','1'],
// ['1','0','1','0','1'],
// ['1','1','1','0','1']
