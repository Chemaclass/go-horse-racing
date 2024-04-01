package src

import "math/rand"

func NewRaceBoard(lines, lineLength int) [][]*Horse {
	board := make([][]*Horse, lines)
	for line := range board {
		board[line] = make([]*Horse, lineLength)
		board[line][0] = &Horse{
			Name: generateHorseName(),
			Line: line,
		}
	}
	return board
}

func generateHorseName() string {
	name := HorseNames[rand.Intn(len(HorseNames))][0]
	surname := HorseNames[rand.Intn(len(HorseNames))][1]

	return name + " " + surname
}
