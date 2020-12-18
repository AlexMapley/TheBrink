package party

// GetHealth returns the current total health of the party
func (party *Party) GetHealth() (healthTotal float64) {
	for _, member := range party.Members {
		healthTotal += member.Health
	}
	return healthTotal
}