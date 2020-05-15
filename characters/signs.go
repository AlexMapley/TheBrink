package characters

func Warrior(character Character) Character{
	res := character

	res.Stats.Class = "Warrior"

	res.Stats.Vitality +=4
	res.Stats.Strength +=5
	res.Stats.Agility +=3
	res.Stats.Intelligence +=3

	res.Stats.MaxHealth = res.Stats.DetermineMaxHealth()
	res.Stats.MaxFocus = res.Stats.DetermineMaxFocus()

	return res
}

func Rogue(character Character) Character{
	res := character

	res.Stats.Class = "Rogue"

	res.Stats.Vitality +=3
	res.Stats.Strength +=3
	res.Stats.Agility +=6
	res.Stats.Intelligence +=3

	res.Stats.MaxHealth = res.Stats.DetermineMaxHealth()
	res.Stats.MaxFocus = res.Stats.DetermineMaxFocus()

	return res
}

func Wizard(character Character) Character{
	res := character

	res.Stats.Class = "Wizard"

	res.Stats.Vitality +=2
	res.Stats.Strength +=2
	res.Stats.Agility +=3
	res.Stats.Intelligence +=8

	res.Stats.MaxHealth = res.Stats.DetermineMaxHealth()
	res.Stats.MaxFocus = res.Stats.DetermineMaxFocus()

	return res
}