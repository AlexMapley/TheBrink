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

	dodgeThreshold := 220 + (self.Stats.AccuracyRating() * 2)
	criticalThreshold := 220
	blockThreshold := 180 + (self.Stats.Strength * 2) + (self.Stats.AccuracyRating())

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
	DamangeMultiplier := rand.Intn(70) + 30
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
	self.BasicAttack(other, self.Stats.Strength + (self.Stats.Agility/2))
	self.BasicAttack(other, self.Stats.Strength + (self.Stats.Agility/2))
}

// Rend
func (self *Character) Rend(other *Character) {
	color.HiGreen("* %s uses Rend *\n", self.Stats.Name)
	self.BasicAttack(other, (self.Stats.Strength*3) + (self.Stats.Agility*2))
}

// GhostBlade
func (self *Character) GhostBlade(other *Character) {
	color.HiGreen("* %s uses GhostBlade *\n", self.Stats.Name)
	self.BasicAttack(other, (self.Stats.Strength/2) + (self.Stats.Agility*3) + (self.Stats.Intelligence*2))

	color.Magenta("%s %s stuns for 1 turn\n", self.Stats.Name, self.Stats.DisplayHealth())

	other.Status.Stunned +=1
}

// Heal
func (self *Character) Heal() {
	color.HiGreen("* %s uses Heal *\n", self.Stats.Name)

	heal := int(float64(self.Stats.Intelligence) * 2.0 + float64(self.Stats.Vitality) * 0.3)

	// Critical Chance
	criticalThreshold := 170
	if (self.Stats.CriticalValue() >= rand.Intn(criticalThreshold)) {
		critical = true
		heal = damage * 2
		color.HiMagenta("%s %s Heals %d damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), heal)
	} else {
		color.Magenta("%s %s Heals %d damage\n", self.Stats.Name, self.Stats.DisplayHealth(), heal)

	}

	self.Stats.Health += heal
	if self.Stats.Health > self.Stats.MaxHealth() {
		self.Stats.Health = self.Stats.MaxHealth()
	}
}

// IceBlast
func (self *Character) IceBlast(other *Character) {
	color.HiGreen("* %s uses Ice Blast *\n", self.Stats.Name)
	damage := int(float64(self.Stats.Intelligence) * 1.7)

	dodgeThreshold := 220 + (self.Stats.AccuracyRating() * 2)
	// Dodge Chance
	if (other.Stats.DodgeValue() >= rand.Intn(dodgeThreshold)) {
		damage = 0
		color.Yellow("%s %s deals %d damage (Dodge)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else {
		color.Magenta("%s %s deals %d magic damage, and stuns for 2 turns\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	}

	other.Stats.Health -= damage
	other.Status.Stunned +=2
}

// LightningBolt
func (self *Character) LightningBolt(other *Character) {
	color.HiGreen("* %s uses Lightning Bolt *\n", self.Stats.Name)
	damage := int(float64(self.Stats.Intelligence) * 2.7)

	dodgeThreshold := 220 + (self.Stats.AccuracyRating() * 2)
	// Dodge Chance
	if (other.Stats.DodgeValue() >= rand.Intn(dodgeThreshold)) {
		damage = 0
		color.Yellow("%s %s deals %d damage (Dodge)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else {
		color.Magenta("%s %s deals %d magic damage\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	}

	other.Stats.Health -= damage
}

// Stun
func (self *Character) Stun(other *Character) {
	color.HiGreen("* %s uses Stun *\n", self.Stats.Name)
	other.Status.Stunned +=3
}

