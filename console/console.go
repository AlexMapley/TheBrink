package console

import (
	"fmt"
	"strconv"

	"the_brink/party"
	"the_brink/world"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

// DisplayActions
func (console *Console) DisplayActions() {

	// List Potential Actions
	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		if option != "" {
			color.Cyan("%d. %s\n", (number + 1), option)
		}
	}
}

// ChooseAction
func (console *Console) ChooseAction(playerParty *party.Party, gameWorld *world.World) int {

	// Open keyboard
	if err := keyboard.Open(); err != nil {
		logError(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	// List Potential Actions
	console.DisplayActions()

	// Read action as key input
	char, key, err := keyboard.GetKey()
	if err != nil {
		logError(err)
	}
	fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)

	option, err := strconv.Atoi(string(char))
	fmt.Printf("Casted to action %d\r\n", option)

	// Print from top left to bottom right
	if gameWorld != nil {
		mapScreen := (*gameWorld).PrintMap()
		color.Green("%s", trim)
		color.Cyan("%s", mapScreen)
		color.Green("%s", trim)
	}

	switch {
	// Move Left
	case key == keyboard.KeyArrowLeft, char == 'a':
		// TODO: Standardize ChooseAction() use cases so we 
		// don't have to sanitize inputs like this 
		if (playerParty == nil || gameWorld == nil) {
			return -2
		}
		gameWorld.Move(playerParty, -1, 0)
		gameWorld.UpdateMap()
		return -2

	// Move Right
	case key == keyboard.KeyArrowRight, char == 'd':
		if (playerParty == nil || gameWorld == nil) {
			return -2
		}
		gameWorld.Move(playerParty, 1, 0)
		gameWorld.UpdateMap()
		return -2

	// Move Up
	case key == keyboard.KeyArrowUp, char == 'w':
		if (playerParty == nil || gameWorld == nil) {
			return option
		}
		gameWorld.Move(playerParty, 0, 1)
		gameWorld.UpdateMap()
		return option

	// Move Down
	case key == keyboard.KeyArrowDown, char == 's':
		if (playerParty == nil || gameWorld == nil) {
			return option
		}
		gameWorld.Move(playerParty, 0, -1)
		gameWorld.UpdateMap()
		return option
	// Error Case
	case err == nil && option >= 0:
		return option
	// Exit
	case key == keyboard.KeyEsc, key == keyboard.KeyCtrlC, key == keyboard.KeyCtrlD, char == 'q', char == 'x':
		return -1
	default:
		return -1
	}
}

func logError(err error) {
	color.Red("\n\n-------------------\nEncountered Error: %s\n-------------------\n\n", err.Error())
}
