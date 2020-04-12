package main

type Player struct {
	Character Character
}

func NewPlayer(name string) Player {
	player := Player{}

	player.Character.Name = name

	player.Character.Vitality = 5
	player.Character.Strength = 5
	player.Character.Agility = 5
	player.Character.Intelligence = 5

	player.Character.Health = (player.Character.Vitality * 8) + (player.Character.Strength * 2)
	player.Character.Mana = (player.Character.Intelligence * 10)

	return player
}
