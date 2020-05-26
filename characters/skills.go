package characters

import (
	"time"
	"math/rand"

	"github.com/fatih/color"
)

// BasicAttack
func (self *Character) BasicAttack(other *Character, base int) {
	rand.Seed(time.Now().UnixNano())

	// Events
	dodged := false
	critical := false
	blocked := false
	

	// Generate base multipliers
	damage := base

	dodgeThreshold := 320 + (self.Stats.AccuracyRating() * 2)
	criticalThreshold := 320
	blockThreshold := 250 + (self.Stats.Strength * 2) + (self.Stats.AccuracyRating())

	// Dodge Chance
	if (other.Stats.DodgeValue() >= rand.Intn(dodgeThreshold)) {
		dodged = true
		damage = 0
	}
	// Block Chance
	if (other.Stats.BlockValue() >= rand.Intn(blockThreshold)) {
		blocked = true
		damage = int (damage / 2)
	}
	// Critical Chance
	if (!blocked && self.Stats.CriticalValue() >= rand.Intn(criticalThreshold)) {
		critical = true
		damage = damage * 2
	}

	// Damage Multiplier 
	DamangeMultiplier := rand.Intn(100)
	damage = int(damage * DamangeMultiplier)
	damage = int(damage / 100)
	
	// Event Cases
	if dodged {
		color.Yellow("%s %s deals %d damage (Dodge)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else if blocked {
		color.Cyan("%s %s deals %d damage (Blocked)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else if critical {
		color.HiRed("%s %s deals %d damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else {
		color.White("%s %s deals %d damage\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	}
	other.Stats.Health -= damage
}

// DoubleStrike
func (self *Character) DoubleStrike(other *Character) {
	color.HiGreen("* %s uses Double Strike *\n", self.Stats.Name)
	self.BasicAttack(other)
	self.BasicAttack(other)
}

// GhostBlade
func (self *Character) GhostBlade(other *Character) {
	color.HiGreen("* %s uses GhostBlad *\n", self.Stats.Name)
	self.BasicAttack(other, (self.Stats.Strength*.7) + (self.Stats.Agility*.8) + (self.Stats.Intelligence*.7))

	color.Magenta("%s %s stuns for 1 turn\n", self.Stats.Name, self.Stats.DisplayHealth())

	other.Status.Stunned +=1
}

// Heal
func (self *Character) Heal() {
	color.HiGreen("* %s uses Heal *\n", self.Stats.Name)

	heal := int(float64(self.Stats.Intelligence) * 1.9)
	color.Magenta("%s %s Heals %d damage\n", self.Stats.Name, self.Stats.DisplayHealth(), heal)

	self.Stats.Health += heal
	if self.Stats.Health > self.Stats.MaxHealth() {
		self.Stats.Health = self.Stats.MaxHealth()
	}
}

// IceBlast
func (self *Character) IceBlast(other *Character) {
	color.HiGreen("* %s uses Ice Blast *\n", self.Stats.Name)
	damage := int(float64(self.Stats.Intelligence) * 1.8)
	color.Magenta("%s %s deals %d magic damage, and stuns for 2 turns\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)

	other.Stats.Health -= damage
	other.Status.Stunned +=2
}

// LightningBolt
func (self *Character) LightningBolt(other *Character) {
	color.HiGreen("* %s uses Lightning Bolt *\n", self.Stats.Name)
	damage := int(float64(self.Stats.Intelligence) * 2.9)
	color.Magenta("%s %s deals %d magic damage\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)

	other.Stats.Health -= damage
}

// Stun
func (self *Character) Stun(other *Character) {
	color.HiGreen("* %s uses Stun *\n", self.Stats.Name)
	other.Status.Stunned +=3
}

