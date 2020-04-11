package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	name, _ := reader.ReadString('\n')

	player := newPlayer(name)
	player.Character.Stats()

	fmt.Printf("\n\nGame Over, %s\n\n\n", player.Character.Name)
}
