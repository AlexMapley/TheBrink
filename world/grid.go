package world

import (
	"golang.org/x/mobile/geom"
)

func NewGrid(xMax int, yMax int) map[geom.Point]string {
	grid := map[geom.Point]string{}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			point := geom.Point{
				X: geom.Pt(x),
				Y: geom.Pt(y),
			}
			grid[point] = "."
		}
	}
	return grid
}