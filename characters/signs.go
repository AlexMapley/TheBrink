package characters

func Warrior(character Character) Character{
	res := NewPlayer(player.Character.Stats.Name)
	res = character

	res.Stats.Class = "Warrior"

	res.Stats.Vitality +=4
	res.Stats.Strength +=5
	res.Stats.Agility +=3
	res.Stats.Intelligence +=3

	res.Stats.HealthMax = res.Stats.DetermineMaxHealth()
	res.Stats.FocusMax = res.Stats.DetermineMaxFocus()

	return res
}

func Rogue(character Character) Character{
	res := NewPlayer(player.Character.Stats.Name)
	res.Character = player.Character

	res.Stats.Class = "Rogue"

	res.Stats.Vitality +=3
	res.Stats.Strength +=3
	res.Stats.Agility +=6
	res.Stats.Intelligence +=3

	res.Stats.HealthMax = res.Stats.DetermineMaxHealth()
	res.Stats.FocusMax = res.Stats.DetermineMaxFocus()

	return res
}

func Wizard(character Character) Character{
	res := NewPlayer(player.Character.Stats.Name)
	res.Character = player.Character

	res.Stats.Class = "Wizard"

	res.Stats.Vitality +=2
	res.Stats.Strength +=2
	res.Stats.Agility +=3
	res.Stats.Intelligence +=8

	res.Stats.HealthMax = res.Stats.DetermineMaxHealth()
	res.Stats.FocusMax = res.Stats.DetermineMaxFocus()

	return res
}