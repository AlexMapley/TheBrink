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
	stats.HealthMax = stats.DetermineHealthMax()
	stats.FocusMax = stats.DetermineMaxFocus()
	stats.Health = stats.HealthMax
	stats.Focus = stats.FocusMax
	bandit.Character.Stats = stats

	return bandit
}

func LevelMultiplier(stats *Stats) {
	stats.Level +=1
	stats.Vilaity +=1
	stats.Strength +=1
	stats.Agility +=2
	stats.Vilaity +=1
}
