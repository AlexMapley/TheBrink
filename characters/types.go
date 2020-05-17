package characters


import (
	"the_brink/inventory"
)

// Character
type Character struct {
	Stats Stats
}

// Stats
type Stats struct {
	// Info
	Name string
	Class string
	Level int
	LevelBonuses LevelBonuses

	// Core Attributes
	Strength     int
	Vitality     int
	Agility      int
	Intelligence int

	// Bonuses
	Expertise int
	Block int

	// Resource Pool
	Health int
	Focus int

	// Skills
	SkillSlots SkillSlot[]
}

// Level Bonsuses
type LevelBonuses struct {
	Strength     int
	Vitality     int
	Agility      int
	Intelligence int

	Expertise int
	Block int
}

// Skills
type SkillSlot {
	Name string,
	Cost int,
	CooldDown int,
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

