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

	const lines, lineLength = 12, 30

	board := NewRaceBoard(lines, lineLength)
	go RenderGame(board)

	winnerChan := make(chan Horse)
	for line := range board {
		// each horse will be moved in different processes
		go startRunningHorseInLine(board, line, winnerChan)
	}

	winner := <-winnerChan // wait until one horse reaches the end
	// render one last time to ensure the latest board state
	RenderRaceBoard(board, &winner)

	fmt.Println("Race finished!")
	fmt.Printf("# Winner: %s\n", winner)
}

func startRunningHorseInLine(board [][]*Horse, line int, winnerChan chan Horse) {
	for {
		select {
		case <-winnerChan: // check if another horse finished
			return // in such a case, then stop the for loop
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
		// the following column, so we move it to the
		// current column, and we set `nil`` the other one
		board[line][col] = board[line][col-1]
		board[line][col-1] = nil

		if col+1 == cols {
			winnerChan <- *board[line][col]
		}
		break
	}
}
