package party

// Rest Whole Party
func (self *Party) Rest() {

	for _, partyMember := range self.Members {
		partyMember.Rest()
	}
}

// Battle other party
func (self *Party) Battle(other *Party) {
	for _, selfPartyMember := range self.Members {
		for _, otherPartyMember := range other.Members {
			selfPartyMember.Duel(otherPartyMember)
		}
	}
}
