package characters

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

	player.Character.Health = player.Character.HealthValue()
	player.Character.Mana = player.Character.ManaValue()

	player.Character.CurrentHealth = player.Character.Health
	player.Character.CurrentMana = player.Character.Mana

	return player
}
