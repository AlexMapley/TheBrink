package characters

import (
	"fmt"
)

func (stats Stats) CriticalValue() float64 {
	return float64(stats.Agility) * 1.45 - float64(stats.Level) * 2.25 + float64(stats.Critical)
}

func (stats Stats) DodgeValue() float64 {
	return 1.85 * float64(stats.Agility)  - 2.05 * float64(stats.Level) 
}

func (stats Stats) BlockValue() float64 {
	return 1.9 * float64(stats.Strength) - 2.05  * float64(stats.Level) + float64(stats.Block)
}

func (stats Stats) AccuracyRating() float64 {
	return 0.25 * float64(stats.Strength) + 0.85 * float64(stats.Agility) + 1.15 * float64(stats.Expertise)
}

func (stats Stats) MaxHealth() float64 {
	return 8 * float64(stats.Vitality) + 2.2 * float64(stats.Strength)
}

func (stats Stats) MaxFocus() float64 {
	return 5 * float64(stats.Intelligence)
}

func (stats Stats) DisplayHealth() string {
	return fmt.Sprintf("(%d/%d)", stats.Health, stats.MaxHealth())
}
