package console

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

type Console struct {
	Actions []string
}

func (console *Console) ChooseAction() int {

	// List Potential Actions
	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		color.Cyan("%d. %s\n", (number + 1), option)
	}

	var inputstr string
	_, err := fmt.Scanf("%s", &inputstr)
	if err != nil {
		logError(err)
	}

	input, e := strconv.Atoi(inputstr)
	if e != nil {
		return -1
	}

	return input
}

func logError(err error) {
	color.Red("\n\n-------------------\nEncountered Error: %s\n-------------------\n\n", err.Error())
}
