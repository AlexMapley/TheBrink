package main

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
