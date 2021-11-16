package main

import (
	"bufio"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"the_brink/characters"
	"the_brink/console"
	"the_brink/party"
	"the_brink/world"

	"github.com/Pallinder/go-randomdata"
	"github.com/fatih/color"
)

var trim string = "-----------------------------------------\n"
var player characters.Player
var playerParty party.Party

// SaveFile holds all information from a playthrough
type SaveFile struct {
	Day   int
	Party *party.Party
}

func main() {

	// Build World
	saveFile := SaveFile{
		Day:   1,
		Party: &playerParty,
	}
	world := world.World{
		XMax:  81,
		YMax:  25,
		Tiles: world.CreateMap(81, 25),
		Parties: []*party.Party{
			&playerParty,
		},
	}

	// Run pprof in the background
	// check it out at http://localhost:6060/debug/pprof/
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

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
		X: 42,
		Y: 16,
		Members: []*characters.Character{
			&player.Character,
			&sidekick.Character,
		},
		Rune: 65, // 'A'
	}

	// Display Map
	mapScreen := world.PrintMap()
	color.Green("%s", trim)
	color.Cyan("%s", mapScreen)
	color.Green("%s", trim)

	// MAIN GAME LOOP
dayCounter:
	for playerParty.GetHealth() > 0 {
		// Rest and reset abilities
		playerParty.Rest()

		if saveFile.Day >= 15 {
			color.Magenta("\n\n%sYou feel a darnkness come over the land...\n%s\n", trim, trim)
		}

		color.Green("%sDay %d in town, what do you?\n%s", trim, saveFile.Day, trim)
		townConsole := console.NewTownConsole()

		option := townConsole.ChooseAction(&playerParty, &world)

		// If option is -1, we have hit an escape key and
		// are trying to quit out of the main game loop
		if option == -1 {
			break dayCounter
		}

		// Main keyboard switch for non map options
		if option > 0 && option <= len(townConsole.Actions) {
			color.Green("You have chosen option %d, %s", option, townConsole.Actions[option-1])

			switch townConsole.Actions[option-1] {

			// Character Menu
			case "Character":
				console.DisplayCharacterConsole(&player.Character)

			// Defend the Town
			case "Defend the Town":
				color.HiRed("\n\n%sNew Fight\n%s\n", trim, trim)
				fmt.Println("** An agry thug blocks you off")
				fmt.Println("** A strange bandit appears not far behind")
				fmt.Println("** An shifty trickster lurks beside them")

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

			// Level Up Party
			case "Level Up":
				// level up player and bandit
				success := playerParty.LevelUp()
				if success {
					color.Yellow("\n\n%sYou have leveled up\n%s\n\n", trim, trim)
				} else {
					color.Red("\n\n%sNot enough xp to level up\n%s\n\n", trim, trim)
				}

			// Time Warp
			case "Time Warp":
				for i := 0; i < 5; i++ {
					color.HiRed("\n\n%sNew Fight\n%s\n", trim, trim)
					fmt.Println("** An agry thug appears")
					fmt.Println("** A strange bandit appears")
					fmt.Println("** An shifty rogue appears")

					bandit := characters.NewBandit(randomdata.FirstName(randomdata.Female), player.Character.Stats.Level)
					thug := characters.NewThug(randomdata.FirstName(randomdata.Male), player.Character.Stats.Level-1)
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

					// Loot enemy party if won
					if player.Character.Stats.Health > 0 {
						player.Character.Inventory.Loot(&bandit.Character.Inventory)
						player.Character.Inventory.Loot(&thug.Character.Inventory)
						player.Character.Inventory.Loot(&trickster.Character.Inventory)
					} else {
						break dayCounter
					}

					player.Character.LevelUp()
					sidekick.Character.LevelUp()

					// Increment day
					saveFile.Day++
				}
				player.Character.Stats.Display()
				sidekick.Character.Stats.Display()
			}
		}
		// Day ends
		if option > 0 {
			saveFile.Day++
		}
	}

	color.Cyan("\n\nGame Over %s, Day %d\n\n\n", player.Character.Stats.Name, saveFile.Day)
	fmt.Println("Your Stats:")
	player.Character.Stats.Display()
	color.Cyan("\n\nOne day later (Day %d), %s is dead.\n\n\n", saveFile.Day, player.Character.Stats.Name)
}
