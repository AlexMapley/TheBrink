package characters

import (
	"time"

	"github.com/fatih/color"
)

// ChooseSkill returns the highest cooldown
// skill that is currently available
// and castable
func (self *Character) ChooseSkill() Skill {

	selectedSkill := self.Stats.SkillSlots[0]

	for _, skill := range self.Stats.SkillSlots {

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

	return selectedSkill
}

// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	self.Stats.Display()
	other.Stats.Display()

	for (self.Stats.Health > 0 && other.Stats.Health  > 0) {

		time.Sleep(100 * time.Millisecond)

		// self action
		chosenSkill := self.ChooseSkill()
		switch chosenSkill.Name {
		case "BasicAttack":
			self.BasicAttack(other)
		case "DoubleStrike":
			self.DoubleStrike(other)
		}
		case "LightningBolt":
			self.LightningBolt(other)
		}
		// self cooldowns
		for i, skill := range self.Stats.SkillSlots {
			if skill.Name == chosenSkill.Name {
				self.Stats.SkillSlots[i].CoolDown = skill.CoolDownMax
			}
			if skill.CoolDown > 0 {
				self.Stats.SkillSlots[i].CoolDown--
			}
		}

		// other action
		if (other.Stats.Health  > 0) {
			chosenSkill := other.ChooseSkill()
			switch chosenSkill.Name {
			case "BasicAttack":
				other.BasicAttack(self)
			}
			// other cooldowns
			for i, skill := range other.Stats.SkillSlots {
				if skill.Name == chosenSkill.Name {
					other.Stats.SkillSlots[i].CoolDown = skill.CoolDownMax
				}
				if skill.CoolDown > 0 {
					other.Stats.SkillSlots[i].CoolDown--
				}
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
	for i := range self.Stats.SkillSlots {
		self.Stats.SkillSlots[i].CoolDown = 0
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
