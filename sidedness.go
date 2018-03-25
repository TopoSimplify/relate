package relate

import (
	"simplex/side"
	"simplex/ctx"
	"github.com/intdxdt/geom"
	"fmt"
	"strings"
	"github.com/intdxdt/mbr"
	"math"
	"github.com/intdxdt/vect"
	"github.com/intdxdt/cart"
)

type homoFragments struct {
	a          *geom.Polygon
	b          *geom.Polygon
	simple     *geom.Segment
	linestring *geom.LineString
}

//Homotopy Relate
func Homotopy(coordinates []*geom.Point, contexts *ctx.ContextGeometries) bool {
	coordinates = cloneCoordinates(coordinates)
	var bln = true
	var geometries = contexts.DataView()

	var homo = boundedRegions(coordinates, false)
	printHomo(homo)

	bln = disjointContextRelate(homo, geometries)
	if !bln {
		homo = boundedRegions(coordinates, true)
		printHomo(homo)

		bln = disjointContextRelate(homo, geometries)
	}
	return bln
}

func diagonal(box *mbr.MBR) float64 {
	return math.Hypot(box.Height(), box.Width())
}

func extendEndPoints(a, b *geom.Point, box *mbr.MBR) (*geom.Point, *geom.Point) {
	var diag = 2 * diagonal(box)
	v := vect.NewVector(a, b)
	m := v.Magnitude() + diag
	ux, uy := cart.Extend(v, m, math.Pi, true)
	up := a.Add(cart.NewCoord(ux, uy))

	lx, ly := cart.Extend(v, diag, math.Pi, false)
	lp := a.Add(cart.NewCoord(lx, ly))
	return geom.NewPointXY(lp[0], lp[1]), geom.NewPointXY(up[0], up[1])
}

func printHomo(homo *homoFragments) {
	fmt.Println(strings.Repeat("--", 80))
	if homo.a != nil {
		fmt.Println(homo.a.WKT())
	}

	if homo.b != nil {
		fmt.Println(homo.b.WKT())
	}
}

func disjointContextRelate(homo *homoFragments, geometries []*ctx.ContextGeometry) bool {
	var bln = true
	var simple = homo.simple
	var linestring = homo.linestring
	for i, n := 0, len(geometries); bln && i < n; i++ {
		g := geometries[i]
		if simple.Intersects(g.Geom) && linestring.Intersects(g.Geom) {
			continue //bln = true : continue
		}
		if homo.a != nil && homo.a.Intersects(g.Geom) {
			bln = false
			continue
		}
		if homo.b != nil && homo.b.Intersects(g.Geom) {
			bln = false
			continue
		}
	}
	return bln
}

func boundedRegions(coordinates []*geom.Point, reverse bool) *homoFragments {
	if reverse {
		reverseCoordinates(coordinates)
	}

	var n = len(coordinates)
	var a, b, inters []*geom.Point
	var ga, gb *geom.Polygon

	var linestring = geom.NewLineString(coordinates)
	var ptA, ptB = extendEndPoints(coordinates[0], coordinates[n-1], linestring.BBox())
	var simple = geom.NewSegment(ptA, ptB )
	fmt.Println(simple.WKT())

	var i, j = homoSplit(simple, coordinates)
	if i < 0 && j < 0 {
		ga, gb = geom.NewPolygon(coordinates), nil
		return &homoFragments{
			a:          ga,
			b:          gb,
			simple:     simple,
			linestring: linestring,
		}
	}

	var segment = geom.NewSegment(coordinates[i], coordinates[j])
	inters = simple.Intersection(segment)

	a = append([]*geom.Point{}, coordinates[:i+1]...)
	a = append(a, inters[0])
	b = append([]*geom.Point{inters[0]}, coordinates[j:]...)

	ga, gb = geom.NewPolygon(a), geom.NewPolygon(b)

	return &homoFragments{
		a:          ga,
		b:          gb,
		simple:     simple,
		linestring: linestring,
	}
}

func homoSplit(segment *geom.Segment, coordinates []*geom.Point) (int, int) {
	var i, j = -1, -1
	var curSide, prevSide *side.Side
	var ln *geom.Segment
	var c *geom.Point

	for idx, n := 1, len(coordinates)-1; idx < n; idx++ {
		c = coordinates[idx]
		curSide = segment.SideOf(c)
		if (prevSide != nil) && !(curSide.IsSameSide(prevSide)) {
			ln = geom.NewSegment(coordinates[idx-1], coordinates[idx])
			if segment.Intersects(ln) {
				i, j = idx-1, idx
			}
		}
		prevSide = curSide
	}
	return i, j
}

func cloneCoordinates(coordinates []*geom.Point) []*geom.Point {
	var n = len(coordinates)
	var clone = make([]*geom.Point, n, n)
	copy(clone, coordinates)
	return clone
}
