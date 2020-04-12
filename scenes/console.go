package scenes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Console struct {
	Options map[int]string
}

func (console *Console) ChooseOption() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose option:")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSuffix(option, "\n")

	index, err := strconv.Atoi(option)
	if err == nil {
		panic(err)
	}

	return index
}
