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


	HealthMax int
	FocusMax   int
	Health int
	Focus int

	Strength     int
	Vitality     int
	Agility      int
	Intelligence int
}

func (stats *Stats) Display() {
	color.Cyan("\n-------------\n%s\n-------------\nLevel: %d\nClass: %s\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\n\n",
		stats.Name,
		stats.Level,
		stats.Class,
		stats.Health,
		stats.HealthMax,
		stats.FocusMax,
		stats.Focus,
		stats.Vitality,
		stats.Strength,
		stats.Agility,
		stats.Intelligence,
		stats.CriticalValue(),
		stats.DodgeValue(),
	)
}

func (stats *Stats) CriticalValue() int {
	return stats.Agility * 2
}

func (stats *Stats) DodgeValue() int {
	return stats.Agility * 3
}

func (stats *Stats) DetermineMaxHealth() int {
	return (stats.Vitality * 8) + (stats.Strength * 2)
}

func (stats *Stats) DetermineMaxFocus() int {
	return stats.Intelligence * 10
}

func (stats *Stats) DisplayHealth() string {
	return fmt.Sprintf("(%d/d)", stats.Health, stats.MaxHealth)
}

