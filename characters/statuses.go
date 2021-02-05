package characters

func (self *Character) Stun(other *Character, turns int) {
	color.HiGreen("* %s knocks the wind out of %s *\n", self.Stats.Name, other.Stats.Name)
	other.Status.Stunned += turns
}