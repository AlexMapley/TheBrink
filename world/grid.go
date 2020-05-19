package world

func newGrid(xMax int, yMax int) map[geom.Point]string {
	grid := map[geom.Point]string{}
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			point := geom.Point{
				X: geom.Pt(x),
				Y: geom.Pt(y),
			}
			grid[point] = "0"
		}
	}

}