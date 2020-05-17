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
	SkillSlots []Skill
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
type Skill struct {
	Name string
	Cost int
	CoolDownMax int
	CoolDown int
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

