package world

// NewGrid will return a blank grid
func NewGrid(xMax int, yMax int) map[Tile]rune {
	grid := map[Tile]rune{}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			tile := Tile{
				X: x,
				Y: y,
			}
			grid[tile] = '.'
		}
	}
	return grid
}

// UpdateParties will set the position and rune for all existent parties
func (world *World) UpdateParties() {
	for _, party := range world.Parties {
		coordinates := Tile{
			X: party.X,
			Y: party.Y,
		}
		world.Tiles[coordinates] = party.Rune
	}
}