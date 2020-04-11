package main

import (
	"bufio"
	"fmt"
	"os"
)

// var player PlayerCharacter

func main() {
	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')

	player := PlayerCharacter{}
	player.Character.Name = name
	player.Character.health = 10

	fmt.Printf("Your answer: %s\nGame Over", player.Character.Name)
}
