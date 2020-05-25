package party


import (
	"the_brink/characters"
	"the_brink/world"
)

type Party struct {
	XCoordinate int
	YCoordinate int
	Tile world.Tile
	Characters []characters.Character
	Rune rune
}
