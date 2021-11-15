package world

import (
	"the_brink/party"
)

// Tile represents a point on the map
type Tile struct {
	X int
	Y int
}

// World holds all info for the map
type World struct {
	XMax    int
	YMax    int
	Tiles   map[Tile]int
	Parties []*party.Party
}
