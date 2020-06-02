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
	res.Stats.ClassHash *= 3

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
			CoolDownInitial: 0,
			CoolDownMax: 4,
			CoolDown: 0,
		},
	)

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
	res.Stats.ClassHash *= 2

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
			CoolDownInitial: 0,
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
	res.Stats.ClassHash *= 7

	// Raw Stat Boosts
	res.Stats.Vitality += 1
	res.Stats.Strength += 1
	res.Stats.Agility += 1
	res.Stats.Intelligence += 4
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
			CoolDownInitial: 0,
			CoolDownMax: 6,
			CoolDown: 0,
		},
	)
	res.SkillSlots = append(
		res.SkillSlots, 
		Skill{
			Name: "IceBlast",
			Cost: 40,
			CoolDownInitial: 0,
			CoolDownMax: 9,
			CoolDown: 0,
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
func Infected(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
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

	// Levelling Stat Boosts
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
func Paladin(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
	if character.Stats.Level < 5 {
		return character, false
	}
	// - ClassHash mod condition
	if character.Stats.ClassHash % 11 == 0 || character.Stats.ClassHash % 13 == 0 || character.Stats.ClassHash % 17 == 0 {
		return character, false
	}
	
	res := character
	res.Stats.Class = res.Stats.Class + " Paladin"
	res.Stats.ClassHash *= 11

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Strength += 2
	res.Stats.LevelBonuses.Intelligence += 1
	res.Stats.LevelBonuses.Vitality += 1
	res.Stats.LevelBonuses.Block += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots, 
		Skill{
			Name: "Heal",
			Cost: 50,
			CoolDownInitial: 5,
			CoolDownMax: 5,
			CoolDown: 5,
		},
	)

	return res, true
}

// NightBlade
func NightBlade(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
	if character.Stats.Level < 5 {
		return character, false
	}
	// - ClassHash mod condition
	if character.Stats.ClassHash % 11 == 0 || character.Stats.ClassHash % 13 == 0 || character.Stats.ClassHash % 17 == 0 {
		return character, false
	}
	
	res := character
	res.Stats.Class = res.Stats.Class + " NightBlade"
	res.Stats.ClassHash *= 13

	// Levelling Stat Boosts
	res.Stats.LevelBonuses.Strength += 1
	res.Stats.LevelBonuses.Intelligence += 1
	res.Stats.LevelBonuses.Agility += 2
	res.Stats.LevelBonuses.Critical += 1

	// Add Skills
	res.SkillSlots = append(
		res.SkillSlots, 
		Skill{
			Name: "GhostBlade",
			Cost: 45,
			CoolDownInitial: 0,
			CoolDownMax: 5,
			CoolDown: 0,
		},
	)

	return res, true
}

// Duelist
func Duelist(character Character) (Character, bool){
	// Reject If:
	// - unit has a class already
	if character.Stats.Level < 5 {
		return character, false
	}
	// - ClassHash mod condition
	if character.Stats.ClassHash % 11 == 0 || character.Stats.ClassHash % 13 == 0 || character.Stats.ClassHash % 17 == 0 {
		return character, false
	}
	
	res := character
	res.Stats.Class = res.Stats.Class + " Duelist"
	res.Stats.ClassHash *= 17

	// Levelling Stat Boosts
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
			Name: "Rend",
			Cost: 70,
			CoolDownInitial: 0,
			CoolDownMax: 7,
			CoolDown: 0,
		},
	)

	return res, true
}
