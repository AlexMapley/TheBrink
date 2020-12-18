package party

// Level Up Party
func (self *Party) LevelUp() bool{
	success := false
	for _, selfPartyMember := range self.Members {
		if selfPartyMember.LevelUp() {
			success = true
		}
	}
	return success
}

// Move updates a party's coordinates
func (self *Party) Move(x, y int) {
	self.X += x
	self.Y += y
}

// Rest Whole Party
func (self *Party) Rest() {
	for _, partyMember := range self.Members {
		partyMember.Rest()
	}
}

// Battle other party
func (self *Party) Battle2(other *Party) {
	for _, selfPartyMember := range self.Members {
		for _, otherPartyMember := range other.Members {
			selfPartyMember.Duel(otherPartyMember)
		}
	}
}


// Battle other party
func (self *Party) Duel(otherParty *Party) {

	for selfParty.Stats.GetHealth() > 0 && otherParty.Stats.GetHealth() > 0 {

		time.Sleep(200 * time.Millisecond)

		for member := range selfParty.Members {
			target := otherParty.TargetMember()


			if member.Status.Stunned == 0 {
				chosenSkill := member.ChooseSkill()
				switch chosenSkill.Name {
				case "Double Strike":
					member.DoubleStrike(target)
				case "Flash Heal":
					member.FlashHeal()
				case "Ghost Blade":
					member.GhostBlade(target)
				case "Heal":
					member.Heal()
				case "Ice Blast":
					member.IceBlast(target)
				case "Lightning Bolt":
					member.LightningBolt(target)
				case "Rend":
					member.Rend(target)
				case "Smite":
					member.Smite(target)
				case "Stun":
					member.Stun(target)
				default:
					member.BasicAttack(target, member.Stats.Strength+(member.Stats.Agility/2))
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
			} else {
				member.Status.Stunned--
			}
	

		}
		
		for member := range otherParty.Members {

		}

		

	}

	// if self.Stats.Health >= other.Stats.Health {
	// 	color.Cyan("\n%s Wins the duel\n", self.Stats.Name)
	// 	color.Red("\nOther xp is %d\n", other.Stats.XP)
	// 	self.Stats.XP += other.Stats.XP
	// 	other.Stats.XP = 0

	// 	return
	// }

	// color.Cyan("%s Wins the duel\n", other.Stats.Name)
}