package api

import (
	"encoding/json"
	"net/http"

	"github.com/xykadra/go-sudoku/internal/sudoku"
)

type GoSudokuResponseWithSolution struct {
	Board    sudoku.Grid `json:"board"`
	Solution sudoku.Grid `json:"solution"`
}

func GeneratSolve(w http.ResponseWriter, r *http.Request) {
	board, solution := sudoku.GenerateWithSolution()

	response := GoSudokuResponseWithSolution{
		Board:    board,
		Solution: solution,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Generate(w http.ResponseWriter, r *http.Request) {
	grid := sudoku.Generate() // Calls the generator logic
	response, _ := json.Marshal(grid)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Solve(w http.ResponseWriter, r *http.Request) {
	var request sudoku.Grid
	json.NewDecoder(r.Body).Decode(&request)

	if sudoku.Solve(&request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"solved"}`))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"invalid solution"}`))
	}
}
