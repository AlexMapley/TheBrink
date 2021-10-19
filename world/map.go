package world

import (
	"fmt"
	"math/rand"
)

var persistentTerrain map[Tile]rune

// CreateMap will return a blank grid
func CreateMap(xMax int, yMax int) map[Tile]rune {
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

	// Create randomized persistent terrain
	// Bushes
	for i := 0; i < 10; i++ {
		x := rand.Intn(xMax)
		y := rand.Intn(yMax)
		tile := Tile{
			X: x,
			Y: y,
		}
		fmt.Println(x)
		fmt.Println(y)
		grid[tile] = '*'
	}

	return grid
}

// UpdateMap will set the position and rune for all existent parties
func (world *World) UpdateMap() {

	// Regerate map from scratch
	// TODO: Optimize me
	grid := map[Tile]rune{}
	for y := 0; y < world.YMax; y++ {
		for x := 0; x < world.XMax; x++ {
			tile := Tile{
				X: x,
				Y: y,
			}
			grid[tile] = '.'
		}
	}

	// Load persistent terrain


	// Move parties
	for _, party := range world.Parties {
		coordinates := Tile{
			X: party.X,
			Y: party.Y,
		}
		grid[coordinates] = party.Rune
	}

	world.Tiles = grid
}
