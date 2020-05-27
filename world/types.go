package world

import (
	"the_brink/party"
)

// SaveFile holds all information from a playthrough 
type SaveFile struct {
	Day int
	Party *party.Party
}

type Tile struct {
	X int
	Y int
}

type World struct {
	XMax int
	YMax int
	Tiles map[Tile]rune
	Parties []*party.Party
}