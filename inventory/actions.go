package inventory


// Loot
func (self *Inventory) Loot(other *Inventory) {

	// Loot gold
	self.Gold += other.Gold
	other.Gold = 0
}
