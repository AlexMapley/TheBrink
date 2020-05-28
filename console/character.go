package console

import (
	"the_brink/characters"
	
	"github.com/fatih/color"
)


func NewCharacterConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 8)
	actions[0] = "Stats"
	actions[1] = "Skills"
	actions[2] = "Inventory"
	actions[3] = "Level Up"
	actions[4] = "Exit"
	actions[5] = "Become Paladin"
	actions[6] = "Become NightBlade"
	actions[7] = "Become Duelist"

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
					character.Stats.DisplaySkills()
			
			// Inventory
			case "Inventory":
				character.Inventory.Display()
			
			// Level Up
			case "Level Up":
				// level up player and bandit
				character.LevelUp()
				character.Rest()

			// Become Paladin
			case "Become Paladin":
				var accepted bool
				*character, accepted = characters.Paladin(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Paladin\n%s\n\n", trim, trim)
				} else {
					color.HiRed("\n\n%sYou cannot become a Paladin\n%s\n\n", trim, trim)
				}

			// Become NightBlade
			case "Become NightBlade":
				var accepted bool
				*character, accepted = characters.NightBlade(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Nightblade\n%s\n\n", trim, trim)
				} else {
					color.HiRed("\n\n%sYou cannot become a Nightblade\n%s\n\n", trim, trim)
				}
			
			// Become Duelist
			case "Become Duelist":
				var accepted bool
				*character, accepted = characters.Duelist(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Duelist\n%s\n\n", trim, trim)
				} else {
					color.HiRed("\n\n%sYou cannot become a Duelist\n%s\n\n", trim, trim)
				}
				

			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}