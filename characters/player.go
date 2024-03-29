package characters

import (
	"the_brink/inventory"
)

func NewPlayer(name string) Player {

	// Base layer
	player := Player{
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
			},
		},
	}

	// Set stats
	stats := Stats{
		Name:         name,
		ClassHash:    1,
		Level:        1,
		Vitality:     8,
		Strength:     8,
		Agility:      8,
		Intelligence: 9,
		LevelBonuses: LevelBonuses{
			Vitality:     2,
			Strength:     2,
			Agility:      2,
			Intelligence: 2,
		},
	}
	player.Character.Stats = stats

	// Set Player Inventory
	player.Character.Inventory = inventory.Inventory{
		Owner: player.Character.Stats.Name,
		Gold:  100,
	}

	player.Character.Rest()
	return player
}
