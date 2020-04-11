package main

import "fmt"

type Player struct {
	Name      string
	Character Character
}

func newPlayer(name string) Player {
	player := Player{}

	player.Name = name
	player.Character.Health = 100
	player.Character.Mana = 100

	return player
}

func (player *Player) Stats() {
	fmt.Printf("\n-------------\n%s-------------\nHealth: %d\nMana: %d", player.Name, player.Character.Health, player.Character.Mana)
}
