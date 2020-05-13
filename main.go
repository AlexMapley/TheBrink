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

func main() {
	metaGame := world.MetaGame{
		Day: 1,
	}

	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	player := characters.NewPlayer(name)

	// main game loop
	for (player.Character.CurrentHealth >= 0) {
		fmt.Println("Your Stats:")
		player.Character.Stats()

		townConsole := console.NewTownConsole()

		color.Cyan("Day %d in town, what do you?\n", metaGame.Day)
		option := townConsole.ChooseAction()

		color.Cyan("\n\nYou have chosen option %d, %s", option+1, townConsole.Actions[option])

		if townConsole.Actions[option] == "Patrol the town" {
			player.Character.Strength++
			player.Character.Agility++
			player.Character.Vitality++
		} 
		if townConsole.Actions[option] == "Rest" {
			color.Cyan("\n\n\nFOOOOOOO\n\n\n")
			player.Character.Rest()
		}

		fmt.Println("\n\nA strange bandit appears")

		bandit := characters.NewBandit("Mel")

		player.Character.Duel(&bandit.Character)

		// Day ends
		metaGame.Day++
	}

	color.Cyan("\n\nGame Over %s, Day %d\n\n\n", player.Character.Name, metaGame.Day)
	fmt.Println("Your Stats:")
	player.Character.Stats()
	color.Cyan("\n\nOne day later (Day %d), %s is dead.\n\n\n", metaGame.Day, player.Character.Name)
}
