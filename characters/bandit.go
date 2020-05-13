package characters

type Bandit struct {
	Character Character
}

func NewBandit(name string) Bandit {

	// Base Layer
	bandit := Bandit{}

	// Set Stats
	stats := Stats{
		Name: name,
		Level: 1,
		Vitality: 3,
		Strength: 3,
		Agility: 6,
		Intelligence: 4,
	}
	stats.Health = stats.HealthValue()
	stats.Focus = stats.FocusValue()
	stats.CurrentHealth = stats.Health
	stats.CurrentFocus = stats.Focus
	bandit.Character.Stats = stats

	return bandit
}
