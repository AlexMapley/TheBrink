package console

import (
	"golang.org/x/mobile/geom"
	"github.com/fatih/color"
)


var trim string = "^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^\n"

func NewMapConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 1)
	actions[0] = "Exit"
	

	console.Actions = actions
	return console
}


func Display(world *World) {
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

	console := NewCharacterConsole()

	// Main Menu Loop
	menuLoop:
	for {
		option := console.ChooseAction()

		if option > 0 && option <= len(console.Actions) {
			color.Green("You have chosen option %d, %s", option, console.Actions[option-1])

			switch console.Actions[option-1] {
			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}