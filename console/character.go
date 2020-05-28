package console

import (
	"the_brink/characters"
	
	"github.com/fatih/color"
)


func NewCharacterConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 7)
	actions[0] = "Stats"
	actions[1] = "Inventory"
	actions[2] = "Level Up"
	actions[3] = "Exit"
	actions[4] = "Become Paladin"
	actions[5] = "Become NightBlade"
	actions[6] = "Become Duelist"

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
				player.Character, accepted = characters.Paladin(player.Character)
				if !accepted {
					color.HiBlue("\n\n%sYou have become a Paladin\n%s\n\n", trim, trim)
				} else {
					color.HiBlue("\n\n%You have become a Paladin\n%s\n\n", trim, trim)
				}

			// Become NightBlade
			case "Become NightBlade":
				var accepted bool
				player.Character, accepted = characters.NightBlade(player.Character)
				if !accepted {
					color.HiRed("\n\n%sYou cannot become a Nightblade\n%s\n\n", trim, trim)
				} else {
					color.HiBlue("\n\n%sYou have become a Nightblade\n%s\n\n", trim, trim)
				}
			
			// Become Duelist
			case "Become Duelist":
				var accepted bool
				player.Character, accepted = characters.Duelist(player.Character)
				if !accepted {
					color.HiRed("\n\n%sYou cannot become a Duelist\n%s\n\n", trim, trim)
				} else {
					color.HiBlue("\n\n%sYou have become a Duelist\n%s\n\n", trim, trim)
				}
				

			// Exit
			case "Exit":
				break menuLoop
			}
		}
	}
}