package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"the_brink/characters"
	"the_brink/console"
	"the_brink/world"

	"github.com/fatih/color"
)

var DayCounter int
var trim string = "-----------------------------------------\n"

func main() {
	metaGame := world.MetaGame{
		Day: 1,
	}

	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	player := characters.NewPlayer(name, "rogue")



	// main game loop
	for (player.Character.Stats.Health >= 0) {
		townConsole := console.NewTownConsole()

		color.Green("%sDay %d in town, what do you?\n%s", trim, metaGame.Day, trim)


		for {
			option := townConsole.ChooseAction()

			if option > 0 && option <= len(townConsole.Actions) {
				color.Green("You have chosen option %d, %s", option, townConsole.Actions[option-1])

				if townConsole.Actions[option-1] == "Stats" {
					player.Character.Stats.Display()
				}
				if townConsole.Actions[option-1] == "Inventory" {
					player.Inventory.Display()
				}
				if townConsole.Actions[option-1] == "Patrol the town" {

					// declare duel opponent
					bandit := characters.NewBandit("Mel", player.Character.Stats.Level)

					fmt.Println("\n\nA strange bandit appears")
					player.Character.Duel(&bandit.Character)
					
					// loot bandit if won duel
					if (player.Character.Stats.Health >= 0) {
						player.Inventory.Loot(&bandit.Inventory)
					}
					// reset bandit
					bandit.Character.Rest()
					break
				}
				if townConsole.Actions[option-1] == "Rest" {
					player.Character.Rest()
					fmt.Println("Your stats have been restored")
					break
				}
				if townConsole.Actions[option-1] == "Level Up" {
					// level up player and bandit
					player.Character.LevelUp()

					player.Character.Rest()

					break
				}
			}
		}

		// Day ends
		metaGame.Day++
	}

	color.Cyan("\n\nGame Over %s, Day %d\n\n\n", player.Character.Stats.Name, metaGame.Day)
	fmt.Println("Your Stats:")
	player.Character.Stats.Display()
	color.Cyan("\n\nOne day later (Day %d), %s is dead.\n\n\n", metaGame.Day, player.Character.Stats.Name)
}
