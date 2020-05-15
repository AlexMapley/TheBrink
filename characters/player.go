package characters

import (
	"the_brink/inventory"
)

type Player struct {
	Character Character
	Inventory inventory.Inventory
}

func NewPlayer(name string) Player {

	// Base layer
	player := Player{}

	// Set stats
	stats := Stats{
		Name: name,
		Level: 1,
		XPCap: 100,
		Vitality: 5,
		Strength: 5,
		Agility: 5,
		Intelligence: 5,
	}
	stats.HealthMax = stats.DetermineHealthMax()
	stats.FocusMax = stats.DetermineMaxFocus()
	stats.Health = stats.HealthMax
	stats.Focus =  stats.FocusMax
	player.Character.Stats = stats

	// Set Player Inventory
	player.Inventory = inventory.Inventory {
		Gold: 100,
	}

	return player
}

func LevelMultiplier(stats *Stats) {
	stats.Level +=1
	stats.Vilaity +=1
	stats.Strength +=1
	stats.Agility +=1
	stats.Vilaity +=1
}