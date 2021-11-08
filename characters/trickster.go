package characters

import (
	"the_brink/inventory"
)

func NewTrickster(name string, level int) Trickster {

	// Base Layer
	trickster := Trickster{
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
					Name:            "SneakAttack",
					Cost:            35,
					CoolDownInitial: 7,
					CoolDownMax:     7,
					CoolDown:        0,
				},
			},
		},
	}

	// Set Stats
	stats := Stats{
		Name:         name,
		Class:        "Trickster",
		ClassHash:    1,
		Level:        1,
		XP:           450,
		Vitality:     6,
		Strength:     2,
		Agility:      4,
		Intelligence: 6,
		LevelBonuses: LevelBonuses{
			Vitality:     2,
			Strength:     1,
			Agility:      1,
			Intelligence: 2,
		},
	}
	trickster.Character.Stats = stats

	// Set Inventory
	trickster.Character.Inventory = inventory.Inventory{
		Owner: trickster.Character.Stats.Name,
		Gold:  10,
	}

	// Level Up
	for i := 1; i < level; i++ {
		trickster.Character.Stats.XP += 1000
		if i == 10 {
			_ = trickster.Character.MutateInfected()
		}
		_ = trickster.Character.LevelUp()
	}

	trickster.Character.Rest()
	return trickster
}
