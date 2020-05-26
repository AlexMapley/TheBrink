package console

import (
	"the_brink/characters"
	
	"github.com/fatih/color"
)


func NewCharacterConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 4)
	actions[0] = "Stats"
	actions[1] = "Inventory"
	actions[2] = "Level Up"
	actions[3] = "Exit"
	

	console.Actions = actions

	return console
}


func DisplayCharacterConsole(character characters.Character) {

	console := NewCharacterConsole()

	menuLoop:
	for {
		option := console.ChooseAction()

		if option > 0 && option <= len(console.Actions) {
			color.Green("You have chosen option %d, %s", option, console.Actions[option-1])

			switch console.Actions[option-1] {
			
			// Stats
			case "Stats":
				character.Stats.Display()
			
			// Inventory
			case "Inventory":
				character.Inventory.Display()
			
			// Level Up
			case "Level Up":
				// level up player and bandit
				character.LevelUp()
				character.Rest()
				break menuLoop

			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}