package relate

import (
	"github.com/intdxdt/geom"
	"simplex/seg"
)

func reverseCoordinates(coordinates []*geom.Point) []*geom.Point {
	for i, j := 0, len(coordinates)-1; i < j; i, j = i+1, j-1 {
		coordinates[i], coordinates[j] = coordinates[j], coordinates[i]
	}
	return coordinates
}

//ring : P0 == Pn
func isRing(coordinates []*geom.Point) bool {
	if len(coordinates) < 2 {
		return false
	}
	return coordinates[0].Equals2D(coordinates[len(coordinates)-1])
}


func lineSegments(coordinates []*geom.Point) []*seg.Seg {
	var i, j int
	var segments = make([]*seg.Seg, 0, len(coordinates)-1)
	for i = 0; i < len(coordinates)-1; i++ {
		j = i + 1
		segments = append(segments, seg.NewSeg(coordinates[i], coordinates[j], i, j))
	}
	return segments
}
