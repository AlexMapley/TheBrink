package characters

import (
	"fmt"
)


type Character struct {
	Name string

	Health int
	CurrentHealth int
	Mana   int

	Strength     int
	Vitality     int
	Agility      int
	Intelligence int
}

func (character *Character) Stats() {
	fmt.Printf("\n-------------\n%s\n-------------\nHealth: %d/%d\nMana: %d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\n\n",
		character.Name,
		character.CurrentHealth,
		character.Health,
		character.Mana,
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
	return self.Intelligence * 10
}

func (self *Character) ManaValue() int {
	return (self.Vitality * 8) + (self.Strength * 2)
}
