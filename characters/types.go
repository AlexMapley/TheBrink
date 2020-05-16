package characters

import (
	"time"
	"math/rand"

	"github.com/fatih/color"
)

type Character struct {
	Stats Stats
}

type Player struct {
	Character Character
	Inventory inventory.Inventory
}

type Bandit struct {
	Character Character
	Inventory inventory.Inventory
}