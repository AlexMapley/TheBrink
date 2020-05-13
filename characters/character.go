package characters

import (
	"github.com/fatih/color"
)

type Character struct {
	Name string
	Level int

	Health int
	Focus   int
	CurrentHealth int
	CurrentFocus int

	Strength     int
	Vitality     int
	Agility      int
	Intelligence int
}

func (character *Character) Stats() {
	color.Cyan("\n-------------\n%s\n-------------\nLevel: %d\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\n\n",
		character.Name,
		character.Level,
		character.CurrentHealth,
		character.Health,
		character.Focus,
		character.CurrentFocus,
		character.Vitality,
		character.Strength,
		character.Agility,
		character.Intelligence,
		character.CriticalValue(),
		character.DodgeValue(),
	)
}

func (self *Character) CriticalValue() int {
	return self.Agility * 2
}

func (self *Character) DodgeValue() int {
	return self.Agility * 3
}

func (self *Character) HealthValue() int {
	return (self.Vitality * 8) + (self.Strength * 2)
}

func (self *Character) FocusValue() int {
	return self.Intelligence * 10
}

// stub method
func (self *Character) LevelUp() {

}