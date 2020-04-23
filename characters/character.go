package characters

import (
	"fmt"
	"time"
)

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
	other.CurrentHealth -= self.Strength
}


// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	for (self.CurrentHealth > 0 && other.CurrentHealth  > 0) {
		time.Sleep(1 * time.Second)
		self.Attack(other)
		other.Attack(self)

		fmt.Printf("%s Health: %d, ", self.Name, self.CurrentHealth)
		fmt.Printf("%s Health: %d\n", other.Name, other.CurrentHealth)
	}

	if (self.CurrentHealth <= other.CurrentHealth) {
		fmt.Printf("%s Wins the duel", self.Name)
		return
	}
	fmt.Printf("%s Wins the duel", other.Name)
}