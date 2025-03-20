package server

import (
	"github.com/xykadra/go-sudoku/api"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/generate-solve", api.GenerateSolve).Methods("GET")
	router.HandleFunc("/generate", api.Generate).Methods("GET")
	router.HandleFunc("/solve", api.Solve).Methods("POST")

	return router
}
