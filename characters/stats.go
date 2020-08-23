package characters

import (
	"fmt"
)

func (stats Stats) CriticalValue() float64 {
	return float64(stats.Agility * 1.45) - float64(stats.Level * 2.25) + float64(stats.Critical)
}

func (stats Stats) DodgeValue() float64 {
	return float64(stats.Agility * 1.85) - float64(stats.Level * 2.05)
}

func (stats Stats) BlockValue() float64 {
	return float64(stats.Strength * 1.9) - float64(stats.Level * 2.05) + float64(stats.Block)
}

func (stats Stats) AccuracyRating() float64 {
	return float64(stats.Agility) + float64(stats.Expertise)
}

func (stats Stats) MaxHealth() float64 {
	return float64(stats.Vitality * 8) + float64(stats.Strength * 2)
}

func (stats Stats) MaxFocus() float64 {
	return float64(stats.Intelligence * 5)
}

func (stats Stats) DisplayHealth() string {
	return fmt.Sprintf("(%d/%d)", stats.Health, stats.MaxHealth())
}
