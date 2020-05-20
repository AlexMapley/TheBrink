package characters

import (
	"fmt"
)

func (stats Stats) CriticalValue() int {
	return int(float64(stats.Agility) * float64(1.45)) - stats.Level
}

func (stats Stats) DodgeValue() int {
	return (stats.Agility * 1.9) - stats.Level
}

func (stats Stats) BlockValue() int {
	return (stats.Strength * 1.9 + stats.Block) - stats.Level
}

func (stats Stats) AccuracyRating() int {
	return (stats.Agility + stats.Expertise)
}

func (stats Stats) MaxHealth() int {
	return (stats.Vitality * 8) + (stats.Strength * 2)
}

func (stats Stats) MaxFocus() int {
	return stats.Intelligence * 10
}

func (stats Stats) DisplayHealth() string {
	return fmt.Sprintf("(%d/%d)", stats.Health, stats.MaxHealth())
}

