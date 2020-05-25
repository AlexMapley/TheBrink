package party


import (
	"the_brink/characters"
	
	"golang.org/x/mobile/geom"
)

type Party struct {
	Coodinates geom.Point
	Characters []characters.Character
}
