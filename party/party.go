package party

func (*party Party) Move(x,y int) {
	party.Coodinates = geom.Point{
		X: geom.Pt(party.Coodinates.X - x),
		Y: geom.Pt(party.Coodinates.Y + y),
	}
}