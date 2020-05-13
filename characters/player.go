package characters

import (
	"the_brink/inventory"
)

type Player struct {
	Stats Stats
	Inventory inventory.Inventory
}

func NewPlayer(name string) Player {

	// Base layer
	player := Player{}

	// Set stats
	stats = Stats{
		Name: name,
		Level: 1,
		Vitality: 5,
		Strength: 5,
		Agility: 5,
		Intelligence: 5,
	}
	stats.Health = statsHealthValue()
	stats.Focus = statsFocusValue()
	stats.CurrentHealth = stats.Health
	stats.CurrentFocus =  = stats.Focus

	player.Character.Stats = stats

	// Set Player Inventory
	player.Inventory = inventory.Inventory {
		Gold: 100,
	}

	return player
}
