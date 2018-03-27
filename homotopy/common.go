package homotopy

import "github.com/intdxdt/geom"

func cloneCoordinates(coordinates []*geom.Point) []*geom.Point {
	var n = len(coordinates)
	var clone = make([]*geom.Point, n, n)
	copy(clone, coordinates)
	return clone
}

func reverseCoordinates(coordinates []*geom.Point) []*geom.Point {
	for i, j := 0, len(coordinates)-1; i < j; i, j = i+1, j-1 {
		coordinates[i], coordinates[j] = coordinates[j], coordinates[i]
	}
	return coordinates
}


