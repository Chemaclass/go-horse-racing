package main

import "fmt"

const milliDelay = 350

func main() {
	defer func() { ShowCursor() }()
	SetUpBoard()

	const lines, lineLength = 12, 30
	board := NewRaceBoard(lines, lineLength)
	RenderGame(board)

	winner := Horse{}
	StartRace(board, &winner)

	// render one last time to ensure the latest board state
	RenderRaceBoard(board)
	renderWinner(winner)
}

func renderWinner(h Horse) {
	fmt.Println("Race finished!")
	fmt.Printf("# Winner: %s\n", h)
}
