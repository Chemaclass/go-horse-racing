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
		Name: "Spirit",
		Line: 2,
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
