package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var DayCounter int

func main() {
	metaGame := game.MetaGame{
		Day: 1,
	}

	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	player := game.NewPlayer(name)

	fmt.Println("Your Stats:")
	player.Character.Stats()

	townConsole := game.NewTownConsole()

	fmt.Printf("Day %d in town, what do you?\n", metaGame.Day)
	option := townConsole.ChooseAction()

	fmt.Printf("\n\nYou have chosen option %d, %s", option, townConsole.Actions[option])

	// fmt.Println("\n\nA bandit appears")

	// bandit := NewBandit("Mel")
	// bandit.Character.Stats()

	// battleConsole := NewBattleConsole()

	// for i := 0; i < 100; i++ {
	// 	actionIndex := battleConsole.ChooseAction()
	// 	fmt.Printf("You choose option: %s", battleConsole.Actions[actionIndex])
	// }

	fmt.Printf("\n\nGame Over %s, Day %d\n\n\n", player.Character.Name, metaGame.Day)
	metaGame.Day++
	fmt.Printf("\n\nOne day later (Day %d), %s is dead.\n\n\n", metaGame.Day, player.Character.Name)
}
