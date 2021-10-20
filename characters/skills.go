package characters

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Attack
func (self *Character) Attack(other *Character, base int) {
	rand.Seed(time.Now().UnixNano())

	// Generate base multipliers
	damage := float64(base)

	// Dodge Threshold
	dodgeThreshold := 235 + int(self.Stats.AccuracyRating() * 1.75)
	dodged := other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold)))
	if dodged {
		damage = 0

		color.Yellow("%s %s deals %s 0 damage (%s Dodged)\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, other.Stats.Name)
		return
	}

	// Block Threshold
	blockThreshold := 180 + (self.Stats.Strength * 2) + int(self.Stats.AccuracyRating())
	blocked := other.Stats.BlockValue() >= float64(rand.Intn(int(blockThreshold)))
	if blocked {
		blocked = true
		damage /= 2.0
	}

	// Critical Threshold
	criticalThreshold := 225 + (2 * self.Stats.Level)
	critical := self.Stats.CriticalValue() >= float64(rand.Intn(int(criticalThreshold)))
	if critical {
		critical = true
		damage *= 2.0
	}

	// Damage Multiplier
	DamangeMultiplier := rand.Intn(1200) + rand.Intn(800)
	damage = damage * float64(DamangeMultiplier) / 1000

	// Event Cases
	if blocked {
		color.Cyan("%s %s deals %s %.2f damage (%s Blocked)\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage, other.Stats.Name)
	} else if critical {
		color.HiRed("%s %s deals %s %.2f damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage)
	} else {
		color.White("%s %s deals %s %.2f damage\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage)
	}
	other.Stats.Health -= damage
}

// Bark
func (self *Character) Bark(other *Character) {
	color.HiGreen("* %s barks at %s, and stuns for 1 turn *\n", self.Stats.Name, other.Stats.Name)
	self.Attack(other, self.Stats.Intelligence)

	self.Stun(other, 1) 
}

// DoubleStrike
func (self *Character) DoubleStrike(other *Character) {
	color.HiGreen("* %s hits twice at %s *\n", self.Stats.Name,other.Stats.Name)
	self.Attack(other, self.Stats.Strength+(self.Stats.Agility/2))
	self.Attack(other, self.Stats.Strength+(self.Stats.Agility/2))
}

// Flash Heal
func (self *Character) FlashHeal() {
	color.HiGreen("* %s uses Flash Heal *\n", self.Stats.Name)

	heal := (1.1 * float64(self.Stats.Intelligence)) + (0.25 * float64(self.Stats.Vitality))

	// Critical Chance
	criticalThreshold := 140
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

// Icicle
func (self *Character) Icicle(other *Character) {
	color.HiGreen("* %s shoots an icicle at %s *\n", self.Stats.Name,other.Stats.Name)
	damage := float64(self.Stats.Intelligence) * 1.7

	dodgeThreshold := float64(220) + (self.Stats.AccuracyRating() * 2)
	// Dodge Chance
	if other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold))) {
		damage = 0
		color.Yellow("%s %s deals %s %.2f damage (%s Dodged)\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage, other.Stats.Name)
	} else {
		color.Magenta("%s %s deals %s %.2f magic damage, and stuns for 2 turns\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage)
	}

	other.Stats.Health -= damage
	self.Stun(other, 2)
}

// Knock The Wind Out
func (self *Character) KnockTheWindOut(other *Character) {
	color.HiGreen("* %s knocks the wind out of %s, stunning for 2 rounds *\n", self.Stats.Name, other.Stats.Name)
	self.Attack(other, self.Stats.Strength)
	self.Stun(other, 2)
}

// LightningBolt
func (self *Character) LightningBolt(other *Character) {
	color.HiGreen("* %s rips some lightning at %s *\n", self.Stats.Name,other.Stats.Name)
	damage := float64(self.Stats.Intelligence) * 2.1

	// Dodge Chance
	dodgeThreshold := float64(220) + (self.Stats.AccuracyRating() * 2)
	dodged := other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold)))
	if dodged {
		damage = 0
		color.Yellow("%s %s deals %s %.2f damage (%s Dodged)\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage, other.Stats.Name)
		return
	}

	// Critical Chance
	criticalThreshold := 160
	critical := int(self.Stats.CriticalValue()) >= rand.Intn(criticalThreshold)
	if critical && !dodged {
		critical = true
		damage *= 2
	}

	if critical {
		color.HiRed("%s %s deals %s %.2f damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage)
	} else {
		color.Magenta("%s %s deals %s %.2f magic damage\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage)
	}
	other.Stats.Health -= damage
}

// Smite
func (self *Character) Smite(other *Character) {
	color.HiGreen("* %s smites %s *\n", self.Stats.Name, other.Stats.Name)
	damage := 1.5 * float64(self.Stats.Intelligence) + 0.5 * float64(self.Stats.Vitality)

	// Critical Chance
	criticalThreshold := 160
	critical := int(self.Stats.CriticalValue()) >= rand.Intn(criticalThreshold)
	if critical {
		critical = true
		damage *= 1.8
	}

	if critical {
		color.HiRed("%s %s deals %s %.2f damage (Critical)\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage)
	} else {
		color.Magenta("%s %s deals %s %.2f magic damage\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name, damage)
	}
	other.Stats.Health -= damage
}

// Slash
func (self *Character) Slash(other *Character) {
	color.HiGreen("* %s slashes %s *\n", self.Stats.Name, other.Stats.Name)
	self.Attack(other, (self.Stats.Strength*3)+(self.Stats.Agility*2))
}

// Sneak Attack
func (self *Character) SneakAttack(other *Character) {
	color.HiGreen("* %s uses SneakAttack *\n", self.Stats.Name)
	self.Attack(other, (self.Stats.Strength/2)+(self.Stats.Agility*3)+(self.Stats.Intelligence*2))

	color.Magenta("%s %s stuns %s for 1 turn\n", self.Stats.Name, self.Stats.DisplayHealth(), other.Stats.Name)

	self.Stun(other, 1)
}
