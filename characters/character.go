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


// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	for (self.CurrentHealth > 0 || other.CurrentHealth  > 0) {
		self.Attack(other)
		other.Attack(self)
	}
}