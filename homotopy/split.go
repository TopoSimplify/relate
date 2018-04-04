package homotopy

import (
	"simplex/side"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/vect"
	"github.com/intdxdt/cart"
	"github.com/intdxdt/math"
)

func IntersectTokens(segment *geom.Segment, coordinates []*geom.Point) []*Inter {
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
			if segment.SegSegIntersects(ln) {
				var pt = segment.SegSegIntersection(ln)
				var sideGroup = (curSide.IsLeft() && prevSide.IsOn()) ||
					(curSide.IsOn() && prevSide.IsLeft())
				if !sideGroup {
					intersections = append(
						intersections, &Inter{
							Intr: pt[0], I: i, J: j, //sort pnts and pick first
						})
				}
			}
		}
		prevSide = curSide
	}

	return intersections
}

func updateInterBySidedness(segment *geom.Segment, intersects []*Inter) {
	var lseg, rseg = perpdicularSegs(segment)
	for _, o := range intersects {
		if lseg.SideOf(o.Intr.Point).IsLeft() {
			o.S = -1
		} else if rseg.SideOf(o.Intr.Point).IsRight() {
			o.S = 1
		} else {
			o.S = 0
		}
	}
}

func perpdicularSegs(segment *geom.Segment) (*geom.Segment, *geom.Segment) {
	var a, b = segment.A, segment.B
	var round = func(f float64) float64 {
		return math.Round(f, precision)
	}
	var perpendicular = func(fromEnd bool) *geom.Segment {
		var v = vect.NewVector(a, b)
		var k = a
		var m = -1.0
		if fromEnd {
			k = b
			m = -m
		}
		var ux, uy = cart.Extend(v, 100, math.Deg2rad(90*m), fromEnd)
		var ptA = k.Add(round(ux), round(uy))
		ux, uy = cart.Extend(v, 100, math.Deg2rad(-90*m), fromEnd)
		var ptB = k.Add(round(ux), round(uy))
		return geom.NewSegment(ptA, ptB)
	}

	return perpendicular(false), perpendicular(true)
}
