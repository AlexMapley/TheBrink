package character

import (
	"github.com/fatih/color"
)

type Stats struct {
	Name string
	Level int

	Health int
	Focus   int
	CurrentHealth int
	CurrentFocus int

	Strength     int
	Vitality     int
	Agility      int
	Intelligence int
}

func (stats *Stats) Display() {
	color.Cyan("\n-------------\n%s\n-------------\nLevel: %d\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\n\n",
		character.Name,
		character.Level,
		character.CurrentHealth,
		character.Health,
		character.Focus,
		character.CurrentFocus,
		character.Vitality,
		character.Strength,
		character.Agility,
		character.Intelligence,
		character.CriticalValue(),
		character.DodgeValue(),
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
