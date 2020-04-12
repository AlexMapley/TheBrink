package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var DayCounter int

func main() {
	createWorld()

	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	player := NewPlayer(name)
	player.Character.Stats()

	fmt.Println("\n\nA bandit appears")

	bandit := NewBandit("Mel")
	bandit.Character.Stats()

	for i := 0; i < 10; i++ {
		fmt.Printf("\n\nBandit Attacks %s\n\n\n", player.Character.Name)
		bandit.Character.Attack(&player.Character)
		player.Character.Stats()
	}

	fmt.Printf("\n\nGame Over %s, Day %d\n\n\n", player.Character.Name, DayCounter)
}

func createWorld() {
	DayCounter = 1
}
