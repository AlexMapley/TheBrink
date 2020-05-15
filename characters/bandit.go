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

func LevelMultiplier(bandit Bandit) Bandit{
	res = NewBandit(bandit.stats.Name)
	res.Character = bandit.Character

	res.Character.Stats.Level +=1
	res.Character.Stats.Vilaity +=1
	res.Character.Stats.Strength +=1
	res.Character.Stats.Agility +=2
	res.Character.Stats.Vilaity +=1

	return res
}
