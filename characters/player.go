package characters

import (
	"the_brink/inventory"
)

type Player struct {
	Stats Stats
	Inventory inventory.Inventory
}

func NewPlayer(name string) Player {
	player := Player{}

	// Set Player Stats
	player.Stats.Name = name
	player.Stats.Level = 1

	player.Stats.Vitality = 5
	player.Stats.Strength = 5
	player.Stats.Agility = 5
	player.Stats.Intelligence = 5

	player.Stats.Health = player.Stats.HealthValue()
	player.Stats.Focus = player.Stats.FocusValue()

	player.Stats.CurrentHealth = player.Stats.Health
	player.Stats.CurrentFocus = player.Stats.Focus

	// Set Player Inventory
	player.Inventory = inventory.Inventory {
		Gold: 100,
	}

	return player
}
