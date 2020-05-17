package characters

import (
	"the_brink/inventory"
)

func NewBandit(name string, level int) Bandit {

	// Base Layer
	bandit := Bandit{}

	// Set Stats
	stats := Stats{
		Name: name,
		Class: "Vagabond",
		Level: 1,
		Vitality: 3,
		Strength: 2,
		Agility: 4,
		Intelligence: 3,
		LevelBonuses: LevelBonuses {
			Vitality: 1,
			Strength: 1,
			Agility: 1,
			Intelligence: 0,
		},
	}
	bandit.Character.Stats = stats

	// Set class
	bandit.Character = Vagabond(bandit.Character)

	// Set Inventory
	bandit.Inventory = inventory.Inventory {
		Owner: bandit.Character.Stats.Name,
		Gold: 25,
	}

	// Level Up
	for i := 1; i < level; i++ {
		bandit.Character.LevelUp()
	}

	bandit.Character.Rest()
	return bandit
}

