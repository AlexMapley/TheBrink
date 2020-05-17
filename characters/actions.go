package characters

import (
	"time"
	"math/rand"

	"github.com/fatih/color"
)
// Attack
func (self *Character) Attack(other *Character) {
	rand.Seed(time.Now().UnixNano())

	damage := self.Stats.Strength + (self.Stats.Agility/2)

	// Dodge Chance
	dodgeThreshold := rand.Intn(200)
	if (other.Stats.DodgeValue() >= dodgeThreshold) {
		damage = 0
		color.Cyan("%s dodges the hit\n", other.Stats.Name)
	}

	// Critical Chance
	criticalThreshold := rand.Intn(200)
	if (self.Stats.CriticalValue() >= criticalThreshold) {
		damage = damage*2
		color.Cyan("%s scores a critical hit\n", self.Stats.Name)
	}

	// Damage Multiplier 
	DamangeMultiplier := rand.Intn(100)
	damage = int(damage * DamangeMultiplier)
	damage = int(damage / 100)
	
	color.Red("%s %s deals %d damage\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	other.Stats.Health -= damage
}


// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	self.Stats.Display()
	other.Stats.Display()

	for (self.Stats.Health > 0 && other.Stats.Health  > 0) {

		time.Sleep(100 * time.Millisecond)

		// self Attack
		self.Attack(other)

		// other Attack
		if (other.Stats.Health  > 0) {
			other.Attack(self)
		}
	}

	if (self.Stats.Health >= other.Stats.Health) {
		color.Cyan("\n%s Wins the duel\n", self.Stats.Name)
		return
	}
	color.Cyan("%s Wins the duel\n", other.Stats.Name)
}

// Rest
func (self *Character) Rest() {
	self.Stats.MaxHealth = self.Stats.DetermineMaxHealth()
	self.Stats.MaxFocus = self.Stats.DetermineMaxFocus()

	self.Stats.Health = self.Stats.DetermineMaxHealth()
	self.Stats.Focus = self.Stats.DetermineMaxFocus()
}

// Level Up
func (self *Character) LevelUp() {

	// increase level
	self.Stats.Level++

	// increase core stats
	self.Stats.Vitality += self.Stats.LevelBonuses.Vitality
	self.Stats.Strength += self.Stats.LevelBonuses.Strength
	self.Stats.Agility += self.Stats.LevelBonuses.Agility
	self.Stats.Intelligence += self.Stats.LevelBonuses.Intelligence

	// rest
	self.Rest()
}
