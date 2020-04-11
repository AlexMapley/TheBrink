package main

import "fmt"

type Player struct {
	Name      string
	Character Character
}

func newPlayer(name string) Player {
	player := Player{}

	player.Name = name

	player.Character.Vilaity = 5
	player.Character.Strength = 5
	player.Character.Agility = 5
	player.Character.Intelligence = 5

	player.Character.Health = (player.Character.Vilaity * 10)
	player.Character.Mana = (player.Character.Intelligence * 10)

	return player
}

func (player *Player) Stats() {
	fmt.Printf("\n-------------\n%s-------------\nHealth: %d\nMana: %d", player.Name, player.Character.Health, player.Character.Mana)
}
