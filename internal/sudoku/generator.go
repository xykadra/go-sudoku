package sudoku

import (
	"math/rand"
	"time"
)

// GenerateWithSolution generates a Sudoku board with its solution
func GenerateWithSolution() (Grid, Grid) {
	var board Grid
	rand.Seed(time.Now().UnixNano())

	// Fill diagonal 3x3 boxes
	for i := 0; i < 9; i += 3 {
		fillBox(&board, i, i)
	}

	Solve(&board) // Generate a fully solved board

	// Making a copy before removing digits to keep the solution
	solution := board

	removeDigits(&board, 40) // Remove numbers to create a playable board

	return board, solution
}

func Generate() Grid {
	var grid Grid
	rand.Seed(time.Now().UnixNano())

	// Fill diagonal 3x3 boxes
	for i := 0; i < 9; i += 3 {
		fillBox(&grid, i, i)
	}

	Solve(&grid)            // Fill the board
	removeDigits(&grid, 40) // Remove some numbers to create a puzzle

	return grid
}

func fillBox(grid *Grid, row, col int) {
	nums := rand.Perm(9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid[row+i][col+j] = nums[i*3+j] + 1
		}
	}
}

func removeDigits(grid *Grid, count int) {
	for i := 0; i < count; i++ {
		row := rand.Intn(9)
		col := rand.Intn(9)
		grid[row][col] = 0
	}
}
