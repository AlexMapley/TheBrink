package party


import (
	"the_brink/characters"
)

type Party struct {
	X int
	Y int
	Members []*characters.Character
	Rune rune
}
