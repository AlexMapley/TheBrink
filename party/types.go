package party


import (
	"the_brink/characters"
)

// Party hold all characters
// as well as coordinates any other info
type Party struct {
	X int
	Y int
	Members []*characters.Character
	Rune rune
}
