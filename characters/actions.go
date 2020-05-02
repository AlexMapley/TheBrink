package characters

import (
	"fmt"
	"time"
	"math/rand"
)

// Attack
func (self *Character) Attack(other *Character) {
	damage := self.Strength + (self.Agility/2)

	criticalThreshold := rand.Intn(100)
	// fmt.Printf("Critical Threshold: %d\n", criticalThreshold)
	if (self.CriticalValue() >= criticalThreshold) {
		damage = damage*2
		fmt.Printf("%s scores a critical hit\n", self.Name)
	}

	dodgeThreshold := rand.Intn(100)
	// fmt.Printf("Dodge Threshold: %d\n", criticalThreshold)
	if (other.DodgeValue() >= dodgeThreshold) {
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

// Rest
func (self *Character) Rest() {
	self.Mana = self.Intelligence * 10
}