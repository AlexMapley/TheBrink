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

// Act will perform a skill
// against onself or an option target,
// lowering cooldowns and status effects
func (character *Character) Act(target *Character) {
	if character.Status.Stunned == 0 {
		chosenSkill := character.ChooseSkill()
		switch chosenSkill.Name {
		case "Bark":
			character.Bark(target)
		case "Double Strike":
			character.DoubleStrike(target)
		case "Flash Heal":
			character.FlashHeal()
		case "Knock The Wind Out":
			character.KnockTheWindOut(target)
		case "Heal":
			character.Heal()
		case "Icicle":
			character.Icicle(target)
		case "Lightning Bolt":
			character.LightningBolt(target)
		case "Slash":
			character.Slash(target)
		case "Smite":
			character.Smite(target)
		case "Sneak Attack":
			character.SneakAttack(target)
		default:
			character.Attack(target, character.Stats.Strength+(character.Stats.Agility/2))
		}
		// self cooldowns
		for i, skill := range character.SkillSlots {
			if skill.Name == chosenSkill.Name {
				character.SkillSlots[i].CoolDown = skill.CoolDownMax
			}
			if skill.CoolDown > 0 {
				character.SkillSlots[i].CoolDown--
			}
		}
		if target.Stats.Health <= 0 {
			color.HiRed("%s Dies", target.Stats.Name)
		}
	}
	character.Status.Stunned--
}

// Duel will update both characters health
func (self *Character) Duel(other *Character) {
	self.Stats.Display()
	other.Stats.Display()

	for self.Stats.Health > 0 && other.Stats.Health > 0 {

		time.Sleep(100 * time.Millisecond)

		// self action
		self.Act(other)

		// other action
		other.Act(self)
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
