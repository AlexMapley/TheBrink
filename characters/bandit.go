package characters

type Bandit struct {
	Stats Stats
}

func NewBandit(name string) Bandit {
	bandit := Bandit{}

	bandit.Stats.Name = name
	bandit.Stats.Level = 1

	bandit.Stats.Vitality = 3
	bandit.Stats.Strength = 3
	bandit.Stats.Agility = 6
	bandit.Stats.Intelligence = 4

	bandit.Stats.Health = bandit.Stats.HealthValue()
	bandit.Stats.Focus = bandit.Stats.FocusValue()

	bandit.Stats.CurrentHealth = bandit.Stats.Health
	bandit.Stats.CurrentFocus = bandit.Stats.Focus

	return bandit
}
