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


func DisplayMapConsole(gameWorld *world.World, playerParty *party.Party) {

	console := NewMapConsole()

	// Main Menu Loop
	menuLoop:
	for {
		
		// Generate Map fields
		terminal := ""
		for y := 0; y < gameWorld.YMax; y++ {
			line := []rune{}
			for x := 0; x < gameWorld.XMax; x++ {
				tile := world.Tile{
					X: x,
					Y: y,
				}
				line = append(line, gameWorld.Tiles[tile])
			}
			terminal += string(line) + "\n"
		}

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
				gameWorld.UpdateMap()

			// Down
			case "Down":
				playerParty.Move(0,-1)
				gameWorld.UpdateMap()

			// Left
			case "Left":
				playerParty.Move(-1,0)
				gameWorld.UpdateMap()

			// Righ
			case "Right":
				playerParty.Move(1,0)
				gameWorld.UpdateMap()

			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}