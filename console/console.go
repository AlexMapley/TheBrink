package console

import (
	"fmt"
	"strconv"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

// ChooseAction
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

// ChooseDirection
func (console *Console) ChooseKey() string {

	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		color.Cyan("%d. %s\n", (number + 1), option)
	}

	_ := exec.Command("/bin/stty", "-F", "/dev/tty", "-icanon", "min", "1") // worked ok
	var keybuf [1]byte
	char, err := os.Stdin.Read(keybuf[0:1])
	if err != nil {
		fmt.Println(err.Error())
	}

	return char
}

func logError(err error) {
	color.Red("\n\n-------------------\nEncountered Error: %s\n-------------------\n\n", err.Error())
}
