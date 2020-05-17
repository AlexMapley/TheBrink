package characters

import (
	"time"
	"math/rand"

	"github.com/fatih/color"
)
// Attack
func (self *Character) Attack(other *Character) {
	rand.Seed(time.Now().UnixNano())

	// Events
	dodged := false
	critical := false

	// Generate base multipliers
	damage := self.Stats.Strength + (self.Stats.Agility/2)

	dodgeThreshold := 200 + (self.Stats.AccuracyRating() * 2)
	criticalThreshold := 200

	// Dodge Chance
	if (other.Stats.DodgeValue() >= rand.Intn(dodgeThreshold)) {
		dodged = true
		damage = 0
	}
	// Critical Chance,
	// cannot happen on a successful dodge
	if (self.Stats.CriticalValue() >= rand.Intn(criticalThreshold)) {
		critical = true
		damage = damage*2
	}

	// Damage Multiplier 
	DamangeMultiplier := rand.Intn(100)
	damage = int(damage * DamangeMultiplier)
	damage = int(damage / 100)
	
	// Event Cases
	if dodged {
		color.Yellow("%s %s deals %d damage (Dodge)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else if critical {
		color.HiRed("%s %s deals %d damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else {
		color.White("%s %s deals %d damage\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	}
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
	self.Stats.Health = self.Stats.MaxHealth()
	self.Stats.Focus = self.Stats.MaxFocus()
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
	self.Stats.Expertise += self.Stats.LevelBonuses.Expertise

	// rest
	self.Rest()
}
