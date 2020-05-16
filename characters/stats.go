package characters

import (
	"fmt"

	"github.com/fatih/color"
)

type Stats struct {
	Name string
	Class string
	Level int

	XPCap int


	MaxHealth int
	MaxFocus   int
	Health int
	Focus int

	Strength     int
	Vitality     int
	Agility      int
	Intelligence int

	LevelBonuses LevelBonuses
}

type LevelBonuses struct {
	Strength     int
	Vitality     int
	Agility      int
	Intelligence int
}



func (stats Stats) CriticalValue() int {
	return stats.Agility * 2
}

func (stats Stats) DodgeValue() int {
	return stats.Agility * 3
}

func (stats Stats) DetermineMaxHealth() int {
	return (stats.Vitality * 8) + (stats.Strength * 2)
}

func (stats Stats) DetermineMaxFocus() int {
	return stats.Intelligence * 10
}

func (stats Stats) DisplayHealth() string {
	return fmt.Sprintf("(%d/%d)", stats.Health, stats.MaxHealth)
}

