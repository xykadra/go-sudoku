package server

import (
	"github.com/xykadra/go-sudoku/api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/generate", api.GenerateSudoku).Methods("GET")
	router.HandleFunc("/solve", api.SolveSudoku).Methods("POST")

	return router
}
