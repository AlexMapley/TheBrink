package console

import (
	"the_brink/characters"
	
	"github.com/fatih/color"
)


func NewCharacterConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 6)
	actions[0] = "Stats"
	actions[1] = "Skills"
	actions[2] = "Inventory"
	actions[3] = "Level Up"
	actions[5] = "Class"
	actions[4] = "Exit"

	console.Actions = actions

	return console
}


func DisplayCharacterConsole(character *characters.Character) {

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
			
			// Stats
				case "Skills":
					character.DisplaySkills()
			
			// Inventory
			case "Inventory":
				character.Inventory.Display()
			
			// Level Up
			case "Level Up":
				// level up player and bandit
				success := character.LevelUp()
				if success {
					color.Yellow("\n\n%sYou have levelled up\n%s\n\n", trim, trim)
					character.Rest()
				} else {
					color.Red("\n\n%sNot enough xp to level up\n%s\n\n", trim, trim)
				}


			// Become Paladin
			case "Class":
				DisplayClassConsole(character)

			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}