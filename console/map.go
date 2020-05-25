package console

import (
	"the_brink/party"
	"the_brink/world"

	"golang.org/x/mobile/geom"
	"github.com/fatih/color"
)


var trim string = "^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^\n"

func NewMapConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 5)
	actions[0] = "Up"
	actions[1] = "Down"
	actions[2] = "Left"
	actions[3] = "Right"
	actions[4] = "Exit"
	

	console.Actions = actions
	return console
}


func DisplayMapConsole(world *world.World, playerParty *party.Party) {
	color.Green("%s", trim)

	terminal := ""

	// Generate Map fields
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

	console := NewMapConsole()

	// Main Menu Loop
	menuLoop:
	for {
		option := console.ChooseAction()

		if option > 0 && option <= len(console.Actions) {
			color.Green("You have chosen option %d, %s", option, console.Actions[option-1])

			switch console.Actions[option-1] {
			// Exit
			case "Up":
				break menuLoop

			// Exit
			case "Down":
				break menuLoop

			// Exit
			case "Left":
				break menuLoop

			// Exit
			case "Right":
				break menuLoop

			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}