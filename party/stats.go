package party

func (party *Party) TotalHealth() float64 {
	var healthTotal float64
	var healthMaxTotal float64

	for _, member := range party.Members {
		healthTotal += member.Health
		healthMaxTotal += member.MaxHealth()
	}

}