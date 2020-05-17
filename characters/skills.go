package characters

import (
	"time"
	"math/rand"

	"github.com/fatih/color"
)

// BasicAttack
func (self *Character) BasicAttack(other *Character) {
	rand.Seed(time.Now().UnixNano())

	// Events
	dodged := false
	critical := false
	blocked := false
	

	// Generate base multipliers
	damage := self.Stats.Strength + (self.Stats.Agility/2)

	dodgeThreshold := 200 + (self.Stats.AccuracyRating() * 2)
	criticalThreshold := 200
	blockThreshold := 150 + (self.Stats.Strength * 2)

	// Dodge Chance
	if (other.Stats.DodgeValue() >= rand.Intn(dodgeThreshold)) {
		dodged = true
		damage = 0
	}
	// Block Chance
	if (self.Stats.BlockValue() >= rand.Intn(blockThreshold)) {
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


func (self *Character) DoubleStrike(other *Character) {
	color.HiGreen("* %s uses Double Strike *\n", self.Stats.Name)
	self.BasicAttack(other)
	self.BasicAttack(other)
}

func (self *Character) LightningBolt(other *Character) {
	color.HiGreen("* %s uses LightningBolt *\n", self.Stats.Name)
	damage := int(float64(self.Stats.Intelligence) * 3.5)

	color.Magenta("%s %s deals %d magic damage\n", self.Stats.Name, self.Stats.DisplayHealth(), damage)
}