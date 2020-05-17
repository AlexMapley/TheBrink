package characters

import (
	"github.com/fatih/color"
)

// Core Structs
type Character struct {
	Stats Stats
}

type Stats struct {
	Name string
	Class string
	Level int
	LevelBonuses LevelBonuses

	Strength     int
	Vitality     int
	Agility      int
	Intelligence int

	MaxHealth int
	MaxFocus int
	Health int
	Focus int
}

type LevelBonuses struct {
	Strength     int
	Vitality     int
	Agility      int
	Intelligence int
}

// Character Types
type Player struct {
	Character Character
	Inventory inventory.Inventory
}

type Bandit struct {
	Character Character
	Inventory inventory.Inventory
}

