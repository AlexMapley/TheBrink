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
	action, _ := reader.ReadString('\n')
	action = strings.TrimSuffix(action, "\n")

	index, err := strconv.Atoi(action)
	if err == nil {
		panic(err)
	}

	return index
}
