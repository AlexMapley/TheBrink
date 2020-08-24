package party

// Battle other party
func (self *Party) Battle(other *Party) {
	for _, selfPartyMember := range self.Members {
		for _, otherPartyMember := range other.Members {
			selfPartyMember.Duel(otherPartyMember)
		}
	}
}

// Level Up Part
func (self *Party) Battle(other *Party) {
	for _, selfPartyMember := range self.Members {
		for _, otherPartyMember := range other.Members {
			selfPartyMember.Duel(otherPartyMember)
		}
	}
}

// Move updates a party's coordinates
func (party *Party) Move(x, y int) {
	party.X += x
	party.Y += y
}

// Rest Whole Party
func (self *Party) Rest() {

	for _, partyMember := range self.Members {
		partyMember.Rest()
	}
}