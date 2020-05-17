package characters

import (
	"the_brink/inventory"
)

func NewPlayer(name string, class string) Player {

	// Base layer
	player := Player{}

	// Set stats
	stats := Stats{
		Name: name,
		Level: 1,
		Vitality: 5,
		Strength: 5,
		Agility: 5,
		Intelligence: 5,
		LevelBonuses: LevelBonuses {
			Vitality: 1,
			Strength: 1,
			Agility: 1,
			Intelligence: 1,
		},
	}
	stats.MaxHealth = stats.DetermineMaxHealth()
	stats.MaxFocus = stats.DetermineMaxFocus()
	stats.Health = stats.MaxHealth
	stats.Focus =  stats.MaxFocus
	player.Character.Stats = stats

	// Set player class
	switch class {
	case "warrior":
		player.Character = Warrior(player.Character)
	case "rogue":
		player.Character = Rogue(player.Character)
	case "wizard":
		player.Character = Wizard(player.Character)
	}

	player.Character.Rest()

	// Set Player Inventory
	player.Inventory = inventory.Inventory {
		Owner: player.Character.Stats.Name,
		Gold: 100,
	}

	return player
}

func (player *Player) LevelUpPlayer() {

	// increase level
	player.Character.Stats.Level++

	// increase core stats
	player.Character.Stats.Vitality += player.Character.Stats.LevelBonuses.Vitality
	player.Character.Stats.Strength += player.Character.Stats.LevelBonuses.Strength
	player.Character.Stats.Agility += player.Character.Stats.LevelBonuses.Agility
	player.Character.Stats.Intelligence += player.Character.Stats.LevelBonuses.Intelligence

	// rest
	player.Character.Rest()
}
