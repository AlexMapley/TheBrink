package console

import (
	"the_brink/characters"

	"github.com/fatih/color"
)

func NewClassConsole(character *characters.Character) Console {
	// Create default actions
	var actions []string
	console := Console{}

	if character.Stats.Level < 6 {
		actions := make([]string, 5)
		actions[0] = "Rogue"
		actions[1] = "Warrior"
		actions[2] = "Wizard"
		actions[3] = "Cleric"
		actions[4] = "x | q | esc to exit"

		console.Actions = actions
		return console
	}

	actions = make([]string, 8)
	actions[0] = "Rogue"
	actions[1] = "Warrior"
	actions[2] = "Wizard"
	actions[3] = "Cleric"
	actions[4] = "Paladin"
	actions[5] = "Swordsman"
	actions[6] = "Nightblade"
	actions[7] = "x | q | esc to exit"

	console.Actions = actions
	return console
}

func DisplayClassConsole(character *characters.Character) {

	console := NewClassConsole(character)

	menuLoop:
	for {
		option := console.ChooseAction(nil, nil)
		if option == -1 {
			break menuLoop
		}

		color.HiRed("\n\nGot option %d\n", option)

		if option > 0 && option <= len(console.Actions) {
			switch console.Actions[option-1] {

			// Become Warrior
			case "Warrior":
				var accepted bool
				accepted = character.MutateWarrior()
				if accepted {
					color.HiGreen("\n\n%s\nYou have become a Warrior\n%s\n\n", trim, trim)
					break menuLoop
				} else {
					color.HiRed("\n\n%s\nYou cannot become a Warrior\n%s\n\n", trim, trim)
				}

			// Become Rogue
			case "Rogue":
				var accepted bool
				accepted = character.MutateRogue()
				if accepted {
					color.HiGreen("\n\n%s\nYou have become a Rogue\n%s\n\n", trim, trim)
					break menuLoop
				} else {
					color.HiRed("\n\n%s\nYou cannot become a Rogue\n%s\n\n", trim, trim)
				}

			// Become Wizard
			case "Wizard":
				var accepted bool
				accepted = character.MutateWizard()
				if accepted {
					color.HiGreen("\n\n%s\nYou have become a Wizard\n%s\n\n", trim, trim)
					break menuLoop
				} else {
					color.HiRed("\n\n%s\nYou cannot become a Wizard\n%s\n\n", trim, trim)
				}

			// Become Wizard
			case "Cleric":
				var accepted bool
				accepted = character.MutateCleric()
				if accepted {
					color.HiGreen("\n\n%s\nYou have become a Cleric\n%s\n\n", trim, trim)
					break menuLoop
				} else {
					color.HiRed("\n\n%s\nYou cannot become a Cleric\n%s\n\n", trim, trim)
				}

			// Become Paladin
			case "Paladin":
				var accepted bool
				accepted = character.MutatePaladin()
				if accepted {
					color.HiGreen("\n\n%s\nYou have become a Paladin\n%s\n\n", trim, trim)
					break menuLoop
				} else {
					color.HiRed("\n\n%s\nYou cannot become a Paladin\n%s\n\n", trim, trim)
				}

			// Become NightBlade
			case "Nightblade":
				var accepted bool
				accepted = character.MutateNightBlade()
				if accepted {
					color.HiGreen("\n\n%s\nYou have become a Nightblade\n%s\n\n", trim, trim)
					break menuLoop
				} else {
					color.HiRed("\n\n%s\nYou cannot become a Nightblade\n%s\n\n", trim, trim)
				}

			// Become Swordsman
			case "Swordsman":
				var accepted bool
				accepted = character.MutateSwordsman()
				if accepted {
					color.HiGreen("\n\n%s\nYou have become a Swordsman\n%s\n\n", trim, trim)
					break menuLoop
				} else {
					color.HiRed("\n\n%s\nYou cannot become a Swordsman\n%s\n\n", trim, trim)
				}
			}
		}
	}
}
