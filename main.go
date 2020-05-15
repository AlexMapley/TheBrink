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

	// declare game enemies
	bandit := characters.NewBandit("Mel")

	// main game loop
	for (player.Character.Stats.Health >= 0) {
		fmt.Println("Your Stats:")
		player.Character.Stats.Display()

		townConsole := console.NewTownConsole()

		color.Green("%sDay %d in town, what do you?\n%s", trim, metaGame.Day, trim)
		option := townConsole.ChooseAction()

		if option <= len(townConsole.Actions) {
			color.Cyan("\n\nYou have chosen option %d, %s", option+1, townConsole.Actions[option])
		}

		if townConsole.Actions[option] == "Patrol the town" {
			// level up player and bandit
			player = characters.LevelUpPlayer(player)
			bandit = characters.LevelUpBandit(bandit)
		} 
		if townConsole.Actions[option] == "Rest" {
			color.Cyan("\n\n\nFOOOOOOO\n\n\n")
			player.Character.Rest()
			bandit.Character.Rest()
		}

		fmt.Println("\n\nA strange bandit appears")

		

		player.Character.Duel(&bandit.Character)

		// Day ends
		metaGame.Day++
	}

	color.Cyan("\n\nGame Over %s, Day %d\n\n\n", player.Character.Stats.Name, metaGame.Day)
	fmt.Println("Your Stats:")
	player.Character.Stats.Display()
	color.Cyan("\n\nOne day later (Day %d), %s is dead.\n\n\n", metaGame.Day, player.Character.Stats.Name)
}
