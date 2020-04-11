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
	text, _ := reader.ReadString('\n')

	fmt.Printf("Your answer: %s\nGame Over", text)

	player := PlayerCharacter{}
	player.Character.Name = text
	player.Character.Health = 10

	fmt.Printf("Your answer: %s\nGame Over", player.Character.Name)
}
