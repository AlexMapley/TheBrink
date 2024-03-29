package characters

import (
	"the_brink/inventory"
)

func NewBandit(name string, level int) Bandit {

	// Base Layer
	bandit := Bandit{
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
					Name:        "Double Strike",
					Cost:        35,
					CoolDownMax: 4,
					CoolDown:    0,
				},
			},
		},
	}

	// Set Stats
	stats := Stats{
		Name:         name,
		Class:        "Bandit",
		ClassHash:    1,
		Level:        1,
		XP:           300,
		Vitality:     7,
		Strength:     3,
		Agility:      5,
		Intelligence: 3,
		LevelBonuses: LevelBonuses{
			Vitality:     2,
			Strength:     1,
			Agility:      2,
			Intelligence: 1,
		},
	}
	bandit.Character.Stats = stats

	// Set Inventory
	bandit.Character.Inventory = inventory.Inventory{
		Owner: bandit.Character.Stats.Name,
		Gold:  25,
	}

	// Level Up
	for i := 1; i < level; i++ {
		bandit.Character.Stats.XP += 1000
		if i == 10 {
			_ = bandit.Character.MutateInfected()
		}
		_ = bandit.Character.LevelUp()
	}

	bandit.Character.Rest()
	return bandit
}
