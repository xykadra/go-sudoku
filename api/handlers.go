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

type SudokuRequest struct {
	Board []int `json:"board"`
}

func GenerateSolve(w http.ResponseWriter, r *http.Request) {
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
	var request SudokuRequest
	// Decode the request body into the sudoku grid
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// If decoding fails, send a bad request response
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"invalid request", "message":"Failed to decode request body"}`))
		return
	}

	// Convert the 1D array to the 2D grid
	var grid sudoku.Grid
	for i, num := range request.Board {
		row := i / 9
		col := i % 9
		grid[row][col] = num
	}

	// Try to solve the puzzle
	if sudoku.Solve(&grid) {
		// Return the solved grid as part of the response if solvable
		w.WriteHeader(http.StatusOK)
		// Encode the solved grid into the response
		response := map[string]interface{}{
			"status":   "solved",
			"solution": grid, // The solved grid will be in the 'grid' variable
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"status":"error", "message":"Failed to create JSON response"}`))
			return
		}
		w.Write(jsonResponse)
	} else {
		// Return a bad request response if the puzzle is unsolvable
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"invalid solution"}`))
	}
}
