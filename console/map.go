package console

import (
	"the_brink/party"
	"the_brink/world"

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

	// Generate Map fields
	terminal := ""
	for y := 0; y < world.YMax; y++ {
		line := []rune{}
		for x := 0; x < world.XMax; x++ {
			tile := world.Tiles{
				X: x,
				Y: y,
			}
			line = append(line, world.Tiles[tile])
		}
		terminal += string(line) + "\n"
	}

	console := NewMapConsole()

	// Main Menu Loop
	menuLoop:
	for {
		
		// Diplay
		color.Green("%s", trim)
		color.Cyan("%s", terminal)
		color.Green("%s", trim)

		option := console.ChooseAction()

		if option > 0 && option <= len(console.Actions) {
			color.Green("You have chosen option %d, %s", option, console.Actions[option-1])

			switch console.Actions[option-1] {
			// Up
			case "Up":
				playerParty.Move(0,1)
				world.UpdateParties()

			// Down
			case "Down":
				playerParty.Move(0,-1)
				world.UpdateParties()

			// Left
			case "Left":
				playerParty.Move(-1,0)
				world.UpdateParties()

			// Righ
			case "Right":
				playerParty.Move(1,0)
				world.UpdateParties()

			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}