package main

import (
	"os"
	"testing"
)

var horse *Horse

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	horse = &Horse{
		Name:     "horse name 1",
		Line:     2,
		Position: 11,
		IsWinner: true,
	}
}

func TestHorseClone(t *testing.T) {

	h2 := Horse{}

	h2.Clone(horse)

	if h2.Name != horse.Name {
		t.Error("Name should be the same")
	}

	if h2.Line != horse.Line {
		t.Error("Line should be the same")
	}

	if h2.Position != horse.Position {
		t.Error("Position should be the same")
	}

	if h2.IsWinner != horse.IsWinner {
		t.Error("IsWinner should be the same")
	}
}
