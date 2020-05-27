package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"the_brink/characters"
	"the_brink/console"
	"the_brink/party"
	"the_brink/world"

	"github.com/fatih/color"
)

var trim string = "-----------------------------------------\n"
var player characters.Player
var playerParty party.Party


func main() {
	// Create Character
	color.Cyan("What is your name?\n")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	// Choose Character
	color.Cyan("\nWhat Class Do You Pick?\n")
	classConsole := console.NewClassConsole()
	for len(player.Character.Stats.Class) == 0 {
		option := classConsole.ChooseAction()
		if option > 0 && option <= len(classConsole.Actions) {
			switch classConsole.Actions[option-1] {
			case "Rogue":
				player = characters.NewPlayer(name, "rogue")
				break
			case "Warrior":
				player = characters.NewPlayer(name, "warrior")
				break
			case "Wizard":
				player = characters.NewPlayer(name, "wizard")
				break
			default:
				continue
			}
		}
	}

	// Create Party
	playerParty = party.Party{
		X: 25,
		Y: 25,
		Members: []*characters.Character{
			&player.Character,
		},
		Rune: 'A',
	}

	// Build World
	saveFile := world.SaveFile{
		Day: 1,
		Party: playerParty,
	}
	world := world.World{
		XMax: 100,
		YMax: 50,
		Tiles: world.CreateMap(100, 50),
		Parties: []*party.Party{
			&playerParty,
		},
	}

	// Set Map
	world.UpdateMap()

	// main game loop
	for (player.Character.Stats.Health > 0) {

		if saveFile.Day == 15 {
			color.Magenta("\n\n%sYou feel a darnkness come over the land...\n%s\n\n", trim, trim)
		}

		color.Green("%sDay %d in town, what do you?\n%s", trim, saveFile.Day, trim)
		townConsole := console.NewTownConsole()

		dayLoop:
		for {
			option := townConsole.ChooseAction()

			if option > 0 && option <= len(townConsole.Actions) {
				color.Green("You have chosen option %d, %s", option, townConsole.Actions[option-1])

				switch townConsole.Actions[option-1] {
				
				// Character Menu
				case "Character":
					console.DisplayCharacterConsole(&player.Character)

				// Map
				case "Map":
					console.DisplayMapConsole(&world, &playerParty)
				
				// Fight
				case "Patrol the town":

					if (saveFile.Day % 2 == 0) {
						fmt.Println("\n\nA strange bandit appears")
						bandit := characters.NewBandit("Mel", player.Character.Stats.Level)

						player.Character.Duel(&bandit.Character)
					
						// loot bandit if won
						if (player.Character.Stats.Health > 0) {
							player.Character.Inventory.Loot(&bandit.Character.Inventory)
						}
					} else {
						fmt.Println("\n\nAn agry thug appears")
						thug := characters.NewThug("Dougy", player.Character.Stats.Level)

						player.Character.Duel(&thug.Character)
					
						// loot thug if won
						if (player.Character.Stats.Health > 0) {
							player.Character.Inventory.Loot(&thug.Character.Inventory)
						}
					}

					break dayLoop

				// Rest Party
				case "Rest":
					playerParty.Rest()
					color.Green("Your party's stats have been restored\n")
					break dayLoop

				// Become Paladin
				case "Become Paladin":
					var accepted bool
					player.Character, accepted = characters.Paladin(player.Character)
					if !accepted {
						color.HiRed("\n\n%sAttempting to become a Paladin, you fail\n%s\n\n", trim, trim)
					}
					break dayLoop

				// Become NightBlade
				case "Become NightBlade":
					var accepted bool
					player.Character, accepted = characters.NightBlade(player.Character)
					if !accepted {
						color.HiRed("\n\n%sAttempting to become a Nightblade, you foolishly die\n%s\n\n", trim, trim)
					}
					break dayLoop
				}
			}
		}

		// Day ends
		saveFile.Day++
	}

	color.Cyan("\n\nGame Over %s, Day %d\n\n\n", player.Character.Stats.Name, saveFile.Day)
	fmt.Println("Your Stats:")
	player.Character.Stats.Display()
	color.Cyan("\n\nOne day later (Day %d), %s is dead.\n\n\n", saveFile.Day, player.Character.Stats.Name)
}
