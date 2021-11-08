package characters

import (
	"time"

	"github.com/fatih/color"
)

// ChooseSkill returns the highest cooldown
// skill that is currently available
// and castable
func (character *Character) ChooseSkill() Skill {

	selectedSkill := Skill{}

	for _, skill := range character.SkillSlots {

		// check if skill is on cooldown
		if skill.CoolDown > 0 {
			continue
		}

		// check if we have the focus to cast
		if skill.Cost > character.Stats.Focus {
			continue
		}

		// check if skill is higher preference
		if skill.CoolDownMax > selectedSkill.CoolDownMax {
			selectedSkill = skill
		}
	}

	character.Stats.Focus -= selectedSkill.Cost
	
	// don't want to go negative focus
	if character.Stats.Focus < 0 {
		character.Stats.Focus = 0
	}

	return selectedSkill
}

// Act will perform a skill
// against onself or an option target,
// lowering cooldowns and status effects
func (character *Character) Act(target *Character) {

	var chosenSkill Skill
	if character.Status.Stunned <= 0 {
		chosenSkill = character.ChooseSkill()
		switch chosenSkill.Name {
		case "Bark":
			character.Bark(target)
		case "Double Strike":
			character.DoubleStrike(target)
		case "Flash Heal":
			character.FlashHeal(target)
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
		case "Uppercut":
			character.Uppercut(target)
		default:
			character.Attack(target, character.Stats.Strength+(character.Stats.Agility/2))
		}
		if target.Stats.Health <= 0 {
			color.HiRed("%s Dies", target.Stats.Name)
		}
	}
	// self cooldowns
	character.Status.Stunned--
	for i, skill := range character.SkillSlots {
		if skill.Name == chosenSkill.Name {
			character.SkillSlots[i].CoolDown = skill.CoolDownMax
		}
		if skill.CoolDown > 0 {
			character.SkillSlots[i].CoolDown--
		}
	}
}

// Duel will update both characters health
func (character *Character) Duel(other *Character) {
	character.Stats.Display()
	other.Stats.Display()

	for character.Stats.Health > 0 && other.Stats.Health > 0 {

		time.Sleep(100 * time.Millisecond)

		// self action
		character.Act(other)

		// other action
		other.Act(character)
	}

	if character.Stats.Health >= other.Stats.Health {
		color.Cyan("\n%s Wins the duel\n", character.Stats.Name)
		color.Red("\nOther xp is %d\n", other.Stats.XP)
		character.Stats.XP += other.Stats.XP
		other.Stats.XP = 0

		return
	}

	color.Cyan("%s Wins the duel\n", other.Stats.Name)

}

// Rest
func (character *Character) Rest() {

	// Reset Resource Pools
	character.Stats.Health = character.Stats.MaxHealth()
	character.Stats.Focus = character.Stats.MaxFocus()

	// Reset Skill Cooldowns
	for i := range character.SkillSlots {
		character.SkillSlots[i].CoolDown = character.SkillSlots[i].CoolDownInitial
	}
}

// Level Up
func (character *Character) LevelUp() bool {
	if character.Stats.XP < 1000 {
		return false
	}

	// increase level
	character.Stats.Level++

	// increase core stats
	character.Stats.Vitality += character.Stats.LevelBonuses.Vitality
	character.Stats.Strength += character.Stats.LevelBonuses.Strength
	character.Stats.Agility += character.Stats.LevelBonuses.Agility
	character.Stats.Intelligence += character.Stats.LevelBonuses.Intelligence

	// increase secondary stats
	character.Stats.Expertise += character.Stats.LevelBonuses.Expertise
	character.Stats.Block += character.Stats.LevelBonuses.Block
	character.Stats.Critical += character.Stats.LevelBonuses.Critical
	character.Stats.Dodge += character.Stats.LevelBonuses.Dodge

	// Dock XP
	character.Stats.XP -= 1000

	// rest
	character.Rest()

	return true
}
