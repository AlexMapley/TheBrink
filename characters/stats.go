package characters

import (
	"github.com/fatih/color"
)

type Stats struct {
	Name string
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
	color.Cyan("\n-------------\n%s\n-------------\nLevel: %d\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\n\n",
		stats.Name,
		stats.Level,
		stats.Health,
		stats.HeallthMax,
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

func (stats *Stats) HealthValue() int {
	return (stats.Vitality * 8) + (stats.Strength * 2)
}

func (stats *Stats) FocusValue() int {
	return stats.Intelligence * 10
}
