package characters


import (
	"the_brink/inventory"
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

	// Core Stats
	Strength     int
	Vitality     int
	Agility      int
	Intelligence int

	// Resource Pool
	MaxHealth int
	MaxFocus int
	Health int
	Focus int

	// Bonuses
	Expertise int
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

