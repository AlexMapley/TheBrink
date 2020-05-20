package characters

import (
	"time"

	"github.com/fatih/color"
)

// ChooseSkill returns the highest cooldown
// skill that is currently available
// and castable
func (self *Character) ChooseSkill() Skill {

	selectedSkill := self.SkillSlots[0]

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

// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	self.Stats.Display()
	other.Stats.Display()

	for (self.Stats.Health > 0 && other.Stats.Health  > 0) {

		time.Sleep(100 * time.Millisecond)

		// self action
		if self.Status.Stunned == 0 {
			chosenSkill := self.ChooseSkill()
			switch chosenSkill.Name {
			case "DoubleStrike":
				self.DoubleStrike(other)
			case "LightningBolt":
				self.LightningBolt(other)
			case "IceBlast":
				self.IceBlast(other)
			case "Stun":
				self.Stun(other)
			case "Heal":
				self.Heal()
			default:
				self.BasicAttack(other)
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
		if (other.Stats.Health  > 0) {
			if (other.Status.Stunned  == 0) {
				chosenSkill := other.ChooseSkill()
				switch chosenSkill.Name {
				case "DoubleStrike":
					other.DoubleStrike(self)
				case "LightningBolt":
					other.LightningBolt(self)
				case "Stun":
					other.Stun(self)
				case "Heal":
					other.Heal()
				default:
					other.BasicAttack(self)
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
				other.Status.Stunned--
			}
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

	// Reset Resource Pools
	self.Stats.Health = self.Stats.MaxHealth()
	self.Stats.Focus = self.Stats.MaxFocus()

	// Reset Skill Cooldowns
	for i := range self.SkillSlots {
		self.SkillSlots[i].CoolDown = 0
	}
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
