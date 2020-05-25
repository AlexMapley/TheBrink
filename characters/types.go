package characters


import (
	"the_brink/inventory"
	"golang.org/x/mobile/geom"
)

// Character
type Character struct {
	Stats Stats
	Status Status
	SkillSlots []Skill
	Coordinates geom.Point
	Inventory inventory.Inventory
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

// Status
type Status struct {
	Stunned int
}

// Skills
type Skill struct {
	Name string
	Cost int
	CoolDownInitial int
	CoolDownMax int
	CoolDown int
}



// Character Types
type Player struct {
	Character Character
}

type Bandit struct {
	Character Character
}

type Thug struct {
	Character Character
}

