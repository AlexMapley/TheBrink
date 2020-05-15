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
	stats.HealthMax = stats.DetermineMaxHealth()
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


func LevelUpPlayer(player Player) Player{
	res := NewPlayer(player.Character.Stats.Name)
	res.Character = player.Character

	res.Character.Stats.Level +=1
	res.Character.Stats.Vitality +=1
	res.Character.Stats.Strength +=1
	res.Character.Stats.Agility +=1
	res.Character.Stats.Vitality +=1

	res.Character.Stats.HealthMax = res.Character.Stats.DetermineMaxHealth()
	res.Character.Stats.FocusMax = res.Character.Stats.DetermineMaxFocus()

	return res
}