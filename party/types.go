package party


import (
	"the_brink/characters"
)

type Party struct {
	XCoordinate int
	YCoordinate int
	Characters []characters.Character
	Rune rune
}
