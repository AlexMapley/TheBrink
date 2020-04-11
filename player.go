package main

import "fmt"

type Player struct {
	Name      string
	Character Character
}

func newPlayer(name string) Player {
	player := Player{}

	player.Name = name

	player.Character.Vitality = 5
	player.Character.Strength = 5
	player.Character.Agility = 5
	player.Character.Intelligence = 5

	player.Character.Health = (player.Character.Vitality * 8) + (player.Character.Strength * 2)
	player.Character.Mana = (player.Character.Intelligence * 10)

	return player
}

func (player *Player) Stats() {
	fmt.Printf("\n-------------\n%s-------------\nHealth: %d\nMana: %d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d",
		player.Name,
		player.Character.Health,
		player.Character.Mana,
		player.Character.Vitality,
		player.Character.Strength,
		player.Character.Agility,
		player.Character.Intelligence,
	)
}
