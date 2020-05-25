package characters

import (
	"fmt"
)

func (stats Stats) CriticalValue() float32 {
	return float64(stats.Agility) * 1.45 - float32(stats.Level)
}

func (stats Stats) DodgeValue() float32 {
	return (float32(stats.Agility) * 1.9) - float32(stats.Level)
}

func (stats Stats) BlockValue() float32 {
	return float32(stats.Strength) * 1.9 + float32(stats.Block) - float32(stats.Level)
}

func (stats Stats) AccuracyRating() float32 {
	return float32(stats.Agility) + float32(stats.Expertise)
}

func (stats Stats) MaxHealth() float32 {
	return (float32(stats.Vitality) * 8) + (float32(stats.Strength) * 2)
}

func (stats Stats) MaxFocus() float32 {
	return float32(stats.Intelligence) * 10
}

func (stats Stats) DisplayHealth() string {
	return fmt.Sprintf("(%d/%d)", stats.Health, stats.MaxHealth())
}

