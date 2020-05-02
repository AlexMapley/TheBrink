package characters

type Player struct {
	Character Character
}

func NewPlayer(name string) Player {
	player := Player{}

	player.Character.Name = name
	player.Character.Level = 1

	player.Character.Vitality = 5
	player.Character.Strength = 5
	player.Character.Agility = 5
	player.Character.Intelligence = 5

	player.Character.Health = player.Character.HealthValue()
	player.Character.Focus = player.Character.FocusValue()

	player.Character.CurrentHealth = player.Character.Health
	player.Character.CurrentFocus = player.Character.Focus

	return player
}
