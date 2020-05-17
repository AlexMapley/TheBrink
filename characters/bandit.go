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
		Strength: 3,
		Agility: 6,
		Intelligence: 4,
		LevelBonuses: LevelBonuses {
			Vitality: 1,
			Strength: 1,
			Agility: 3,
			Intelligence: 1,
		},
	}
	bandit.Character.Stats = stats

	// Set Bandit Inventory
	bandit.Inventory = inventory.Inventory {
		Owner: bandit.Character.Stats.Name,
		Gold: 25,
	}

	// Level Up Bandit
	for i := 1; i < level; i++ {
		bandit.Character.LevelUp()
	}

	// Determine starting health and focus
	stats.MaxHealth = stats.GetMaxHealth()
	stats.MaxFocus = statsGet.MaxFocus()
	stats.Health = stats.MaxHealth
	stats.Focus =  stats.MaxFocus

	return bandit
}

