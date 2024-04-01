package src

import "fmt"

var HorseNames = [20][2]string{
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
	Name string // The name of the horse
	Line int    // The competition line
}

func (h *Horse) Clone(other *Horse) {
	h.Name = other.Name
	h.Line = other.Line
}

func (h *Horse) Letter() string {
	return fmt.Sprintf("%c", h.Name[0])
}

func (h *Horse) String() string {
	return fmt.Sprintf("%s (line:%d)", h.Name, h.Line)
}
