package main

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
		fmt.Println(err)
	}

	input, e := strconv.Atoi(inputstr)
	if e != nil {
		fmt.Println(err)
	}
	// output := (input *

	return input
}
