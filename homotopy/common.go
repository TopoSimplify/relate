package homotopy

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/vect"
	"github.com/intdxdt/cart"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/math"
)

const precision = 12

func asIndices(vals []*Inter) [][2]int {
	var tokens = make([][2]int, len(vals))
	for i, o := range vals {
		tokens[i] = [2]int{o.I, o.J}
	}
	return tokens
}

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

func round(f float64) float64{
	return math.Round(f, precision)
}

func extendEndPoints(a, b *geom.Point, box *mbr.MBR) (*geom.Point, *geom.Point) {
	var diag = diagonal(box)
	v := vect.NewVector(a, b)
	m := v.Magnitude() + diag
	ux, uy := cart.Extend(v, m, math.Pi, true)
	ux, uy = round(ux), round(uy)
	up := a.Add(cart.NewCoord(ux, uy))

	lx, ly := cart.Extend(v, diag, math.Pi, false)
	lx, ly = round(lx), round(ly)
	lp := a.Add(cart.NewCoord(lx, ly))
	return geom.NewPointXY(lp[0], lp[1]), geom.NewPointXY(up[0], up[1])
}

func diagonal(box *mbr.MBR) float64 {
	return math.Hypot(box.Width(), box.Height())
}

func boundingBox(coordinates []*geom.Point) *mbr.MBR {
	var pt = coordinates[0]
	var box = mbr.NewMBR(pt[0], pt[1], pt[0], pt[1])
	for _, pt = range coordinates[1:] {
		box.ExpandIncludeXY(pt[0], pt[1])
	}
	return box
}
