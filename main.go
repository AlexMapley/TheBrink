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
	"github.com/Pallinder/go-randomdata"
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

	sidekick := characters.NewSidekick("Doug")

	// Create Player Party
	playerParty = party.Party{
		X: 25,
		Y: 25,
		Members: []*characters.Character{
			&player.Character,
			&sidekick.Character,
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

				// Defend the Town
				case "Defend the Town":
					color.HiRed("\n\n%sNew Fight\n%s\n", trim, trim)
					fmt.Println("** An agry thug appears")
					fmt.Println("** A strange bandit appears")
					fmt.Println("** An weaker trainee appears")

					bandit := characters.NewBandit(randomdata.FirstName(randomdata.Female), player.Character.Stats.Level)
					thug := characters.NewThug(randomdata.FirstName(randomdata.Male), player.Character.Stats.Level)
					trickster := characters.NewTrickster(randomdata.FirstName(randomdata.Male), player.Character.Stats.Level-1)

					enemyParty := party.Party{
						X: 25,
						Y: 25,
						Members: []*characters.Character{
							&bandit.Character,
							&thug.Character,
							&trickster.Character,
						},
						Rune: 'B',
					}
					playerParty.Battle(&enemyParty)

					// loot enemy party if won
					if player.Character.Stats.Health > 0 {
						player.Character.Inventory.Loot(&bandit.Character.Inventory)
						player.Character.Inventory.Loot(&thug.Character.Inventory)
						player.Character.Inventory.Loot(&trickster.Character.Inventory)
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
					success := playerParty.LevelUp()
					if success {
						color.Yellow("\n\n%sYou have leveled up\n%s\n\n", trim, trim)
						player.Character.Rest()
					} else {
						color.Red("\n\n%sNot enough xp to level up\n%s\n\n", trim, trim)
					}

				// Time Warp
				case "Time Warp":
					for i := 0; i < 5; i++ {
						color.HiRed("\n\n%sNew Fight\n%s\n", trim, trim)
						fmt.Println("** An agry thug appears")
						fmt.Println("** A strange bandit appears")
						fmt.Println("** An weaker trainee appears")
	
						bandit := characters.NewBandit(randomdata.FirstName(randomdata.Female), player.Character.Stats.Level)
						thug := characters.NewThug(randomdata.FirstName(randomdata.Male), player.Character.Stats.Level)
						trickster := characters.NewTrickster(randomdata.FirstName(randomdata.Male), player.Character.Stats.Level-1)
	
						enemyParty := party.Party{
							X: 25,
							Y: 25,
							Members: []*characters.Character{
								&bandit.Character,
								&thug.Character,
								&trickster.Character,
							},
							Rune: 'B',
						}
						playerParty.Battle(&enemyParty)
	
						// loot enemy party if won
						if player.Character.Stats.Health > 0 {
							player.Character.Inventory.Loot(&bandit.Character.Inventory)
							player.Character.Inventory.Loot(&thug.Character.Inventory)
							player.Character.Inventory.Loot(&trickster.Character.Inventory)
						}

						playerParty.Rest()
						player.Character.LevelUp()
						sidekick.Character.LevelUp()
					}
					player.Character.Stats.Display()
					sidekick.Character.Stats.Display()
					break dayCounter
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
