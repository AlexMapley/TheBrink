package party

// Rest Whole Party
func (self *Party) Rest() {

	for _, partyMember := range self.Members {
		partyMember.Rest()
	}
}