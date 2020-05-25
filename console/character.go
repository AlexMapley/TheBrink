package console


func NewCharacterConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 3)
	actions[0] = "Stats"
	actions[1] = "Inventory"
	actions[2] = "Exit"
	

	console.Actions = actions

	return console
}


func DisplayCharacterConsole(character *characters.Character) {

	console := NewCharacterConsole()

	menuLoop:
	for {
		option := console.ChooseAction()

		if option > 0 && option <= len(townConsole.Actions) {
			color.Green("You have chosen option %d, %s", option, townConsole.Actions[option-1])

			switch townConsole.Actions[option-1] {
			
			// Stats
			case "Stats":
				character.Stats.Display()
			
			// Inventory
			case "Inventory":
				character.Inventory.Display()
			}

			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}