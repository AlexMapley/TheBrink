package game

import "fmt"

func logError(err error) {
	fmt.Printf("\n\n-------------------\nEncountered Error: %s\n-------------------\n\n", err.Error())
}
