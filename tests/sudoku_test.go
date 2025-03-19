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
