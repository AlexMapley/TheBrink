package world

import "golang.org/x/mobile/geom"

type MetaGame struct {
	Day int
}

type Map struct {
	XMax int
	YMax int
	Grid []geom.Point
}