package inventory

func (inventory *Inventory) Display() {
	color.Cyan("\n-------------\n%s's Inventory\n-------------\nLevel: %d\bGold: %s\n\n",
		Inventory.Gold,
	)
}