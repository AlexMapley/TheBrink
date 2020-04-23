package characters

import "fmt"

type Character struct {
	Name string

	Health int64
	CurrentHealth int64
	Mana   int64

	Strength     int64
	Vitality     int64
	Agility      int64
	Intelligence int64
}

func (character *Character) Stats() {
	fmt.Printf("\n-------------\n%s\n-------------\nHealth: %d/%d\nMana: %d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\n\n",
		character.Name,
		character.CurrentHealth,
		character.Health,
		character.Mana,
		character.Vitality,
		character.Strength,
		character.Agility,
		character.Intelligence,
	)
}

func (self *Character) Attack(other *Character) {
	other.Health -= self.Strength
}


func fight(character1, character2  *Character) (*Character, *Character) {
	return character1, character2
}