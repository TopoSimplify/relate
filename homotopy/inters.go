package homotopy

import (
	"sort"
	"fmt"
	"github.com/intdxdt/geom"
)

type Inter struct {
	Intr *geom.InterPoint
	I    int
	J    int
	S    int
}

type Intersects struct {
	Simple     *geom.Segment
	ExtSimple  *geom.Segment
	Intersects []*Inter
	Indices    IntPairs
}

func NewIntersects(coordinates []*geom.Point) *Intersects {
	var simple, extSimple = SimpleSegs(coordinates)
	var tokens = IntersectTokens(extSimple, coordinates)
	updateInterBySidedness(simple, tokens)

	fmt.Println(extSimple.WKT())

	var intersects = &Intersects{
		Simple:     simple,
		ExtSimple:  extSimple,
		Intersects: tokens,
	}
	sort.Sort(intersects)
	return intersects
}

func SimpleSegs(coordinates []*geom.Point) (*geom.Segment, *geom.Segment) {
	var n = len(coordinates) - 1
	var a, b = coordinates[0], coordinates[n]
	var simple = geom.NewSegment(a, b)
	var bounds = boundingBox(coordinates)
	var ax, bx = extendEndPoints(a, b, bounds)
	var extSimple = geom.NewSegment(ax, bx)
	return simple, extSimple
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
