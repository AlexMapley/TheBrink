package world

import (
	"golang.org/x/mobile/geom"
	"github.com/fatih/color"
)


var trim string = "^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^\n"

func (world *World) Display() {
	color.Green("%s", trim)

	terminal := ""

	for y := 0; y < world.YMax; y++ {
		line := []rune{}
		for x := 0; x < world.XMax; x++ {
			point := geom.Point{
				X: geom.Pt(x),
				Y: geom.Pt(y),
			}
			line = append(line, world.Grid[point])
		}
		terminal += string(line) + "\n"
	}

	color.Cyan("%s", terminal)
	color.Green("%s", trim)
}