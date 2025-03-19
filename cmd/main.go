package main

import (
	"fmt"
	"net/http"

	"github.com/xykadra/go-sudoku/config"
	"github.com/xykadra/go-sudoku/server"
)

func main() {
	router := server.NewRouter()
	port := ":" + config.ServerPort
	fmt.Println("Starting Sudoku API on port", config.ServerPort)
	http.ListenAndServe(port, router)
}
