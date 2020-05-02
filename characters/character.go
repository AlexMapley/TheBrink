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
	fmt.Printf("\n-------------\n%s\n-------------\nHealth: %d/%d\nMana: %d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\n\n",
		character.Name,
		character.CurrentHealth,
		character.Health,
		character.Mana,
		character.Vitality,
		character.Strength,
		character.Agility,
		character.Intelligence,
		character.Critical(),
	)
}

func (self *Character) Critical() int {
	return self.Agility * 2
}

func (self *Character) Dodge() int {
	return self.Agility * 3
}
