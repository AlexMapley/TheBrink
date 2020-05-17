package characters

import (
	"github.com/fatih/color"
)

func (stats *Stats) Display() {
	color.Cyan("\n-------------\n%s\n-------------\nLevel: %d\nClass: %s\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\nAccuracy Rating: %d\n\n",
		stats.Name,
		stats.Level,
		stats.Class,
		stats.Health,
		stats.GetMaxHealth(),
		stats.GetMaxFocus(),
		stats.Focus,
		stats.Vitality,
		stats.Strength,
		stats.Agility,
		stats.Intelligence,
		stats.GetCriticalValue(),
		stats.GetDodgeValue(),
		stats.GetAccuracyRating(),
	)
}