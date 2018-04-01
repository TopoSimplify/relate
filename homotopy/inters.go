package homotopy

import (
	"sort"
	"github.com/intdxdt/geom"
	"fmt"
)

type Inter struct {
	Intr *geom.InterPoint
	I    int
	J    int
}

type Intersects struct {
	Simple     *geom.Segment
	Intersects []*Inter
	Indices    IntPairs
}

func NewIntersects(coordinates []*geom.Point) *Intersects {
	var n = len(coordinates) - 1
	var a, b = coordinates[0], coordinates[n]
	var simple = geom.NewSegment(a, b)
	fmt.Println(simple.WKT())

	var tokens = intersectTokens(simple, coordinates)
	if len(tokens)%2 == 0 {
		var bounds = boundingBox(coordinates)
		var ax, bx = extendEndPoints(a, b, bounds)
		var extSimple = geom.NewSegment(ax, bx)
		var extTokens = intersectTokens(extSimple, coordinates)
		if len(extTokens)%2 == 1 { //if odd after ext - make tokens odd
			tokens = extTokens
		}
	}

	var intersects = &Intersects{
		Simple:     simple,
		Intersects: tokens,
	}

	sort.Sort(intersects)
	return intersects
}

func (o *Intersects) evenAggregate() {
	var ranges IntPairs
	if o.Len()%2 == 1 {
		panic("odd number of intersects - homotopy pairing error")
	}

	for i := 0; i < o.Len(); i += 2 {
		pair := NewIntPair(o.Intersects[i], o.Intersects[i+1])
		ranges = append(ranges, pair)
	}
	o.Indices = compactPairs(ranges)
	sort.Sort(o.Indices)
}

//Checks if number of original intersects with simple line even
//excludes endpoints of polyline
func (o *Intersects) IsEven() bool {
	return o.Len()%2 == 0
}

func (o *Intersects) Len() int {
	return len(o.Intersects)
}

func (o *Intersects) Swap(i, j int) {
	o.Intersects[i], o.Intersects[j] = o.Intersects[j], o.Intersects[i]
}

func (o *Intersects) Less(i, j int) bool {
	var a, b = o.Intersects[i], o.Intersects[j]
	if o.IsEven() {
		var origin = o.Simple.A
		return origin.Magnitude(a.Intr.Point) < origin.Magnitude(b.Intr.Point)
	}
	return a.I < b.I
}
