package characters

import (
	"the_brink/inventory"
)

func NewThug(name string, level int) Thug {

	// Base Layer
	thug := Thug{
		Character: Character {
			Status: Status{
				Stunned: 0,
			},
		},
	}

	// Set Stats
	stats := Stats{
		Name: name,
		Class: "Vagabond",
		Level: 1,
		Vitality: 4,
		Strength: 5,
		Agility: 2,
		Intelligence: 2,
		LevelBonuses: LevelBonuses {
			Vitality: 2,
			Strength: 1,
			Agility: 0,
			Intelligence: 1,
		},
		SkillSlots: []Skill{
			Skill {
				Name: "BasicAttack",
				Cost: 0,
				CoolDownMax: 1,
				CoolDown: 0,
			},
			Skill{
				Name: "DoubleStrike",
				Cost: 35,
				CoolDownMax: 4,
				CoolDown: 0,
			},
		},
	}
	thug.Character.Stats = stats

	// Set class
	thug.Character = Vagabond(thug.Character)

	// Set Inventory
	thug.Inventory = inventory.Inventory {
		Owner: thug.Character.Stats.Name,
		Gold: 10,
	}

	// Level Up
	for i := 1; i < level; i++ {
		thug.Character.LevelUp()
	}

	thug.Character.Rest()
	return thug
}

