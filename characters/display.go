package characters

func (stats *Stats) Display() {
	color.Cyan("\n-------------\n%s\n-------------\nLevel: %d\nClass: %s\nHealth: %d/%d\nFocus: %d/%d\nVitality: %d\nStrength: %d\nAgility: %d\nIntelligence: %d\nCritical: %d\nDodge: %d\n\n",
		stats.Name,
		stats.Level,
		stats.Class,
		stats.Health,
		stats.MaxHealth,
		stats.MaxFocus,
		stats.Focus,
		stats.Vitality,
		stats.Strength,
		stats.Agility,
		stats.Intelligence,
		stats.CriticalValue(),
		stats.DodgeValue(),
	)
}