package characters

func Warrior(character Character) Character{
	res := character

	res.Stats.Class = "Warrior"

	// Raw Stat Boosts
	res.Stats.Vitality += 4
	res.Stats.Strength += 4
	res.Stats.Agility += 2
	res.Stats.Intelligence += 1
	res.Stats.Expertise += 1
	res.Stats.Block += 1

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Vitality += 1
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Expertise += 1
	res.Stats.LevelBonuses.Block += 1

	return res
}

func Rogue(character Character) Character{
	res := character

	res.Stats.Class = "Rogue"

	// Raw Stat Boosts
	res.Stats.Vitality +=2
	res.Stats.Strength +=2
	res.Stats.Agility +=5
	res.Stats.Intelligence +=2
	res.Stats.Expertise += 1

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Agility += 1
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Expertise += 1

	return res
}

func Vagabond(character Character) Character{
	res := character

	res.Stats.Class = "Vagabond"

	// Raw Stat Boosts
	res.Stats.Vitality += 0
	res.Stats.Strength += 1
	res.Stats.Agility += 3
	res.Stats.Intelligence += 1
	res.Stats.Expertise += 1

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Agility += 1
	res.Stats.LevelBonuses.Vitality += 1

	return res
}

func Wizard(character Character) Character{
	res := character

	res.Stats.Class = "Wizard"

	// Raw Stat Boosts
	res.Stats.Vitality += 2
	res.Stats.Strength += 2
	res.Stats.Agility += 3
	res.Stats.Intelligence += 8
	res.Stats.Expertise += 2
	res.Stats.Block += 2

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Intelligence += 2
	res.Stats.LevelBonuses.Vitality += 1

	return res
}

