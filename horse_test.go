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
		Name:     "Spirit",
		Line:     2,
		Position: 11,
		IsWinner: true,
	}
}

func TestHorseCloneName(t *testing.T) {
	h := Horse{}

	if h.Name == horse.Name {
		t.Error("Name should not be the same")
	}

	h.Clone(horse)

	if h.Name != horse.Name {
		t.Error("Name should be the same")
	}
}

func TestHorseCloneLine(t *testing.T) {
	h := Horse{}

	if h.Line == horse.Line {
		t.Error("Line should not be the same")
	}

	h.Clone(horse)

	if h.Line != horse.Line {
		t.Error("Line should be the same")
	}
}

func TestHorseClonePosition(t *testing.T) {
	h := Horse{}

	if h.Position == horse.Position {
		t.Error("Position should not be the same")
	}

	h.Clone(horse)

	if h.Position != horse.Position {
		t.Error("Position should be the same")
	}
}

func TestHorseCloneIsWinner(t *testing.T) {
	h := Horse{}

	if h.IsWinner == horse.IsWinner {
		t.Error("IsWinner should not be the same")
	}

	h.Clone(horse)

	if h.IsWinner != horse.IsWinner {
		t.Error("IsWinner should be the same")
	}
}

func TestHorseString(t *testing.T) {
	if horse.String() != "Spirit (line:2)" {
		t.Error("String() not working")
	}
}

func TestHorseLetter(t *testing.T) {
	if horse.Letter() != "S" {
		t.Error("Letter() wrong")
	}
}

func TestHorseIsFound(t *testing.T) {
	if !horse.IsFound() {
		t.Error("IsFound() should be true")
	}
}
