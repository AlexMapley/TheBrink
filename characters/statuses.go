package characters

func (self *Character) Stun(other *Character, turns int) {
	other.Status.Stunned += turns
}