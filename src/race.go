package src

import (
	"math/rand"
	"sync"
	"time"
)

const maxSleepDelay = 500

func StartRace(
	board [][]*Horse,
	winner *Horse,
) {
	var wg sync.WaitGroup
	// use a channel with a flag to notify when a horse won
	won := make(chan bool)

	for line := range board {
		wg.Add(1)
		// each horse will be moved in different processes
		go startHorseRunning(&wg, board, line, won, winner)
	}

	wg.Wait() // wait until one horse reaches the end
	// which is controlled by the shared channel `won`
}

func startHorseRunning(
	wg *sync.WaitGroup,
	board [][]*Horse,
	line int,
	won chan bool,
	winner *Horse,
) {
	defer wg.Done()
	for {
		select {
		case <-won: // check if another horse finished in
			return // such a case, then stop the for loop
		default:
			sleepRandomMilliseconds()
			moveHorseOnePos(board, line, won, winner)
			if winner.IsFound() {
				return
			}
		}
	}
}

func sleepRandomMilliseconds() {
	randomDuration := time.Duration(rand.Intn(maxSleepDelay))
	time.Sleep(time.Millisecond * randomDuration)
}

func moveHorseOnePos(
	board [][]*Horse,
	line int,
	won chan bool,
	winner *Horse,
) {
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
			declareWinner(won, board[line][col], winner)
		}
		break
	}
}

var once sync.Once

func declareWinner(
	won chan bool,
	actual *Horse,
	winner *Horse,
) {
	// do just once to avoid multiple winners; which is highly
	// possible, as they are running in different processes
	once.Do(func() {
		actual.IsWinner = true
		winner.Clone(actual)

		// the winner will close the channel notifying
		// all other goroutines, so they can stop their
		// loops and finalising the WaitGroup and the
		// main program can also stop.
		close(won)
	})
}
