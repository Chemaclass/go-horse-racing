package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func RenderGame(board [][]*Horse) {
	go func() {
		for {
			time.Sleep(milliDelay * time.Millisecond)
			RenderRaceBoard(board)
		}
	}()
}

func RenderRaceBoard(board [][]*Horse) {
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
