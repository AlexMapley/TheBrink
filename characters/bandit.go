package characters

type Bandit struct {
	Character Character
}

func NewBandit(name string) Bandit {
	bandit := Bandit{}

	bandit.Character.Name = name

	bandit.Character.Vitality = 3
	bandit.Character.Strength = 3
	bandit.Character.Agility = 6
	bandit.Character.Intelligence = 4

	bandit.Character.Health = bandit.Character.HealthValue()
	bandit.Character.Mana = bandit.Character.ManaValue()

	bandit.Character.CurrentHealth = bandit.Character.Health

	return bandit
}
