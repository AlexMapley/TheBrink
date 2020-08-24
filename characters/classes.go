package characters

//////////////////
// Tier 1 Classes
// - Can be applied to any base character
// - requires no prior class
////////////////////
func Rogue(character Character) (Character, bool) {
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return character, false
	}

	res := character
	res.Stats.Class = "Rogue"
	res.Stats.ClassHash *= 2

	// Raw Stat Boosts
	res.Stats.Vitality += 2
	res.Stats.Strength += 4
	res.Stats.Agility += 6
	res.Stats.Intelligence += 2
	res.Stats.Expertise += 1

	// leveling Stat Boosts
	res.Stats.LevelBonuses.Agility += 1
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Critical += 1
	res.Stats.LevelBonuses.Dodge += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Double Strike",
			Cost:            float64(18),
			CoolDownInitial: 0,
			CoolDownMax:     4,
			CoolDown:        0,
		},
	)

	return res, true
}

// Warrior
func Warrior(character Character) (Character, bool) {
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return character, false
	}

	res := character
	res.Stats.Class = "Warrior"
	res.Stats.ClassHash *= 3

	// Raw Stat Boosts
	res.Stats.Vitality += 5
	res.Stats.Strength += 5
	res.Stats.Agility += 2
	res.Stats.Intelligence += 1
	res.Stats.Expertise += 1
	res.Stats.Block += 1

	// leveling Stat Boosts
	res.Stats.LevelBonuses.Vitality += 1
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Expertise += 1
	res.Stats.LevelBonuses.Block += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Rend",
			Cost:            float64(30),
			CoolDownInitial: 0,
			CoolDownMax:     7,
			CoolDown:        0,
		},
	)

	return res, true
}

// Wizard
func Wizard(character Character) (Character, bool) {
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return character, false
	}

	res := character
	res.Stats.Class = "Wizard"
	res.Stats.ClassHash *= 5

	// Raw Stat Boosts
	res.Stats.Vitality += 3
	res.Stats.Strength += 2
	res.Stats.Agility += 3
	res.Stats.Intelligence += 7
	res.Stats.Expertise += 3
	res.Stats.Block += 3
	res.Stats.Dodge += 3
	res.Stats.Critical += 3

	// leveling Stat Boosts
	res.Stats.LevelBonuses.Intelligence += 1
	res.Stats.LevelBonuses.Vitality += 1
	res.Stats.LevelBonuses.Dodge += 1
	res.Stats.LevelBonuses.Expertise += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Lightning Bolt",
			Cost:            float64(24),
			CoolDownInitial: 0,
			CoolDownMax:     4,
			CoolDown:        0,
		},
	)
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Ice Blast",
			Cost:            float64(35),
			CoolDownInitial: 2,
			CoolDownMax:     6,
			CoolDown:        0,
		},
	)

	return res, true
}

// Cleric
func Cleric(character Character) (Character, bool) {
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return character, false
	}

	res := character
	res.Stats.Class = "Cleric"
	res.Stats.ClassHash *= 7

	// Raw Stat Boosts
	res.Stats.Vitality += 5
	res.Stats.Strength += 2
	res.Stats.Agility += 3
	res.Stats.Intelligence += 6
	res.Stats.Expertise += 3
	res.Stats.Block += 3
	res.Stats.Dodge += 2
	res.Stats.Critical += 2

	// leveling Stat Boosts
	res.Stats.LevelBonuses.Intelligence += 1
	res.Stats.LevelBonuses.Vitality += 1
	res.Stats.LevelBonuses.Block += 1
	res.Stats.LevelBonuses.Expertise += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Flash Heal",
			Cost:            float64(32),
			CoolDownInitial: 5,
			CoolDownMax:     5,
			CoolDown:        5,
		},
	)
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Smite",
			Cost:            float64(24),
			CoolDownInitial: 5,
			CoolDownMax:     5,
			CoolDown:        5,
		},
	)

	return res, true
}


//////////////////
// Tier 2 Classes
// - Level Restricted
// - ClassHash Restricted
////////////////////

// Infected
func Infected(character Character) (Character, bool) {
	// Reject If:
	// * unit has a class already * //
	if character.Stats.Level < 10 {
		return character, false
	}

	res := character

	res.Stats.Class = "Infected " + res.Stats.Class

	// Raw Stat Boost
	res.Stats.Vitality += 0
	res.Stats.Strength += 0
	res.Stats.Agility += 0
	res.Stats.Intelligence += 0
	res.Stats.Expertise += 0
	res.Stats.Block += 0

	// leveling Stat Boosts
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Agility += 1
	res.Stats.LevelBonuses.Intelligence += 1
	res.Stats.LevelBonuses.Vitality += 1
	res.Stats.LevelBonuses.Expertise += 1
	res.Stats.LevelBonuses.Block += 0
	res.Stats.LevelBonuses.Critical += 1

	return res, true
}

// Paladin
func Paladin(character Character) (Character, bool) {
	// Reject If:
	// * unit has a class already * //
	if character.Stats.Level < 5 {
		return character, false
	}
	// * ClassHash modulos are incompatible * //
	if character.Stats.ClassHash%11 == 0 || character.Stats.ClassHash%13 == 0 || character.Stats.ClassHash%17 == 0 {
		return character, false
	}

	res := character
	res.Stats.Class = res.Stats.Class + " Paladin"
	res.Stats.ClassHash *= 11

	// Raw stat boosts
	res.Stats.Strength += 3
	res.Stats.Vitality += 2
	res.Stats.Intelligence += 1

	// leveling Stat Boosts
	res.Stats.LevelBonuses.Strength += 2
	res.Stats.LevelBonuses.Intelligence += 1
	res.Stats.LevelBonuses.Vitality += 1
	res.Stats.LevelBonuses.Block += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Heal",
			Cost:            float64(45),
			CoolDownInitial: 5,
			CoolDownMax:     5,
			CoolDown:        5,
		},
	)

	return res, true
}

// NightBlade
func NightBlade(character Character) (Character, bool) {
	// Reject If:
	// * unit has a class already * //
	if character.Stats.Level < 5 {
		return character, false
	}
	// * ClassHash modulos are incompatible * //
	if character.Stats.ClassHash%11 == 0 || character.Stats.ClassHash%13 == 0 || character.Stats.ClassHash%17 == 0 {
		return character, false
	}

	res := character
	res.Stats.Class = res.Stats.Class + " NightBlade"
	res.Stats.ClassHash *= 13

	// Raw stat boosts
	res.Stats.Intelligence += 3
	res.Stats.Agility += 2
	res.Stats.Vitality += 1

	// Leveling Stat Boosts
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Intelligence += 2
	res.Stats.LevelBonuses.Agility += 2
	res.Stats.LevelBonuses.Critical += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "GhostBlade",
			Cost:            float64(25),
			CoolDownInitial: 0,
			CoolDownMax:     4,
			CoolDown:        0,
		},
	)

	return res, true
}

// Duelist
func Duelist(character Character) (Character, bool) {
	// Reject If:
	// * unit has a class already * //
	if character.Stats.Level < 5 {
		return character, false
	}
	// * ClassHash modulos are incompatible * //
	if character.Stats.ClassHash%11 == 0 || character.Stats.ClassHash%13 == 0 || character.Stats.ClassHash%17 == 0 {
		return character, false
	}

	res := character
	res.Stats.Class = res.Stats.Class + " Duelist"
	res.Stats.ClassHash *= 17

	// Raw stat boosts
	res.Stats.Agility += 3
	res.Stats.Strength += 2
	res.Stats.Vitality += 1

	// leveling Stat Boosts
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Vitality += 1
	res.Stats.LevelBonuses.Agility += 1
	res.Stats.LevelBonuses.Critical += 1
	res.Stats.LevelBonuses.Expertise += 1
	res.Stats.LevelBonuses.Block += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Double Strike",
			Cost:            float64(11),
			CoolDownInitial: 0,
			CoolDownMax:     3,
			CoolDown:        0,
		},
	)
	res.SkillSlots = append(
		res.SkillSlots,
		Skill{
			Name:            "Flash Heal",
			Cost:            float64(16),
			CoolDownInitial: 0,
			CoolDownMax:     3,
			CoolDown:        0,
		},
	)

	return res, true
}
