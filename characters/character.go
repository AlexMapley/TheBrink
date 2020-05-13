package characters

import (
	"time"
	"math/rand"

	"github.com/fatih/color"
)

type Character struct {
	Stats Stats
}

// Attack
func (self *Character) Attack(other *Character) {
	rand.Seed(time.Now().UnixNano())

	damage := self.Stats.Strength + (self.Stats.Agility/2)

	criticalThreshold := rand.Intn(100)
	// color.Cyan("Critical Threshold: %d\n", criticalThreshold)
	if (self.Stats.CriticalValue() >= criticalThreshold) {
		damage = damage*2
		color.Cyan("%s scores a critical hit\n", self.Stats.Name)
	}

	dodgeThreshold := rand.Intn(100)
	// color.Cyan("Dodge Threshold: %d\n", criticalThreshold)
	if (other.Stats.DodgeValue() >= dodgeThreshold) {
		damage = 0
		color.Cyan("%s dodges the hit\n", other.Stats.Name)
	}
	
	other.Stats.CurrentHealth -= damage
}


// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	self.Stats.Display()
	other.Stats.Display()

	for (self.Stats.CurrentHealth > 0 && other.Stats.CurrentHealth  > 0) {

		time.Sleep(100 * time.Millisecond)

		self.Attack(other)
		other.Attack(self)

		color.Cyan("%s Health: %d, ", self.Stats.Name, self.Stats.CurrentHealth)
		color.Cyan("%s Health: %d\n", other.Stats.Name, other.Stats.CurrentHealth)
	}

	if (self.Stats.CurrentHealth >= other.Stats.CurrentHealth) {
		color.Cyan("\n%s Wins the duel\n", self.Stats.Name)
		return
	}
	color.Cyan("%s Wins the duel\n", other.Stats.Name)
}

// Rest
func (self *Character) Rest() {
	self.Stats.Health = self.Stats.HealthValue()
	self.Stats.Focus = self.Stats.FocusValue()

	self.Stats.CurrentHealth = self.Stats.HealthValue()
	self.Stats.CurrentFocus = self.Stats.FocusValue()
}