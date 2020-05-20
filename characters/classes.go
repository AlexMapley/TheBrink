package characters

//////////////////
// Tier 1 Classes
// - Can be applied to any base character
// - requires no prior class
////////////////////
func Rogue(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return character, false
	}

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

	return res, true
}

func Vagabond(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return character, false
	}

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

	return res, true
}

func Warrior(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return character, false
	}

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

	return res, true
}

func Wizard(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return character, false
	}

	res := character
	res.Stats.Class = "Wizard"

	// Raw Stat Boosts
	res.Stats.Vitality += 1
	res.Stats.Strength += 1
	res.Stats.Agility += 1
	res.Stats.Intelligence += 8
	res.Stats.Expertise += 1
	res.Stats.Block += 1

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Intelligence += 1
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
			Name: "Ice Blast",
			Cost: 40,
			CoolDownMax: 9,
			CoolDown: 0,
		},
	)

	return res, true
}

//////////////////
// Tier 2 Classes
// - Level Restricted
////////////////////
func Paladin(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
	if character.Stats.Level < 10 {
		return character, false
	}

	res := character

	res.Stats.Class = res.Stats.Class + " Paladin"

	// Raw Stat Boost
	res.Stats.Vitality += 4
	res.Stats.Strength += 4
	res.Stats.Agility += 0
	res.Stats.Intelligence += 3
	res.Stats.Expertise += 2
	res.Stats.Block += 5

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Intelligence += 1
	res.Stats.LevelBonuses.Vitality += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots, 
		Skill{
			Name: "Heal",
			Cost: 50,
			CoolDownMax: 5,
			CoolDown: 0,
		},
	)

	return res, true
}