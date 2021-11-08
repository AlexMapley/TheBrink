package characters

//////////////////
// Tier 1 Classes
// - Can be applied to any base character
// - requires no prior class
////////////////////

// MutateRogue
func (character *Character) MutateRogue() (allowed bool) {

	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return
	}

	
	character.Stats.Class = "Rogue"
	character.Stats.ClassHash *= 2

	// Raw Stat Boosts
	character.Stats.Vitality += 6
	character.Stats.Strength += 6
	character.Stats.Agility += 10
	character.Stats.Intelligence += 4
	character.Stats.Expertise += 4

	// leveling Stat Boosts
	character.Stats.LevelBonuses.Agility += 1
	character.Stats.LevelBonuses.Strength += 1
	character.Stats.LevelBonuses.Critical += 1
	character.Stats.LevelBonuses.Dodge += 1

	// Add Skills
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Double Strike",
			Cost:            float64(14),
			CoolDownInitial: 0,
			CoolDownMax:     3,
			CoolDown:        0,
		},
	)
	return
}

// MutateWarrior
func (character *Character)  MutateWarrior() (allowed bool) {
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return
	}
	character.Stats.Class = "Warrior"
	character.Stats.ClassHash *= 3

	// Raw Stat Boosts
	character.Stats.Vitality += 8
	character.Stats.Strength += 8
	character.Stats.Agility += 4
	character.Stats.Intelligence += 2
	character.Stats.Expertise += 2
	character.Stats.Block += 2

	// leveling Stat Boosts
	character.Stats.LevelBonuses.Vitality += 1
	character.Stats.LevelBonuses.Strength += 1
	character.Stats.LevelBonuses.Expertise += 1
	character.Stats.LevelBonuses.Block += 1

	// Add Skills
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Slash",
			Cost:            float64(30),
			CoolDownInitial: 0,
			CoolDownMax:     7,
			CoolDown:        0,
		},
	)
	return
}

// MutateWizard
func (character *Character) MutateWizard() (allowed bool) {
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return
	}
	character.Stats.Class = "Wizard"
	character.Stats.ClassHash *= 5

	// Raw Stat Boosts
	character.Stats.Vitality += 4
	character.Stats.Strength += 2
	character.Stats.Agility += 3
	character.Stats.Intelligence += 9
	character.Stats.Expertise += 4
	character.Stats.Block += 4
	character.Stats.Dodge += 4
	character.Stats.Critical += 4

	// leveling Stat Boosts
	character.Stats.LevelBonuses.Intelligence += 1
	character.Stats.LevelBonuses.Vitality += 1
	character.Stats.LevelBonuses.Dodge += 1
	character.Stats.LevelBonuses.Expertise += 1

	// Add Skills
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Lightning Bolt",
			Cost:            float64(24),
			CoolDownInitial: 0,
			CoolDownMax:     4,
			CoolDown:        0,
		},
	)
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Icicle",
			Cost:            float64(35),
			CoolDownInitial: 2,
			CoolDownMax:     6,
			CoolDown:        0,
		},
	)
	return
}

// MutateCleric
func (character *Character) MutateCleric() (allowed bool) {
	// Reject If:
	// - unit has a class already
	if character.Stats.Class != "" {
		return
	}
	character.Stats.Class = "Cleric"
	character.Stats.ClassHash *= 7

	// Raw Stat Boosts
	character.Stats.Vitality += 6
	character.Stats.Strength += 3
	character.Stats.Agility += 3
	character.Stats.Intelligence += 7
	character.Stats.Expertise += 4
	character.Stats.Block += 5
	character.Stats.Dodge += 2
	character.Stats.Critical += 2

	// leveling Stat Boosts
	character.Stats.LevelBonuses.Intelligence += 1
	character.Stats.LevelBonuses.Vitality += 1
	character.Stats.LevelBonuses.Block += 1
	character.Stats.LevelBonuses.Expertise += 1

	// Add Skills
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Flash Heal",
			Cost:            float64(24),
			CoolDownInitial: 3,
			CoolDownMax:     4,
			CoolDown:        4,
		},
	)
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Smite",
			Cost:            float64(18),
			CoolDownInitial: 0,
			CoolDownMax:     5,
			CoolDown:        5,
		},
	)
	return
}


//////////////////
// Tier 2 Classes
// - Level Restricted
// - ClassHash Restricted
////////////////////

// MutateInfected
func (character *Character) MutateInfected() (allowed bool) {
	// Reject If:
	// * unit has a class already * //
	if character.Stats.Level < 10 {
		allowed = false; return
	}
	character.Stats.Class = "Infected " + character.Stats.Class

	// Raw Stat Boost
	character.Stats.Vitality += 0
	character.Stats.Strength += 0
	character.Stats.Agility += 0
	character.Stats.Intelligence += 0
	character.Stats.Expertise += 0
	character.Stats.Block += 0

	// leveling Stat Boosts
	character.Stats.LevelBonuses.Strength += 2
	character.Stats.LevelBonuses.Agility += 2
	character.Stats.LevelBonuses.Intelligence += 2
	character.Stats.LevelBonuses.Vitality += 3
	character.Stats.LevelBonuses.Expertise += 1
	character.Stats.LevelBonuses.Block += 1
	character.Stats.LevelBonuses.Critical += 1

	return
}

// Paladin
func (character *Character) MutatePaladin() (allowed bool) {
	// Reject If:
	// * unit has a class already * //
	if character.Stats.Level < 5 {
		return
	}
	// * ClassHash modulos are incompatible * //
	if character.Stats.ClassHash%11 == 0 || character.Stats.ClassHash%13 == 0 || character.Stats.ClassHash%17 == 0 {
		return
	}

	character.Stats.Class = character.Stats.Class + " Paladin"
	character.Stats.ClassHash *= 11

	// Raw stat boosts
	character.Stats.Strength += 3
	character.Stats.Vitality += 2
	character.Stats.Intelligence += 1

	// leveling Stat Boosts
	character.Stats.LevelBonuses.Strength += 2
	character.Stats.LevelBonuses.Intelligence += 1
	character.Stats.LevelBonuses.Vitality += 1
	character.Stats.LevelBonuses.Block += 1

	// Add Skills
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Heal",
			Cost:            float64(45),
			CoolDownInitial: 5,
			CoolDownMax:     5,
			CoolDown:        5,
		},
	)
	return
}

// MutateNightBlade
func (character *Character) MutateNightBlade() (allowed bool) {
	// Reject If:
	// * unit has a class already * //
	if character.Stats.Level < 5 {
		return
	}
	// * ClassHash modulos are incompatible * //
	if character.Stats.ClassHash%11 == 0 || character.Stats.ClassHash%13 == 0 || character.Stats.ClassHash%17 == 0 {
		return
	}

	character.Stats.Class = character.Stats.Class + " Nightblade"
	character.Stats.ClassHash *= 13

	// Raw stat boosts
	character.Stats.Vitality += 2
	character.Stats.Strength += 1
	character.Stats.Agility += 3
	character.Stats.Intelligence += 3
	character.Stats.Expertise += 2
	character.Stats.Block += 0
	character.Stats.Dodge += 2
	character.Stats.Critical += 2


	// Leveling Stat Boosts
	character.Stats.LevelBonuses.Strength += 1
	character.Stats.LevelBonuses.Intelligence += 2
	character.Stats.LevelBonuses.Agility += 2
	character.Stats.LevelBonuses.Critical += 1
	character.Stats.LevelBonuses.Dodge += 1

	// Add Skills
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Sneak Attack",
			Cost:            float64(24),
			CoolDownInitial: 0,
			CoolDownMax:     4,
			CoolDown:        0,
		},
	)
	return
}

// MutateSwordsman
func (character *Character) MutateSwordsman() (allowed bool) {
	// Reject If:
	// * unit has a class already * //
	if character.Stats.Level < 5 {
		return
	}
	// * ClassHash modulos are incompatible * //
	if character.Stats.ClassHash%11 == 0 || character.Stats.ClassHash%13 == 0 || character.Stats.ClassHash%17 == 0 {
		return
	}

	character.Stats.Class = character.Stats.Class + " Swordsman"
	character.Stats.ClassHash *= 17

	// Raw stat boosts
	character.Stats.Agility += 3
	character.Stats.Strength += 2
	character.Stats.Vitality += 1

	// leveling Stat Boosts
	character.Stats.LevelBonuses.Strength += 1
	character.Stats.LevelBonuses.Vitality += 1
	character.Stats.LevelBonuses.Agility += 1
	character.Stats.LevelBonuses.Critical += 1
	character.Stats.LevelBonuses.Expertise += 1
	character.Stats.LevelBonuses.Block += 1
	character.Stats.LevelBonuses.Dodge += 1

	// Add Skills
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Double Strike",
			Cost:            float64(12),
			CoolDownInitial: 0,
			CoolDownMax:     4,
			CoolDown:        0,
		},
	)
		// Add Skills
		character.SkillSlots = append(
			character.SkillSlots,
			Skill{
				Name:            "Slash",
				Cost:            float64(28),
				CoolDownInitial: 1,
				CoolDownMax:     6,
				CoolDown:        0,
			},
		)
	character.SkillSlots = append(
		character.SkillSlots,
		Skill{
			Name:            "Flash Heal",
			Cost:            float64(16),
			CoolDownInitial: 0,
			CoolDownMax:     3,
			CoolDown:        0,
		},
	)
	return
}
