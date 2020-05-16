package inventory

import (
	"github.com/fatih/color"
)

func (inventory *Inventory) Display() {
	color.Yellow("\n-------------\n%s's Inventory\n-------------\nGold: %d\n\n",
		inventory.Owner,
		inventory.Gold,
	)
}
