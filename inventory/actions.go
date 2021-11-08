package inventory

// Loot
func (character *Inventory) Loot(other *Inventory) {

	// Loot gold
	character.Gold += other.Gold
	other.Gold = 0
}
