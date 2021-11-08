package party

// All methods in actions.go will mutate an existing party struct
// Each method additionally return a `bool success`,
// indicating that the method attempted was successful or not allowed

// LevelUp will attempt to level up each of a party's members
func (party *Party) LevelUp() (bool success){
	for _, partyMember := range party.Members {
		if partyMember.LevelUp() {
			// If at least one party member levels up, 
			// we consider the party's level up successful
			success = true
		}
	}
	party.Rest()
}

// Move updates a party's coordinates
func (party *Party) Move(x, y int) (bool success){
	party.X += x
	party.Y += y

	// Return true for by default
	success = true
}

// Rest Whole Party
func (party *Party) Rest() (bool success) {
	for _, partyMember := range party.Members {
		partyMember.Rest()
	}

	// Return true for by default
	success = true
}
