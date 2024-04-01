package main

import (
	"os"
	"testing"

	. "github.com/Chemaclass/go-horse-racing/src"
)

var horse *Horse

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	horse = &Horse{
		Name: "Spirit",
		Line: 2,
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
