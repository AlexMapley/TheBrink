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
	stats.MaxHealth = stats.DetermineMaxHealth()
	stats.MaxFocus = stats.DetermineMaxFocus()
	stats.Health = stats.MaxHealth
	stats.Focus = stats.MaxFocus
	bandit.Character.Stats = stats

	// Set Bandit Inventory
	bandit.Inventory = inventory.Inventory {
		Owner: bandit.Character.Stats.Name,
		Gold: 25,
	}

	// Level Up Bandit
	for i := 1; i < level; i++ {
		bandit.LevelUp()
	}

	return bandit
}

func (bandit *Bandit) LevelUpBandit() {

	// increase level
	bandit.Character.Stats.Level++

	// increase core stats
	bandit.Character.Stats.Vitality += bandit.Character.Stats.LevelBonuses.Vitality
	bandit.Character.Stats.Strength += bandit.Character.Stats.LevelBonuses.Strength
	bandit.Character.Stats.Agility += bandit.Character.Stats.LevelBonuses.Agility
	bandit.Character.Stats.Intelligence += bandit.Character.Stats.LevelBonuses.Intelligence

	// rest
	bandit.Character.Rest()
}
