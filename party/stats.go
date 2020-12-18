package party

// GetHealth returns the current total health of the party
func (party *Party) GetHealth() float64 {
	var healthTotal float64
	for _, member := range party.Members {
		healthTotal += member.Stats.Health
	}
	return healthTotal
}