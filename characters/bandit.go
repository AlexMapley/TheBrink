package characters

func NewBandit(name string) Bandit {

	// Base Layer
	bandit := Bandit{}

	// Set Stats
	stats := Stats{
		Name: name,
		Class: "Vagabond",
		Level: 1,
		XPCap: 50,
		Vitality: 3,
		Strength: 3,
		Agility: 6,
		Intelligence: 4,
		LevelBonuses: LevelBonuses {
			Vitality: 1,
			Strength: 1,
			Agility: 3,
			Intelligence: 1,
		},
	}
	stats.MaxHealth = stats.DetermineMaxHealth()
	stats.MaxFocus = stats.DetermineMaxFocus()
	stats.Health = stats.MaxHealth
	stats.Focus = stats.MaxFocus
	bandit.Character.Stats = stats

	return bandit
}

func LevelUpBandit(bandit Bandit) Bandit{
	res := NewBandit(bandit.Character.Stats.Name)
	res.Character = bandit.Character


	// increase level
	res.Character.Stats.Level++

	// increase core stats
	res.Character.Stats.Vitality += bandit.Character.Stats.LevelBonuses.Vitality
	res.Character.Stats.Strength += bandit.Character.Stats.LevelBonuses.Strength
	res.Character.Stats.Agility += bandit.Character.Stats.LevelBonuses.Agility
	res.Character.Stats.Intelligence += bandit.Character.Stats.LevelBonuses.Intelligence

	// rest
	res.Character.Rest()

	return res
}
