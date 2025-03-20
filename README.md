# Go Sudoku Project

This project is a full-stack solution for solving Sudoku puzzles, implemented in Go. The application provides three key functions for generating and solving Sudoku boards via an API. It is designed to generate new puzzles, solve existing ones, and provide solutions via a web interface.

## Live Website

The website is live and accessible at: [https://sudokuapi.tech](https://sudokuapi.tech)  
Visit the site for more information or to test the Sudoku solver.

## Key Features

### 1. **Generate**
The `generate` function creates a new, random Sudoku puzzle. The puzzle is returned as a grid, which is partially filled with numbers, while the remaining cells are left empty. This function ensures that the generated Sudoku board follows standard Sudoku rules.

### 2. **Generate-Solve**
The `generate-solve` function not only generates a random Sudoku puzzle, but also solves it immediately. This provides both the puzzle and its solution in a single API call, useful for testing or demonstration purposes.

### 3. **Solve**
The `solve` function takes a Sudoku puzzle as input and attempts to solve it. If the puzzle is solvable, it returns the solved grid. If the puzzle cannot be solved, it returns an error indicating that the solution is invalid.

## API Endpoints

### 1. **GET /generate**
- Generates a new Sudoku puzzle.
- **Request Body**: No input required.
- **Response**: A new Sudoku puzzle in the form of a 9x9 grid.

### 2. **GET /generate-solve**
- Generates a Sudoku puzzle and provides the solved grid.
- **Request Body**: No input required.
- **Response**: A Sudoku puzzle with its solved solution in the form of a 9x9 grid.

### 3. **POST /solve**
- Solves a Sudoku puzzle provided in the request body.
- **Request Body**: A Sudoku puzzle in the form of a 9x9 grid (empty cells are represented by 0).
- **Response**: The solved Sudoku puzzle in the form of a 9x9 grid, or an error message if the puzzle is unsolvable.

## How to Use

You can interact with the API via HTTP requests. The following is an example of how to send a POST request to solve a Sudoku puzzle using `curl`:

### Example Request to Solve a Puzzle:

```bash
curl -X POST https://sudokuapi.tech/solve \
  -H "Content-Type: application/json" \
  -d '{
        "board": [
            5, 3, 0, 0, 7, 0, 0, 0, 0,
            6, 0, 0, 1, 9, 5, 0, 0, 0,
            0, 9, 8, 0, 0, 0, 0, 6, 0,
            8, 0, 0, 0, 6, 0, 0, 0, 3,
            4, 0, 0, 8, 0, 3, 0, 0, 1,
            7, 0, 0, 0, 2, 0, 0, 0, 6,
            0, 6, 0, 0, 0, 0, 2, 8, 0,
            0, 0, 0, 4, 1, 9, 0, 0, 5,
            0, 0, 0, 0, 8, 0, 0, 7, 9
        ]
      }'
```

# Setup and Installation

To run the project locally:

Clone the repository:
```bash 
git clone https://github.com/xykadra/go-sudoku.git
```

Install Go: Ensure that you have Go installed on your system. 
Run the application:
Navigate to the project directory and run the Go application:
```bash
cd go-sudoku
go run cmd/main.go
```

The API will be running locally on http://localhost:8080.
Access the API:
You can test the endpoints using `curl` or `Postman`.
Visit the live website at https://sudokuapi.tech for further details.

# Technologies Used

**Go**: The backend of the application is written in Go, providing a fast and efficient server for solving Sudoku puzzles.
**Next.js**: The web application is deployed and hosted for public access.

## Deployment

This application is deployed on a virtual machine (VM) and managed using a system service. The service ensures that the Go Sudoku application runs continuously, automatically restarting the app if it crashes or the VM reboots.

The deployment setup includes:

VM: The application is hosted on a virtual machine.
System Service: A systemd service is used to manage the app, ensuring that it starts automatically on boot and stays running in the background.


Enjoy solving Sudoku with this powerful API! 🎉
