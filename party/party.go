package party

func (party *Party) Move(x, y int) {
	party.X += x
	party.Y += y
}