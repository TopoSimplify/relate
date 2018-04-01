package homotopy

import (
	"github.com/intdxdt/geom"
)

func oddBoundedFaces(coordinates []*geom.Point, intersect *Inter) []*geom.Polygon {
	var a, b []*geom.Point
	var ga, gb *geom.Polygon

	var pt = intersect.Intr.Point
	var i, j = intersect.I, intersect.J

	a = append([]*geom.Point{}, coordinates[:i+1]...)
	a = append(a, pt)
	b = append([]*geom.Point{pt}, coordinates[j:]...)

	ga, gb = geom.NewPolygon(a), geom.NewPolygon(b)

	return []*geom.Polygon{ga, gb}
}
