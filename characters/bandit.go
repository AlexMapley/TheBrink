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
		Class: "Vagabond",
		Level: 1,
		XPCap: 50,
		Vitality: 3,
		Strength: 3,
		Agility: 6,
		Intelligence: 4,

	}
	stats.HealthMax = stats.DetermineMaxHealth()
	stats.FocusMax = stats.DetermineMaxFocus()
	stats.Health = stats.HealthMax
	stats.Focus = stats.FocusMax
	bandit.Character.Stats = stats

	return bandit
}

func LevelUpBandit(bandit Bandit) Bandit{
	res := NewBandit(bandit.Character.Stats.Name)
	res.Character = bandit.Character


	// increase core stats
	res.Character.Stats.Level +=1
	res.Character.Stats.Vitality +=1
	res.Character.Stats.Strength +=1
	res.Character.Stats.Agility +=2
	res.Character.Stats.Intelligence +=1

	// rest
	res.Character.Rest()

	return res
}
