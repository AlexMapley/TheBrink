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

		
		if self.Status.Stunned == 0 {
			chosenSkill := self.ChooseSkill()
			switch chosenSkill.Name {
			case "Double Strike":
				self.DoubleStrike(other)
			case "Flash Heal":
				self.FlashHeal()
			case "Ghost Blade":
				self.GhostBlade(other)
			case "Heal":
				self.Heal()
			case "Ice Blast":
				self.IceBlast(other)
			case "Lightning Bolt":
				self.LightningBolt(other)
			case "Rend":
				self.Rend(other)
			case "Smite":
				self.Smite(other)
			case "Stun":
				self.Stun(other)
			default:
				self.BasicAttack(other, self.Stats.Strength+(self.Stats.Agility/2))
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
				case "Double Strike":
					other.DoubleStrike(self)
				case "Flash Heal":
					other.FlashHeal()
				case "Ghost Blade":
					other.GhostBlade(self)
				case "Heal":
					other.Heal()
				case "Ice Blast":
					other.IceBlast(self)
				case "Lightning Bolt":
					other.LightningBolt(self)
				case "Rend":
					other.Rend(self)
				case "Smite":
					other.Smite(self)
				case "Stun":
					other.Stun(self)
				default:
					other.BasicAttack(self, self.Stats.Strength+(self.Stats.Agility/2))
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

	if self.Stats.Health >= other.Stats.Health {
		color.Cyan("\n%s Wins the duel\n", self.Stats.Name)
		color.Red("\nOther xp is %d\n", other.Stats.XP)
		self.Stats.XP += other.Stats.XP
		other.Stats.XP = 0

		return
	}

	color.Cyan("%s Wins the duel\n", other.Stats.Name)
}