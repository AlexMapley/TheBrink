package characters

import (
	"fmt"

	"github.com/fatih/color"
)

func (stats *Stats) Display() {
	text := fmt.Sprintf"\n-------------\n%s\n-------------\nLevel: %d\nClass: %s\nClassHash: %d\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\nBlock: %d\nAccuracy Rating: %d\n\n",
		stats.Name,
		stats.Level,
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
	color.Cyan(text)
}

func (character *Character) DisplaySkills() {
	text := fmt.Sprintf"\n-------------\n%s's Skills\n-------------\n", character.Stats.Name)

	for _, skill := range character.Skillslots {
		text += fmt.Sprintf("Name: %s\nCost: %d\nMax Cooldown: %d\n Initial Cooldown: %d",
			skill.Name,
			skill.Cost,
			skill.CoolDownMax,
			skill.CoolDownInitial,
		)
	}
	color.Cyan(text)
}

// Skills
type Skill struct {
	Name string
	Cost int
	CoolDownInitial int
	CoolDownMax int
	CoolDown int
}
