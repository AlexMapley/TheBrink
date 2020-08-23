package characters

import (
	"the_brink/inventory"
)

func NewThug(name string, level int) Thug {

	// Base Layer
	thug := Thug{
		Character: Character{
			Status: Status{
				Stunned: 0,
			},
			SkillSlots: []Skill{
				Skill{
					Name:        "Basic Attack",
					Cost:        0,
					CoolDownMax: 1,
					CoolDown:    0,
				},
				Skill{
					Name:        "Stun",
					Cost:        45,
					CoolDownMax: 7,
					CoolDown:    0,
				},
			},
		},
	}

	// Set Stats
	stats := Stats{
		Name:         name,
		Class:        "Thug",
		ClassHash:    1,
		Level:        1,
		XP:           450,
		Vitality:     5,
		Strength:     4,
		Agility:      2,
		Intelligence: 1,
		LevelBonuses: LevelBonuses{
			Vitality:     3,
			Strength:     1,
			Agility:      1,
			Intelligence: 0,
		},
	}
	thug.Character.Stats = stats

	// Set Inventory
	thug.Character.Inventory = inventory.Inventory{
		Owner: thug.Character.Stats.Name,
		Gold:  10,
	}

	// Level Up
	for i := 1; i < level; i++ {
		thug.Character.Stats.XP += 1000
		if i == 10 {
			thug.Character, _ = Infected(thug.Character)
		}
		_ = thug.Character.LevelUp()
	}

	thug.Character.Rest()
	return thug
}
