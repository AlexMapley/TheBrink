package world

import (
	"golang.org/x/mobile/geom"
)

// NewGrid will return a blank grid
func NewGrid(xMax int, yMax int) map[geom.Point]rune {
	grid := map[geom.Point]rune{}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			point := geom.Point{
				X: geom.Pt(x),
				Y: geom.Pt(y),
			}
			grid[point] = '.'
		}
	}
	return grid
}

// UpdateParties will set the position and rune for all existent parties
func (world *World) UpdateParties() {
	for _, party := range world.Parties {
		world.Grid[party.Coodinates] = party.Rune
	}
}