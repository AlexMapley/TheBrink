package characters

import (
	"fmt"
	"time"
	"math/rand"
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

func (self *Character) Attack(other *Character) {
	damage := self.Strength + (self.Agility/2)

	criticalThreshold := rand.Intn(100)
	// fmt.Printf("Critical Threshold: %d\n", criticalThreshold)
	if (self.Critical() >= criticalThreshold) {
		damage = damage*2
		fmt.Printf("%s scores a critical hit\n", self.Name)
	}

	dodgeThreshold := rand.Intn(100)
	// fmt.Printf("Dodge Threshold: %d\n", criticalThreshold)
	if (other.Dodge() >= dodgeThreshold) {
		damage = 0
		fmt.Printf("%s dodges the hit\n", other.Name)
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

	if (self.CurrentHealth >= other.CurrentHealth) {
		fmt.Printf("\n%s Wins the duel\n", self.Name)
		return
	}
	fmt.Printf("%s Wins the duel\n", other.Name)
}