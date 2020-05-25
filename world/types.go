package world

import (
	"golang.org/x/mobile/geom"
	"the_brink/party"
)

type MetaGame struct {
	Day int
}

type World struct {
	XMax int
	YMax int
	Grid map[geom.Point]rune
	Parties []party.Party
}