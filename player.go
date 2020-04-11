package main

import "fmt"

type Player struct {
	Character Character
}

func newPlayer(name string) Player {
	player := Player{}

	player.Character.Name = name
	player.Character.Health = 100
	player.Character.Mana = 100

	return player
}

func (player *Player) Stats() {
	fmt.Printf("Health: %d\nMana: %d", player.Character.Health, player.Character.Mana)
}
