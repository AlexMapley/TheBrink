package characters

import (
	"fmt"
)

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

