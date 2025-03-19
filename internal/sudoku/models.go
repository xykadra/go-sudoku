package sudoku

type Grid [9][9]int

type SolveRequest struct {
	Grid Grid `json:"grid"`
}
