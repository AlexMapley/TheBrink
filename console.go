package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Console struct {
	Actions map[int]string
}

func (console *Console) ChooseAction() int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose option:")
	for number, option := range console.Actions {
		fmt.Printf("%d. %s\n", number, option)
	}

	action, _ := reader.ReadString('\n')
	fmt.Printf("%s%s%s%s", action, action, action, action)
	action = strings.TrimSuffix(fmt.Sprintf("%s", action), "\n")
	fmt.Printf("%s%s%s%s", action, action, action, action)

	choice, err := strconv.Atoi(action)
	if err == nil {
		panic(err)
	}

	// choice, err := strconv.ParseInt(action, 10, 32)
	// if err == nil {
	// 	panic(err)
	// }

	return int(choice)
}
