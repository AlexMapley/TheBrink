package console

import (
	"the_brink/characters"

	"github.com/fatih/color"
)

func NewClassConsole(character *characters.Character) Console {
	// Create default actions
	actions := make([]string, 7)
	actions[0] = "Rogue"
	actions[1] = "Warrior"
	actions[2] = "Wizard"
	if (character.Stats.Level >= 5) {
		actions[3] = "Paladin"
		actions[4] = "Duelist"
		actions[5] = "Nightblade"
		actions[6] = "x | q | esc to exit"
	} else {
		actions[3] = "x | q | esc to exit"
	}

	console := Console{}
	console.Actions = actions
	return console
}

func DisplayClassConsole(character *characters.Character) {

	console := NewClassConsole(character)

	menuLoop:
	for {
		option := console.ChooseAction()
		if option == -1 {
			break menuLoop
		}

		color.HiRed("\n\nGot option %d\n", option)

		if option > 0 && option <= len(console.Actions) {
			switch console.Actions[option-1] {

			// Become Warrior
			case "Warrior":
				var accepted bool
				*character, accepted = characters.Warrior(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Warrior\n%s\n\n", trim, trim)
					character.Rest()
					break menuLoop
				} else {
					color.HiRed("\n\n%sYou cannot become a Warrior\n%s\n\n", trim, trim)
				}

			// Become Rogue
			case "Rogue":
				var accepted bool
				*character, accepted = characters.Rogue(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Rogue\n%s\n\n", trim, trim)
					character.Rest()
					break menuLoop
				} else {
					color.HiRed("\n\n%sYou cannot become a Rogue\n%s\n\n", trim, trim)
				}

			// Become Duelist
			case "Wizard":
				var accepted bool
				*character, accepted = characters.Wizard(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Wizard\n%s\n\n", trim, trim)
					character.Rest()
					break menuLoop
				} else {
					color.HiRed("\n\n%sYou cannot become a Wizard\n%s\n\n", trim, trim)
				}

			// Become Paladin
			case "Paladin":
				var accepted bool
				*character, accepted = characters.Paladin(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Paladin\n%s\n\n", trim, trim)
					character.Rest()
					break menuLoop
				} else {
					color.HiRed("\n\n%sYou cannot become a Paladin\n%s\n\n", trim, trim)
				}

			// Become NightBlade
			case "Nightblade":
				var accepted bool
				*character, accepted = characters.NightBlade(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Nightblade\n%s\n\n", trim, trim)
					character.Rest()
					break menuLoop
				} else {
					color.HiRed("\n\n%sYou cannot become a Nightblade\n%s\n\n", trim, trim)
				}

			// Become Duelist
			case "Duelist":
				var accepted bool
				*character, accepted = characters.Duelist(*character)
				if accepted {
					color.HiGreen("\n\n%sYou have become a Duelist\n%s\n\n", trim, trim)
					character.Rest()
					break menuLoop
				} else {
					color.HiRed("\n\n%sYou cannot become a Duelist\n%s\n\n", trim, trim)
				}
			}
		}
	}
}
