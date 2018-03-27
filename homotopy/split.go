package homotopy

import (
	"github.com/intdxdt/geom"
	"simplex/side"
)

func intersectTokens(segment *geom.Segment, coordinates []*geom.Point) []*Inter {
	var curSide, prevSide *side.Side
	var ln *geom.Segment
	var c *geom.Point
	var intersections []*Inter

	for idx, n := 1, len(coordinates)-1; idx < n; idx++ {
		var i, j = idx-1, idx
		c = coordinates[idx]
		curSide = segment.SideOf(c)
		if (prevSide != nil) && !(curSide.IsSameSide(prevSide)) {
			ln = geom.NewSegment(coordinates[i], coordinates[j])
			if segment.Intersects(ln) {
				pt := segment.Intersection(ln)
				intersections = append(
					intersections, &Inter{
						Point: pt[0], I: i, J: j, //sort pnts and pick first
					})
			}
		}
		prevSide = curSide
	}
	return intersections
}
