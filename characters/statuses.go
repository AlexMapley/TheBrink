package characters

func (character *Character) Stun(other *Character, turns int) {
	other.Status.Stunned += turns
}