package api

import (
	"encoding/json"
	"net/http"

	"github.com/xykadra/go-sudoku/internal/sudoku"
)

func GenerateSudoku(w http.ResponseWriter, r *http.Request) {
	grid := sudoku.Generate() // Calls the generator logic
	response, _ := json.Marshal(grid)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func SolveSudoku(w http.ResponseWriter, r *http.Request) {
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
