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
	actions[4] = "Class"
	actions[5] = "x | q | esc to exit"

	console.Actions = actions

	return console
}

func DisplayCharacterConsole(character *characters.Character) {

	console := NewCharacterConsole()

menuLoop:
	for {
		option := console.ChooseAction(nil, nil)
		
		// Escape character
		if option == -1 {
			break menuLoop
		}
		// Map move
		if option == -2 {
			continue
		}

		if option > 0 && option <= len(console.Actions) {

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
					color.Yellow("\n\n%s\nYou have leveled up\n%s\n\n", trim, trim)
					character.Rest()
				} else {
					color.Red("\n\n%s\nNot enough xp to level up\n%s\n\n", trim, trim)
				}

			// Open up class console
			case "Class":
				DisplayClassConsole(character)
			}
		}
	}
}
