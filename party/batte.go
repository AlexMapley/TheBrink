package party

import (
	"math/rand"
	"time"
	"the_brink/characters"

	"github.com/fatih/color"
  )

// TargetMember returns an attackble player from the party
func (party *Party) TargetMember() *characters.Character {

	rand.Seed(time.Now().Unix())

	for i := 0; i < len(party.Members); i++ {
		randomPartyMember := party.Members[rand.Intn(len(party.Members))]
		if randomPartyMember.Stats.Health > 0 {
			return randomPartyMember
		}
	}
	return nil
}


// Battle other party
func (selfParty *Party) Battle(otherParty *Party) {

	round := 1
	for selfParty.GetHealth() > 0 && otherParty.GetHealth() > 0 {
		time.Sleep(150 * time.Millisecond)
		color.Green("\n\nRound %d", round)

		for _, member := range selfParty.Members {
			if member.Stats.Health <= 0 {
				continue
			}
			target := otherParty.TargetMember()
			if target == nil {
				continue
			}

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
		
		for _, member := range otherParty.Members {
			if member.Stats.Health <= 0 {
				continue
			}
			target := selfParty.TargetMember()
			if target == nil {
				continue
			}

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
		round++
	}

	if selfParty.GetHealth() > 0 {
		for _, otherMember := range otherParty.Members {
			for _, selfMember := range selfParty.Members {
				selfMember.Stats.XP += otherMember.Stats.XP
			}
		}
		color.Cyan("Player Party wins the battle\n")
		return
	}

	color.Cyan("Enemy Party wins the battle\n")
}