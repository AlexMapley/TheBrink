package characters

import (
	"the_brink/inventory"
)

func NewSidekick(name string) Sidekick {

	// Base layer
	sidekick := Sidekick{
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
					Name:        "Bark",
					Cost:        3,
					CoolDownMax: 4,
					CoolDown:    0,
				},
			},
		},
	}

	// Set stats
	stats := Stats{
		Name:         name,
		ClassHash:    1,
		Level:        1,
		Vitality:     7,
		Strength:     3,
		Agility:      4,
		Intelligence: 3,
		LevelBonuses: LevelBonuses{
			Vitality:     1,
			Strength:     1,
			Agility:      1,
			Intelligence: 1,
		},
	}
	sidekick.Character.Stats = stats

	// Set sidekick inventory
	sidekick.Character.Inventory = inventory.Inventory{
		Owner: sidekick.Character.Stats.Name,
		Gold:  0,
	}

	sidekick.Character.Rest()
	return sidekick
}
