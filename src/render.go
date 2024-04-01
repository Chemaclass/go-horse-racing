package src

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const renderDelay = 200

func RenderGame(board [][]*Horse) {
	for {
		time.Sleep(renderDelay * time.Millisecond)
		RenderRaceBoard(board, nil)
	}
}

func RenderRaceBoard(board [][]*Horse, winner *Horse) {
	// use a "string buffer" to save the whole board state
	// so we can later use one IO call to render it
	var buffer bytes.Buffer
	buffer.WriteString("\n")
	for line := range board {
		renderRaceLine(board, line, &buffer, winner)
	}
	clearScreen()
	fmt.Println(buffer.String())
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func renderRaceLine(
	board [][]*Horse,
	line int,
	buffer *bytes.Buffer,
	winner *Horse,
) {
	buffer.WriteString(fmt.Sprintf(" %.2d | ", line))
	var current Horse
	for col := range board[line] {
		renderRacePosition(board, line, col, &current, buffer, winner)
	}
	buffer.WriteString(fmt.Sprintf("| %s", current.Name))

	if winner != nil && current.Name == winner.Name {
		buffer.WriteString(" [Won!]")
	}
	buffer.WriteString("\n")
}

func renderRacePosition(
	board [][]*Horse,
	line, col int,
	current *Horse,
	buffer *bytes.Buffer,
	winner *Horse,
) {
	if board[line][col] == nil {
		buffer.WriteString(" ")
		return
	}

	current.Clone(board[line][col])

	if winner != nil && current.Name == winner.Name {
		removeChars(buffer, col+1)
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
