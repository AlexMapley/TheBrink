package characters

import (
	"github.com/fatih/color"
)

func (stats *Stats) Display() {
	color.Cyan("\n-------------\n%s\n-------------\nLevel: %d\nClass: %s\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\nBlock: %d\nAccuracy Rating: %d\n\n",
		stats.Name,
		stats.Level,
		stats.Class,
		stats.Health,
		stats.MaxHealth(),
		stats.Focus,
		stats.MaxFocus(),
		stats.Vitality,
		stats.Strength,
		stats.Agility,
		stats.Intelligence,
		stats.CriticalValue(),
		stats.DodgeValue(),
		stats.BlockValue(),
		stats.AccuracyRating(),
	)
}