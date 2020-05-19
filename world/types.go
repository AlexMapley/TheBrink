package world

import "golang.org/x/mobile/geom"

type MetaGame struct {
	Day int
}

type World struct {
	XMax int
	YMax int
	Grid map[string]geom.Point

}