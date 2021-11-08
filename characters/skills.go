package characters

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Attack
func (character *Character) Attack(other *Character, base int) {
	rand.Seed(time.Now().UnixNano())

	// Generate base multipliers
	damage := float64(base)

	// Dodge Threshold
	dodgeThreshold := 235 + int(character.Stats.AccuracyRating() * 1.75)
	dodged := other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold)))
	if dodged {
		damage = 0

		color.Yellow("%s's attack %s deals %s 0 damage (%s Dodges)\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, other.Stats.Name)
		return
	}

	// Block Threshold
	blockThreshold := 180 + (character.Stats.Strength * 2) + int(character.Stats.AccuracyRating())
	blocked := other.Stats.BlockValue() >= float64(rand.Intn(int(blockThreshold)))
	if blocked {
		blocked = true
		damage /= 2.0
	}

	// Critical Threshold
	criticalThreshold := 225 + (2 * character.Stats.Level)
	critical := character.Stats.CriticalValue() >= float64(rand.Intn(int(criticalThreshold)))
	if critical {
		critical = true
		damage *= 2.0
	}

	// Damage Multiplier
	DamangeMultiplier := rand.Intn(1200) + rand.Intn(800)
	damage = damage * float64(DamangeMultiplier) / 1000

	// Event Cases
	switch {
	case (blocked && critical):
		color.Blue("%s's attack %s deals %s %.2f damage (%s Criticals, %s Blocks)\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage, character.Stats.Name, other.Stats.Name)
	case blocked:
		color.Cyan("%s's attack %s deals %s %.2f damage (%s Blocks)\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage, other.Stats.Name)
	case critical:
		color.HiRed("%s's attack %s deals %s %.2f damage (%s Criticals)\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage, character.Stats.Name)
	default:
		color.White("%s's attack %s deals %s %.2f damage\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage)
	}

	other.Stats.Health -= damage
}

// Bark
func (character *Character) Bark(other *Character) {
	color.HiGreen("* %s barks at %s, and stuns for 1 turn *\n", character.Stats.Name, other.Stats.Name)
	character.Attack(other, character.Stats.Intelligence)

	character.Stun(other, 1) 
}

// DoubleStrike
func (character *Character) DoubleStrike(other *Character) {
	color.HiGreen("* %s hits twice at %s *\n", character.Stats.Name,other.Stats.Name)
	character.Attack(other, character.Stats.Strength+(character.Stats.Agility/2))
	character.Attack(other, character.Stats.Strength+(character.Stats.Agility/2))
}

// Flash Heal
func (character *Character) FlashHeal() {
	color.HiGreen("* %s uses Flash Heal *\n", character.Stats.Name)

	heal := (1.1 * float64(character.Stats.Intelligence)) + (0.25 * float64(character.Stats.Vitality))

	// Critical Chance
	criticalThreshold := 140
	if int(character.Stats.CriticalValue()) >= rand.Intn(int(criticalThreshold)) {
		heal *= 2.0
		color.HiMagenta("%s %s Heals %.2f damage (Critical)\n", character.Stats.Name, character.Stats.DisplayHealth(), heal)
	} else {
		color.Magenta("%s %s Heals %.2f damage\n", character.Stats.Name, character.Stats.DisplayHealth(), heal)

	}

	character.Stats.Health += heal
	if character.Stats.Health > character.Stats.MaxHealth() {
		character.Stats.Health = character.Stats.MaxHealth()
	}
}

// Heal
func (character *Character) Heal() {
	color.HiGreen("* %s uses Heal *\n", character.Stats.Name)

	heal := (2.0 * float64(character.Stats.Intelligence)) + (0.3 * float64(character.Stats.Vitality))

	// Critical Chance
	criticalThreshold := 170
	if int(character.Stats.CriticalValue()) >= rand.Intn(int(criticalThreshold)) {
		heal *= 2.0
		color.HiMagenta("%s %s Heals %.2f damage (Critical)\n", character.Stats.Name, character.Stats.DisplayHealth(), heal)
	} else {
		color.Magenta("%s %s Heals %.2f damage\n", character.Stats.Name, character.Stats.DisplayHealth(), heal)

	}

	character.Stats.Health += heal
	if character.Stats.Health > character.Stats.MaxHealth() {
		character.Stats.Health = character.Stats.MaxHealth()
	}
}

// Icicle
func (character *Character) Icicle(other *Character) {
	color.HiGreen("* %s shoots an icicle at %s *\n", character.Stats.Name,other.Stats.Name)
	damage := float64(character.Stats.Intelligence) * 1.7

	dodgeThreshold := float64(220) + (character.Stats.AccuracyRating() * 2)
	// Dodge Chance
	if other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold))) {
		damage = 0
		color.Yellow("%s %s deals %s %.2f damage (%s Dodged)\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage, other.Stats.Name)
	} else {
		color.Magenta("%s %s deals %s %.2f magic damage, and stuns for 2 turns\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage)
	}

	other.Stats.Health -= damage
	character.Stun(other, 2)
}

// Uppercut
func (character *Character) Uppercut(other *Character) {
	color.HiGreen("* %s knocks the wind out of %s, stunning for 2 rounds *\n", character.Stats.Name, other.Stats.Name)
	character.Attack(other, character.Stats.Strength)
	character.Stun(other, 2)
}

// LightningBolt
func (character *Character) LightningBolt(other *Character) {
	color.HiGreen("* %s rips some lightning at %s *\n", character.Stats.Name,other.Stats.Name)
	damage := float64(character.Stats.Intelligence) * 2.1

	// Dodge Chance
	dodgeThreshold := float64(220) + (character.Stats.AccuracyRating() * 2)
	dodged := other.Stats.DodgeValue() >= float64(rand.Intn(int(dodgeThreshold)))
	if dodged {
		damage = 0
		color.Yellow("%s %s deals %s %.2f damage (%s Dodged)\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage, other.Stats.Name)
		return
	}

	// Critical Chance
	criticalThreshold := 160
	critical := int(character.Stats.CriticalValue()) >= rand.Intn(criticalThreshold)
	if critical && !dodged {
		critical = true
		damage *= 2
	}

	if critical {
		color.HiRed("%s %s deals %s %.2f damage (Critical)\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage)
	} else {
		color.Magenta("%s %s deals %s %.2f magic damage\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage)
	}
	other.Stats.Health -= damage
}

// Smite
func (character *Character) Smite(other *Character) {
	color.HiGreen("* %s smites %s *\n", character.Stats.Name, other.Stats.Name)
	damage := 1.5 * float64(character.Stats.Intelligence) + 0.5 * float64(character.Stats.Vitality)

	// Critical Chance
	criticalThreshold := 160
	critical := int(character.Stats.CriticalValue()) >= rand.Intn(criticalThreshold)
	if critical {
		critical = true
		damage *= 1.8
	}

	if critical {
		color.HiRed("%s %s deals %s %.2f damage (Critical)\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage)
	} else {
		color.Magenta("%s %s deals %s %.2f magic damage\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name, damage)
	}
	other.Stats.Health -= damage
}

// Slash
func (character *Character) Slash(other *Character) {
	color.HiGreen("* %s slashes %s *\n", character.Stats.Name, other.Stats.Name)
	character.Attack(other, (character.Stats.Strength*3)+(character.Stats.Agility*2))
}

// Sneak Attack
func (character *Character) SneakAttack(other *Character) {
	color.HiGreen("* %s uses SneakAttack *\n", character.Stats.Name)
	character.Attack(other, (character.Stats.Strength/2)+(character.Stats.Agility*3)+(character.Stats.Intelligence*2))

	color.Magenta("%s %s stuns %s for 1 turn\n", character.Stats.Name, character.Stats.DisplayHealth(), other.Stats.Name)

	character.Stun(other, 1)
}
