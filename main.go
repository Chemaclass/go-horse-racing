package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"time"
)

const milliDelay = 350

var once sync.Once

var horseNames = [20][2]string{
	{"Alloping", "Giggles"},
	{"A-lot", "Gallop"},
	{"BoJack", "Jack"},
	{"Baroness", "Belle"},
	{"Bucksnort", "Buckaroo"},
	{"Count", "Clopperstein"},
	{"Duchess", "Whirlwind"},
	{"Lady", "Hoofers"},
	{"Gallopalot", "Gallopadore"},
	{"Hoof", "Hearted"},
	{"Marquis", "Clipclapper"},
	{"Mr.", "Trot-a-lot"},
	{"Neigh", "Sayer"},
	{"Princess", "Neight"},
	{"Professor", "Neighsley"},
	{"Sir", "Trotsworth"},
	{"Sugar", "Cube"},
	{"Whinny", "McWhinerson"},
	{"Thunder", "Hooves"},
	{"Zomby", "McStompface"},
}

type Horse struct {
	Name     string // The name of the horse
	Line     int    // The competition line
	Position int    // The position in its line
	IsWinner bool   // A flag to know if it's the winner
}

func (h *Horse) Clone(other *Horse) {
	h.Name = other.Name
	h.Line = other.Line
	h.Position = other.Position
	h.IsWinner = other.IsWinner
}

func (h Horse) Letter() string {
	return fmt.Sprintf("%c", h.Name[0])
}

func (h Horse) IsFound() bool {
	return h.Name != ""
}

func (h Horse) String() string {
	return fmt.Sprintf("%s (line:%d)", h.Name, h.Line)
}

func main() {
	defer func() { showCursor() }()
	setUpBoard()

	const lines, lineLength = 12, 30
	board := generateRaceBoard(lines, lineLength)
	renderGame(board)

	winner := Horse{}
	startRace(board, &winner)

	// render one last time to ensure the latest board state
	renderRaceBoard(board)
	renderWinner(winner)
}

func setUpBoard() {
	hideCursor()
	// Force call showCursor() after Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			showCursor()
			fmt.Sprintln(sig.String())
			os.Exit(1)
		}
	}()
}

func hideCursor() {
	fmt.Print("\x1b[?25l")
}

func showCursor() {
	fmt.Print("\x1b[?25h")
}

func generateRaceBoard(lines, lineLength int) [][]*Horse {
	board := make([][]*Horse, lines)
	for line := range board {
		board[line] = make([]*Horse, lineLength)
		board[line][0] = &Horse{
			Name:     generateName(),
			Position: 0,
			Line:     line,
			IsWinner: false,
		}
	}
	return board
}

func generateName() string {
	name := horseNames[rand.Intn(len(horseNames))][0]
	surname := horseNames[rand.Intn(len(horseNames))][1]

	return name + " " + surname
}

func renderGame(board [][]*Horse) {
	go func() {
		for {
			time.Sleep(milliDelay * time.Millisecond)
			renderRaceBoard(board)
		}
	}()
}

func renderRaceBoard(board [][]*Horse) {
	var buffer bytes.Buffer
	buffer.WriteString("\n")
	for line := range board {
		renderRaceLine(board, line, &buffer)
	}
	clearScreen()
	fmt.Println(buffer.String())
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func renderRaceLine(board [][]*Horse, line int, buffer *bytes.Buffer) {
	buffer.WriteString(fmt.Sprintf(" %.2d | ", line))
	var current Horse
	for col := range board[line] {
		renderRacePosition(board, line, col, &current, buffer)
	}
	buffer.WriteString(fmt.Sprintf("| %s", current.Name))
	if current.IsWinner {
		buffer.WriteString(" [Won!]")
	}
	buffer.WriteString("\n")
}

func renderRacePosition(
	board [][]*Horse,
	line, col int,
	current *Horse,
	buffer *bytes.Buffer,
) {
	if board[line][col] == nil {
		buffer.WriteString(" ")
		return
	}

	current.Clone(board[line][col])
	if current.IsWinner {
		removeChars(buffer, current.Position+1)
		for range board[line] {
			buffer.WriteString("-")
		}
	}
	buffer.WriteString(current.Letter())
}

func removeChars(buffer *bytes.Buffer, n int) {
	buffer.WriteString(fmt.Sprintf("\033[%dD", n))
	buffer.WriteString("\033[K")
}

func startRace(
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
	randomDuration := time.Duration(rand.Intn(milliDelay))
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

func renderWinner(h Horse) {
	fmt.Println("Race finished!")
	fmt.Printf("# Winner: %s\n", h)
}
