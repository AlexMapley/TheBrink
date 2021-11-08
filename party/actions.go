package party

// Level Up Party
// Mutates an existing party
func (party *Party) LevelUp() bool{
	success := false
	for _, partyMember := range party.Members {
		if partyMember.LevelUp() {
			success = true
		}
	}
	party.Rest()
	return success
}

// Move updates a party's coordinates
func (party *Party) Move(x, y int) {
	party.X += x
	party.Y += y
}

// Rest Whole Party
func (party *Party) Rest() {
	for _, partyMember := range party.Members {
		partyMember.Rest()
	}
}
