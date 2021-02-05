package characters

import (
	"time"

	"github.com/fatih/color"
)

// ChooseSkill returns the highest cooldown
// skill that is currently available
// and castable
func (self *Character) ChooseSkill() Skill {

	selectedSkill := Skill{}

	for _, skill := range self.SkillSlots {

		// check if skill is on cooldown
		if skill.CoolDown > 0 {
			continue
		}

		// check if we have the focus to cast
		if skill.Cost > self.Stats.Focus {
			continue
		}

		// check if skill is higher preference
		if skill.CoolDownMax > selectedSkill.CoolDownMax {
			selectedSkill = skill
		}
	}

	self.Stats.Focus -= selectedSkill.Cost
	return selectedSkill
}

func (self *Character) DoSkill(target *Character) {
	if member.Status.Stunned == 0 {
		chosenSkill := member.ChooseSkill()
		switch chosenSkill.Name {
		case "Bark":
			member.Bark(target)
		case "Double Strike":
			member.DoubleStrike(target)
		case "Flash Heal":
			member.FlashHeal()
		case "Sneak Attack":
			member.SneakAttack(target)
		case "Heal":
			member.Heal()
		case "Icicle":
			member.Icicle(target)
		case "Lightning Bolt":
			member.LightningBolt(target)
		case "Slash":
			member.Slash(target)
		case "Smite":
			member.Smite(target)
		case "Stun":
			member.Stun(target)
		default:
			member.Attack(target, member.Stats.Strength+(member.Stats.Agility/2))
		}
		// self cooldowns
		for i, skill := range member.SkillSlots {
			if skill.Name == chosenSkill.Name {
				member.SkillSlots[i].CoolDown = skill.CoolDownMax
			}
			if skill.CoolDown > 0 {
				member.SkillSlots[i].CoolDown--
			}
		}
		if target.Stats.Health <= 0 {
			color.HiRed("%s Dies", target.Stats.Name)
		}
		continue
	}
	member.Status.Stunned--
}

// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	self.Stats.Display()
	other.Stats.Display()

	for self.Stats.Health > 0 && other.Stats.Health > 0 {

		time.Sleep(100 * time.Millisecond)

		// self action
		if self.Status.Stunned == 0 {
			chosenSkill := self.ChooseSkill()
			switch chosenSkill.Name {
			case "Bark":
				self.Bark(other)
			case "Double Strike":
				self.DoubleStrike(other)
			case "Flash Heal":
				self.FlashHeal()
			case "Heal":
				self.Heal()
			case "Knock The Wind Out":
				self.KnockTheWindOut(other)
			case "Icicle":
				self.Icicle(other)
			case "Lightning Bolt":
				self.LightningBolt(other)
			case "Slash":
				self.Slash(other)
			case "Smite":
				self.Smite(other)
			case "Sneak Attack":
				self.SneakAttack(other)
			default:
				self.Attack(other, self.Stats.Strength+(self.Stats.Agility/2))
			}
			// self cooldowns
			for i, skill := range self.SkillSlots {
				if skill.Name == chosenSkill.Name {
					self.SkillSlots[i].CoolDown = skill.CoolDownMax
				}
				if skill.CoolDown > 0 {
					self.SkillSlots[i].CoolDown--
				}
			}
		} else {
			self.Status.Stunned--
		}

		// other action
		if other.Stats.Health > 0 {
			if other.Status.Stunned == 0 {
				chosenSkill := other.ChooseSkill()
				switch chosenSkill.Name {
				case "Bark":
					other.Bark(self)
				case "Double Strike":
					other.DoubleStrike(self)
				case "Flash Heal":
					other.FlashHeal()
				case "Heal":
					other.Heal()
				case "Knock The Wind Out":
					other.KnockTheWindOut(self)
				case "Icicle":
					other.Icicle(self)
				case "Lightning Bolt":
					other.LightningBolt(self)
				case "Slash":
					other.Slash(self)
				case "Smite":
					other.Smite(self)
				case "Sneak Attack":
					other.SneakAttack(self)
				default:
					other.Attack(self, other.Stats.Strength+(other.Stats.Agility/2))
				}
				// other cooldowns
				for i, skill := range other.SkillSlots {
					if skill.Name == chosenSkill.Name {
						other.SkillSlots[i].CoolDown = skill.CoolDownMax
					}
					if skill.CoolDown > 0 {
						other.SkillSlots[i].CoolDown--
					}
				}
			} else {
				// Recover from round of stun
				other.Status.Stunned--
			}
		}

	}

	if self.Stats.Health >= other.Stats.Health {
		color.Cyan("\n%s Wins the duel\n", self.Stats.Name)
		color.Red("\nOther xp is %d\n", other.Stats.XP)
		self.Stats.XP += other.Stats.XP
		other.Stats.XP = 0

		return
	}

	color.Cyan("%s Wins the duel\n", other.Stats.Name)

}

// Rest
func (self *Character) Rest() {

	// Reset Resource Pools
	self.Stats.Health = self.Stats.MaxHealth()
	self.Stats.Focus = self.Stats.MaxFocus()

	// Reset Skill Cooldowns
	for i := range self.SkillSlots {
		self.SkillSlots[i].CoolDown = self.SkillSlots[i].CoolDownInitial
	}
}

// Level Up
func (self *Character) LevelUp() bool {
	if self.Stats.XP < 1000 {
		return false
	}

	// increase level
	self.Stats.Level++

	// increase core stats
	self.Stats.Vitality += self.Stats.LevelBonuses.Vitality
	self.Stats.Strength += self.Stats.LevelBonuses.Strength
	self.Stats.Agility += self.Stats.LevelBonuses.Agility
	self.Stats.Intelligence += self.Stats.LevelBonuses.Intelligence

	// increase secondary stats
	self.Stats.Expertise += self.Stats.LevelBonuses.Expertise
	self.Stats.Block += self.Stats.LevelBonuses.Block
	self.Stats.Critical += self.Stats.LevelBonuses.Critical
	self.Stats.Dodge += self.Stats.LevelBonuses.Dodge

	// Dock XP
	self.Stats.XP -= 1000

	// rest
	self.Rest()

	return true
}
