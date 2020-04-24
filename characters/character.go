package characters

import (
	"fmt"
	"time"
	"math/rand"
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

func (self *Character) Critical() int64 {
	return self.Agility * 2
}

func (self *Character) Attack(other *Character) {
	damage := self.Strength + (self.Agility/2)

	criticalThreshold := int64(rand.Intn(100))
	fmt.Printf("Critical Threshold: %d\n", criticalThreshold)
	if (self.Critical() >= criticalThreshold) {
		damage = damage*2
		fmt.Printf("%s scores a critical hit\n", self.Name)
	}
	
	other.CurrentHealth -= damage
}


// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	for (self.CurrentHealth > 0 && other.CurrentHealth  > 0) {

		time.Sleep(100 * time.Millisecond)

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