package party


import (
	"the_brink/characters"
)

type Party struct {
	X int
	Y int
	Characters []characters.Character
	Rune rune
}
