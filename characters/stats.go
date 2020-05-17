package characters

import (
	"fmt"
)

func (stats Stats) GetCriticalValue() int {
	return stats.Agility * 2
}

func (stats Stats) GetDodgeValue() int {
	return stats.Agility * 3
}

func (stats Stats) GetAccuracyRating() int {
	return (stats.Agility + stats.Expertise)
}

func (stats Stats) GetMaxHealth() int {
	return (stats.Vitality * 8) + (stats.Strength * 2)
}

func (stats Stats) GetMaxFocus() int {
	return stats.Intelligence * 10
}

func (stats Stats) DisplayHealth() string {
	return fmt.Sprintf("(%d/%d)", stats.Health, stats.MaxHealth)
}

