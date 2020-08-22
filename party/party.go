package party

// Move updates a party's coordinates
func (party *Party) Move(x, y int) {
	party.X += x
	party.Y += y
}
