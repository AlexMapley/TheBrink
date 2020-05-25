package world

import (
	"the_brink/party"
)

type MetaGame struct {
	Day int
}

type Tile struct {
	X int
	Y int
}

type World struct {
	XMax int
	YMax int
	Tiles map[Tile]rune
	Parties map[Tile]party.Party
}