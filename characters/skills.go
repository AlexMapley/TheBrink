package characters

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// BasicAttack
func (self *Character) BasicAttack(other *Character, base int) {
	rand.Seed(time.Now().UnixNano())

	// Generate base multipliers
	damage := float64(base)

	// Dodge Threshold
	dodgeThreshold := 220 + int(self.Stats.AccuracyRating() * 2)
	dodged := other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold)))
	if dodged {
		dodged = true
		damage = 0
	}

	// Block Threshold
	blockThreshold := 180 + (self.Stats.Strength * 2) + int(self.Stats.AccuracyRating())
	blocked := other.Stats.BlockValue() >= float64(rand.Intn(int(blockThreshold)))
	if blocked {
		blocked = true
		damage /= 2.0
	}

	// Critical Threshold
	criticalThreshold := 225
	critical := self.Stats.CriticalValue() >= float64(rand.Intn(int(criticalThreshold)))
	if critical {
		critical = true
		damage *= 2.0
	}

	// Damage Multiplier
	DamangeMultiplier := rand.Intn(1200) + rand.Intn(800)
	damage = damage * float64(DamangeMultiplier) / 1000

	// Event Cases
	if dodged {
		color.Yellow("%s %s deals %.2f damage (Dodge)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else if blocked {
		color.Cyan("%s %s deals %.2f damage (Blocked)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else if critical {
		color.HiRed("%s %s deals %.2f damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else {
		color.White("%s %s deals %.2f damage\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	}
	other.Stats.Health -= damage
}

// DoubleStrike
func (self *Character) DoubleStrike(other *Character) {
	color.HiGreen("* %s uses Double Strike *\n", self.Stats.Name)
	self.BasicAttack(other, self.Stats.Strength+(self.Stats.Agility/2))
	self.BasicAttack(other, self.Stats.Strength+(self.Stats.Agility/2))
}

// Rend
func (self *Character) Rend(other *Character) {
	color.HiGreen("* %s uses Rend *\n", self.Stats.Name)
	self.BasicAttack(other, (self.Stats.Strength*3)+(self.Stats.Agility*2))
}

// GhostBlade
func (self *Character) GhostBlade(other *Character) {
	color.HiGreen("* %s uses GhostBlade *\n", self.Stats.Name)
	self.BasicAttack(other, (self.Stats.Strength/2)+(self.Stats.Agility*3)+(self.Stats.Intelligence*2))

	color.Magenta("%s %s stuns for 1 turn\n", self.Stats.Name, self.Stats.DisplayHealth())

	other.Status.Stunned += 1
}

// Heal
func (self *Character) Heal() {
	color.HiGreen("* %s uses Heal *\n", self.Stats.Name)

	heal := (2.0 * float64(self.Stats.Intelligence)) + (0.3 * float64(self.Stats.Vitality))

	// Critical Chance
	criticalThreshold := 170
	if int(self.Stats.CriticalValue()) >= rand.Intn(int(criticalThreshold)) {
		heal *= 2.0
		color.HiMagenta("%s %s Heals %.2f damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), heal)
	} else {
		color.Magenta("%s %s Heals %.2f damage\n", self.Stats.Name, self.Stats.DisplayHealth(), heal)

	}

	self.Stats.Health += heal
	if self.Stats.Health > self.Stats.MaxHealth() {
		self.Stats.Health = self.Stats.MaxHealth()
	}
}

// IceBlast
func (self *Character) IceBlast(other *Character) {
	color.HiGreen("* %s uses Ice Blast *\n", self.Stats.Name)
	damage := float64(self.Stats.Intelligence) * 1.7

	dodgeThreshold := float64(220) + (self.Stats.AccuracyRating() * 2)
	// Dodge Chance
	if other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold))) {
		damage = 0
		color.Yellow("%s %s deals %.2f damage (Dodge)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else {
		color.Magenta("%s %s deals %.2f magic damage, and stuns for 2 turns\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	}

	other.Stats.Health -= damage
	other.Status.Stunned += 2
}

// LightningBolt
func (self *Character) LightningBolt(other *Character) {
	color.HiGreen("* %s uses Lightning Bolt *\n", self.Stats.Name)
	damage := float64(self.Stats.Intelligence) * 2.1

	// Dodge Chance
	dodgeThreshold := float64(220) + (self.Stats.AccuracyRating() * 2)
	dodged := other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold)))
	if dodged {
		damage = 0
		color.Yellow("%s %s deals %.2f damage (Dodge)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	}

	// Critical Chance
	criticalThreshold := 160
	critical := int(self.Stats.CriticalValue()) >= rand.Intn(criticalThreshold)
	if critical && !dodged {
		critical = true
		damage *= 2
	}

	if critical {
		color.HiRed("%s %s deals %d damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	} else {
		color.Magenta("%s %s deals %.2f magic damage\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
	}
	other.Stats.Health -= damage
}

// Stun
func (self *Character) Stun(other *Character) {
	color.HiGreen("* %s uses Stun *\n", self.Stats.Name)
	other.Status.Stunned += 3
}
