package console

import (
	"fmt"
	"strconv"
)

type Console struct {
	Actions map[int]string
}

func (console *Console) ChooseAction() int {

	// List Potential Actions
	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		fmt.Printf("%d. %s\n", number, option)
	}

	var inputstr string
	_, err := fmt.Scanf("%s", &inputstr)
	if err != nil {
		logError(err)
	}

	input, e := strconv.Atoi(inputstr)
	if e != nil {
		logError(err)
	}
	// output := (input *

	return input
}

func logError(err error) {
	fmt.Printf("\n\n-------------------\nEncountered Error: %s\n-------------------\n\n", err.Error())
}
