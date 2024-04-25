package main

import (
	"fmt"
	"os"
	"os/signal"
)

func SetUpBoard() {
	hideCursor()
	// Force call showCursor() after Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			ShowCursor()
			fmt.Sprintln(sig.String())
			os.Exit(1)
		}
	}()
}

func hideCursor() {
	fmt.Print("\x1b[?25l")
}

func ShowCursor() {
	fmt.Print("\x1b[?25h")
}
