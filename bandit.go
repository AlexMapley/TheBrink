package main

type Bandit struct {
	Character Character
}

func NewBandit(name string) Bandit {
	bandit := Bandit{}

	bandit.Character.Name = name

	bandit.Character.Vitality = 3
	bandit.Character.Strength = 4
	bandit.Character.Agility = 6
	bandit.Character.Intelligence = 5

	bandit.Character.Health = (bandit.Character.Vitality * 8) + (bandit.Character.Strength * 2)
	bandit.Character.Mana = (bandit.Character.Intelligence * 10)

	return bandit
}
