package characters

import (
	"fmt"
)

func (stats Stats) CriticalValue() int {
	return int ( (float32(stats.Agility) * 1.45) - (float32(stats.Level) * 1.25) )  + float32(stats.Critical)
}

func (stats Stats) DodgeValue() int {
	return int ( (float32(stats.Agility) * 1.85) - (float32(stats.Level) *  1.15) )
}

func (stats Stats) BlockValue() int {
	return int ( (float32(stats.Strength) * 1.9) - (float32(stats.Level) * 1.2) )  + float32(stats.Block)
}

func (stats Stats) AccuracyRating() int {
	return int ( float32(stats.Agility) + float32(stats.Expertise) )
}

func (stats Stats) MaxHealth() int {
	return int ( (float32(stats.Vitality) * 8) + (float32(stats.Strength) * 2) )
}

func (stats Stats) MaxFocus() int {
	return int ( float32(stats.Intelligence) * 5 )
}

func (stats Stats) DisplayHealth() string {
	return fmt.Sprintf("(%d/%d)", stats.Health, stats.MaxHealth())
}

