package characters

type Bandit struct {
	Character Character
}

func NewBandit(name string) Bandit {
	bandit := Bandit{}

	bandit.Character.Name = name
	bandit.Character.Level = 1

	bandit.Character.Vitality = 3
	bandit.Character.Strength = 3
	bandit.Character.Agility = 6
	bandit.Character.Intelligence = 4

	bandit.Character.Health = bandit.Character.HealthValue()
	bandit.Character.Focus = bandit.Character.FocusValue()

	bandit.Character.CurrentHealth = bandit.Character.Health
	bandit.Character.CurrentFocus = bandit.Character.Focus

	return bandit
}
