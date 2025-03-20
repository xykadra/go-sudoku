package sudoku

import "fmt"

// Solve tries to solve the Sudoku puzzle using backtracking
func Solve(grid *Grid) bool {
	row, col := findEmpty(grid) // Find the next empty spot on the grid
	if row == -1 {
		return true // No empty cells left, puzzle is solved
	}

	// Try numbers 1 through 9 in the empty spot
	for num := 1; num <= 9; num++ {
		if isValid(grid, row, col, num) { // Check if the number is valid for the spot
			grid[row][col] = num // Place the number

			// Recursively attempt to solve the rest of the grid
			if Solve(grid) {
				return true
			}

			// If placing num leads to an invalid state, backtrack
			grid[row][col] = 0
		}
	}

	// If no valid number can be placed, return false (unsolvable)
	return false
}

// findEmpty finds the first empty spot (represented by 0) in the grid
func findEmpty(grid *Grid) (int, int) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if grid[r][c] == 0 {
				return r, c // Return the first empty row, column
			}
		}
	}
	return -1, -1 // No empty spots, puzzle solved
}

// isValid checks if placing 'num' at grid[row][col] is valid
func isValid(grid *Grid, row, col, num int) bool {
	// Check row and column
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// Check 3x3 subgrid
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	// If no conflicts, it's a valid placement
	return true
}

// PrintGrid prints the Sudoku grid
func PrintGrid(grid *Grid) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			fmt.Printf("%d ", grid[r][c])
		}
		fmt.Println()
	}
}
