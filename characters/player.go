package characters

import (
	"the_brink/inventory"
)

func NewPlayer(name string, class string) Player {

	// Base layer
	player := Player{
		Character: Character {
			Status: Status{
				Stunned: 0,
			},
		},
	}

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
		SkillSlots: []Skill{
			Skill{
				Name: "BasicAttack",
				Cost: 0,
				CoolDownMax: 1,
				CoolDown: 0,
			},
		},
	}
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

	// Set Player Inventory
	player.Inventory = inventory.Inventory {
		Owner: player.Character.Stats.Name,
		Gold: 100,
	}

	player.Character.Rest()
	return player
}
