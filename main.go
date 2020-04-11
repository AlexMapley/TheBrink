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

	player := PlayerCharacter{}
	player.Character.Name = text
	player.Character.Health = 10

	fmt.Printf("Game Over, %s\n", player.Character.Name)
}
