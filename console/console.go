package console

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/eiannone/keyboard"
)

// DisplayActions
func (console *Console) DisplayActions() {

	// List Potential Actions
	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		color.Cyan("%d. %s\n", (number + 1), option)
	}
}

// ChooseAction
func (console *Console) ChooseAction() (int, bool) {

	// Open keyboard
	if err := keyboard.Open(); err != nil {
		logError(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	// List Potential Actions
	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		color.Cyan("%d. %s\n", (number + 1), option)
	}

	// Read action as key input
	char, key, err := keyboard.GetKey()
	if err != nil {
		logError(err)
	}
	fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)


	option, err := strconv.Atoi(string(char))
	fmt.Printf("Casted to action %d\r\n", option)

	switch {
		// Exit
		case key == keyboard.KeyEsc, key == keyboard.KeyCtrlC,  key == keyboard.KeyCtrlD, char == 'q', char == 'x':
			return -1, true
		case err == nil && option >= 0:
			return option, false
		default:
			return -1, false
	}
}

func logError(err error) {
	color.Red("\n\n-------------------\nEncountered Error: %s\n-------------------\n\n", err.Error())
}
