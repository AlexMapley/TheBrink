package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var DayCounter int

func main() {
	metaGame = Game{
		Day: 1,
	}

	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	player := NewPlayer(name)
	player.Character.Stats()

	townConsole := NewTownConsole()

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

	fmt.Printf("\n\nGame Over %s, Day %d\n\n\n", player.Character.Name, DayCounter)
}
