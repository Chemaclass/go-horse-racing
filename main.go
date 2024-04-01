package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/Chemaclass/go-horse-racing/src"
)

const maxSleepDelay = 300

func main() {
	defer func() { ShowCursor() }()
	SetUpBoard()

	const linesCount, lineLength = 12, 30

	board := NewRaceBoard(linesCount, lineLength)
	RenderGame(board)

	winnerChan := make(chan Horse)
	for line := range board {
		// each horse will be moved in different processes
		go startHorseRunning(board, line, winnerChan)
	}

	winner := <-winnerChan // wait until one horse reaches the end
	// render one last time to ensure the latest board state
	RenderRaceBoard(board, &winner)
	renderWinner(winner)
}

func startHorseRunning(board [][]*Horse, line int, winnerChan chan Horse) {
	for {
		select {
		case <-winnerChan: // check if another horse finished in
			return // such a case, then stop the for loop
		default:
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(maxSleepDelay)))
			moveHorseOnePos(board, line, winnerChan)
		}
	}
}

func moveHorseOnePos(board [][]*Horse, line int, winnerChan chan Horse) {
	cols := len(board[line])
	for col := cols - 1; col > 0; col-- {
		if board[line][col-1] == nil {
			continue
		}
		// here we identify that there is a horse in
		// the following position, so we move it to the
		// current pos, and we set nil in the other one
		board[line][col] = board[line][col-1]
		board[line][col].Position++
		board[line][col-1] = nil
		// try to declare a winner
		if board[line][col].Position+1 == cols {
			winnerChan <- *board[line][col]
		}
		break
	}
}

func renderWinner(h Horse) {
	fmt.Println("Race finished!")
	fmt.Printf("# Winner: %s\n", h)
}
