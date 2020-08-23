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
	player := characters.NewPlayer(name)
	color.Cyan("\nWhat Class Do You Pick?\n")
	console.DisplayClassConsole(&player.Character)

	// Create Player Party
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
		Day:   1,
		Party: &playerParty,
	}
	world := world.World{
		XMax:  100,
		YMax:  50,
		Tiles: world.CreateMap(100, 50),
		Parties: []*party.Party{
			&playerParty,
		},
	}

	// Set Map
	world.UpdateMap()

	// MAIN GAME LOOP
	// the inner `dayCounter`
	for player.Character.Stats.Health > 0 {

		if saveFile.Day == 15 {
			color.Magenta("\n\n%sYou feel a darnkness come over the land...\n%s\n\n", trim, trim)
		}

		color.Green("%sDay %d in town, what do you?\n%s", trim, saveFile.Day, trim)
		townConsole := console.NewTownConsole()

	dayCounter:
		for {
			option := townConsole.ChooseAction()
			if option == -1 {
				player.Character.Stats.Health -= 10000
				break dayCounter
			}

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
				case "Patrole the town":
					color.HiRed("\n\n%sNew Fight\n%s\n\n", trim, trim)
					fmt.Println("\n** An agry thug appears")
					fmt.Println("\n** A strange bandit appears")
					fmt.Println("\n** An weaker trainee appears")

					bandit := characters.NewBandit("Mel", player.Character.Stats.Level)
					thug1 := characters.NewThug("Dougy", player.Character.Stats.Level)
					thug2 := characters.NewThug("Dillan", player.Character.Stats.Level/2)
					enemyParty := party.Party{
						X: 25,
						Y: 25,
						Members: []*characters.Character{
							&thug1.Character,
							&bandit.Character,
							&thug2.Character,
						},
						Rune: 'B',
					}
					playerParty.Battle(&enemyParty)

					// loot enemy party if won
					if player.Character.Stats.Health > 0 {
						player.Character.Inventory.Loot(&bandit.Character.Inventory)
						player.Character.Inventory.Loot(&thug1.Character.Inventory)
						player.Character.Inventory.Loot(&thug2.Character.Inventory)
					}
					break dayCounter

				// Rest Party
				case "Rest":
					playerParty.Rest()
					color.Green("Your party's stats have been restored\n")
					break dayCounter

				// Level Up Party
				case "Level Up":
					// level up player and bandit
					success := player.Character.LevelUp()
					if success {
						color.Yellow("\n\n%sYou have leveled up\n%s\n\n", trim, trim)
						player.Character.Rest()
					} else {
						color.Red("\n\n%sNot enough xp to level up\n%s\n\n", trim, trim)
					}

				// Time Warp
				case "Time Warp":
					for i := 0; i < 5; i++ {
						color.HiRed("\n\n%sNew Fight\n%s\n\n", trim, trim)
						fmt.Println("\n** An agry thug appears")
						fmt.Println("\n** A strange bandit appears")
						fmt.Println("\n** An weaker trainee appears")
	
						bandit := characters.NewBandit("Mel", player.Character.Stats.Level)
						thug1 := characters.NewThug("Dougy", player.Character.Stats.Level)
						thug2 := characters.NewThug("Dillan", player.Character.Stats.Level/2)
						enemyParty := party.Party{
							X: 25,
							Y: 25,
							Members: []*characters.Character{
								&thug1.Character,
								&bandit.Character,
								&thug2.Character,
							},
							Rune: 'B',
						}
						playerParty.Battle(&enemyParty)

						// Die as well
						if player.Character.Stats.Health <= 0 {
							break dayCounter
						}
	
						// loot enemy party if won
						if player.Character.Stats.Health > 0 {
							player.Character.Inventory.Loot(&bandit.Character.Inventory)
							player.Character.Inventory.Loot(&thug1.Character.Inventory)
							player.Character.Inventory.Loot(&thug2.Character.Inventory)
						}
						playerParty.Rest()
					}
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
