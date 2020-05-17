package characters

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

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots, 
		Skill{
			Name: "DoubleStrike",
			Cost: 35,
			CoolDownMax: 4,
			CoolDown: 0,
		},
	)

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

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots, 
		Skill{
			Name: "DoubleStrike",
			Cost: 35,
			CoolDownMax: 4,
			CoolDown: 0,
		},
	)

	return res
}

func Wizard(character Character) Character{
	res := character

	res.Stats.Class = "Wizard"

	// Raw Stat Boosts
	res.Stats.Vitality += 2
	res.Stats.Strength += 2
	res.Stats.Agility += 2
	res.Stats.Intelligence += 8
	res.Stats.Expertise += 1
	res.Stats.Block += 1

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Intelligence += 2
	res.Stats.LevelBonuses.Vitality += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots, 
		Skill{
			Name: "LightningBolt",
			Cost: 60,
			CoolDownMax: 6,
			CoolDown: 0,
		},
	)
	res.SkillSlots = append(
		res.SkillSlots, 
		Skill{
			Name: "Heal",
			Cost: 50,
			CoolDownMax: 5,
			CoolDown: 0,
		},
	)

	return res
}

