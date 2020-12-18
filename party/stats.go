package party

// GetHealth returns the current and max health of the part,
// for example (75,100) translating to a party at 75/100 health
func (party *Party) GetHealth() (healthTotal float64, healthMaxTotal float64) {
	for _, member := range party.Members {
		healthTotal += member.Health
		healthMaxTotal += member.MaxHealth()
	}
	return healthTotal, healthMaxTotal
}