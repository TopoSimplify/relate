package homotopy

import "github.com/intdxdt/geom"

func SimpleDeformation(coordinates []*geom.Point) {
	var n = len(coordinates) - 1
	var simple = geom.NewSegment(coordinates[0], coordinates[n])
}
