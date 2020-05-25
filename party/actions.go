package party

// Rest
func (self *Party) Rest() {

	for _, partyMember := range self.Members {
		partyMember.Rest()
	}
}