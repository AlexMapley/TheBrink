package characters

import (
	"fmt"

	"github.com/fatih/color"
)

func (stats *Stats) Display() {
	text := fmt.Sprintf("\n-------------\n%s\n-------------\nLevel: %d\nXP: %d\nClass: %s\nClassHash: %d\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\nBlock: %d\nAccuracy Rating: %d\n\n",
		stats.Name,
		stats.Level,
		stats.XP,
		stats.Class,
		stats.ClassHash,
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
	color.HiCyan(text)
}

func (character *Character) DisplaySkills() {
	text := fmt.Sprintf("\n----------------\n%s's Skills\n----------------\n", character.Stats.Name)

	for _, skill := range character.SkillSlots {
		text += fmt.Sprintf("%s\nCost: %d\nMax Cooldown: %d\nInitial Cooldown: %d\n",
			skill.Name,
			skill.Cost,
			skill.CoolDownMax,
			skill.CoolDownInitial,
		)
		text += "------------\n"
	}
	
	color.HiCyan(text)
}
