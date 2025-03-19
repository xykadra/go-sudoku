package sudoku

func Solve(grid *Grid) bool {
	row, col := findEmpty(grid)
	if row == -1 {
		return true // No empty cells, solved!
	}

	for num := 1; num <= 9; num++ {
		if isValid(grid, row, col, num) {
			grid[row][col] = num
			if Solve(grid) {
				return true
			}
			grid[row][col] = 0 // Backtrack
		}
	}
	return false
}

func findEmpty(grid *Grid) (int, int) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if grid[r][c] == 0 {
				return r, c
			}
		}
	}
	return -1, -1
}

func isValid(grid *Grid, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}
