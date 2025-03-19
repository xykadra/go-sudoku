package tests

import (
	"testing"

	"github.com/xykadra/go-sudoku/internal/sudoku"
)

func TestSolveSudoku(t *testing.T) {
	grid := sudoku.Generate()
	solved := sudoku.Solve(&grid)

	if !solved {
		t.Errorf("Sudoku should be solvable but failed")
	}
}

func TestGenerateSudoku(t *testing.T) {
	board, solution := sudoku.GenerateWithSolution()

	// Ensure that the board has empty cells (zeros)
	hasEmptyCells := false
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				hasEmptyCells = true
				break
			}
		}
	}
	if !hasEmptyCells {
		t.Errorf("Generated board should have empty cells for playing")
	}

	// Ensure that the solution is valid
	if !sudoku.Solve(&solution) {
		t.Errorf("Generated solution should be a valid Sudoku solution")
	}

	// Ensure the solution is not altered
	if !isValidSudoku(solution) {
		t.Errorf("Generated solution is not a valid Sudoku board")
	}
}

// Helper function to check if a Sudoku solution is valid
func isValidSudoku(grid sudoku.Grid) bool {
	rows := make([]map[int]bool, 9)
	cols := make([]map[int]bool, 9)
	boxes := make([]map[int]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]bool)
		cols[i] = make(map[int]bool)
		boxes[i] = make(map[int]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := grid[i][j]
			if num < 1 || num > 9 {
				return false
			}

			boxIndex := (i/3)*3 + j/3
			if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[boxIndex][num] = true
		}
	}
	return true
}
