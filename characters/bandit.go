package characters

import (
	"the_brink/inventory"
)

func NewBandit(name string, level int) Bandit {

	// Base Layer
	bandit := Bandit{
		Character: Character {
			Status: Status{
				Stunned: 0,
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
		},
	}

	// Set Stats
	stats := Stats{
		Name: name,
		Class: "Vagabond",
		Level: 1,
		Vitality: 3,
		Strength: 3,
		Agility: 5,
		Intelligence: 3,
		LevelBonuses: LevelBonuses {
			Vitality: 1,
			Strength: 1,
			Agility: 1,
			Intelligence: 1,
		},
	}
	bandit.Character.Stats = stats

	// Set class
	bandit.Character, _ = Vagabond(bandit.Character)

	// Set Inventory
	bandit.Inventory = inventory.Inventory {
		Owner: bandit.Character.Stats.Name,
		Gold: 25,
	}

	// Level Up
	for i := 1; i < level; i++ {
		if i == 10 {
			bandit.Character, _ = Infected(bandit.Character)
		}
		bandit.Character.LevelUp()
	}

	bandit.Character.Rest()
	return bandit
}

