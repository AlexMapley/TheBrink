package world

import (
	"golang.org/x/mobile/geom"
	"github.com/fatih/color"
)


var trim string = "^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^\n"

func (world *World) Display() {
	terminal := trim
	for y := 0; y < world.YMax; y++ {
		line := ""
		for x := 0; x < world.XMax; x++ {
			point := geom.Point{
				X: geom.Pt(x),
				Y: geom.Pt(y),
			}
			line += world.Grid[point]
		}
		terminal += line + "\n"
	}
	terminal += trim
	color.Cyan("%s", terminal)
}