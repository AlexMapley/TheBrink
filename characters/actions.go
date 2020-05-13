package characters

import (
	"time"
	"math/rand"

	"github.com/fatih/color"
)

// Attack
func (self *Character) Attack(other *Character) {
	rand.Seed(time.Now().UnixNano())

	damage := self.Strength + (self.Agility/2)

	criticalThreshold := rand.Intn(100)
	// color.Cyan("Critical Threshold: %d\n", criticalThreshold)
	if (self.CriticalValue() >= criticalThreshold) {
		damage = damage*2
		color.Cyan("%s scores a critical hit\n", self.Name)
	}

	dodgeThreshold := rand.Intn(100)
	// color.Cyan("Dodge Threshold: %d\n", criticalThreshold)
	if (other.DodgeValue() >= dodgeThreshold) {
		damage = 0
		color.Cyan("%s dodges the hit\n", other.Name)
	}
	
	other.CurrentHealth -= damage
}


// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	self.Stats()
	other.Stats()

	for (self.CurrentHealth > 0 && other.CurrentHealth  > 0) {

		time.Sleep(100 * time.Millisecond)

		self.Attack(other)
		other.Attack(self)

		color.Cyan("%s Health: %d, ", self.Name, self.CurrentHealth)
		color.Cyan("%s Health: %d\n", other.Name, other.CurrentHealth)
	}

	if (self.CurrentHealth >= other.CurrentHealth) {
		color.Cyan("\n%s Wins the duel\n", self.Name)
		return
	}
	color.Cyan("%s Wins the duel\n", other.Name)
}

// Rest
func (self *Character) Rest() {
	self.Health = self.HealthValue()
	self.Focus = self.FocusValue()

	self.CurrentHealth = self.HealthValue()
	self.CurrentFocus = self.FocusValue()
}