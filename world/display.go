package world

import (
	"golang.org/x/mobile/geom"
	"fmt"
)

func (world *World) Display() {
	terminal := ""
	for x := 0; x < world.XMax; x++ {
		line := ""
		for y := 0; y < world.YMax; y++ {
			point := geom.Point{
				X: geom.Pt(x),
				Y: geom.Pt(y),
			}
			line += world.Grid[point]
		}
		terminal += line + "\n"
	}
	fmt.Printf("\n\nTerminal: %s", terminal)
}