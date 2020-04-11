package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your plan?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Printf("Your answer: %s\nGame Over", text)
}
