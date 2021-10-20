package world

import (
	"math/rand"
)

var persistentTerrain map[Tile]int

// CreateMap will return a blank grid
func CreateMap(xMax int, yMax int) map[Tile]int {
	// Declare our map grid
	grid := map[Tile]int{}

	// Creat base map layer
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			tile := Tile{
				X: x,
				Y: y,
			}
			// '.'
			grid[tile] = 46 
		}
	}

	// Create randomized persistent terrain
	// Bushes
	for i := 0; i < 10; i++ {
		x := rand.Intn(xMax)
		y := rand.Intn(yMax)
		tile := Tile{
			X: x,
			Y: y,
		}
		// '*'
		grid[tile] = 42
	}

	return grid
}

// UpdateMap will set the position and rune for all existent parties
func (world *World) UpdateMap() {

	// Regerate map from scratch
	// TODO: Optimize me
	for y := 0; y < world.YMax; y++ {
		for x := 0; x < world.XMax; x++ {
			tile := Tile{
				X: x,
				Y: y,
			}

			// don't overwrite bushes
			if world.Tiles[tile] != 46 {
				world.Tiles[tile] = 42
			}
		}
	}


	// Move parties
	for _, party := range world.Parties {
		coordinates := Tile{
			X: party.X,
			Y: party.Y,
		}
		world.Tiles[coordinates] = party.Rune
	}
}
