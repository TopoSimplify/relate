package homotopy

import "github.com/intdxdt/geom"

func evenBoundedFaces(coordinates []*geom.Point, intersects *Intersects) []*geom.Polygon {
	intersects.evenAggregate()
	var indices = intersects.Indices
	var simpleBounds []*geom.Polygon
	var ptr = 0
	for _, idx := range indices {
		ra, rb := idx.a, idx.b
		pa := boundedPolygon(nil, coordinates[ptr:ra.J], ra.Point)
		pb := boundedPolygon(ra.Point, coordinates[ra.J:rb.J], rb.Point)
		//change ptr origin
		coordinates[rb.I] = rb.Point
		ptr = rb.I
		simpleBounds = append(simpleBounds, pa, pb)
	}
	if ptr < len(coordinates)-1 {
		pa := boundedPolygon(nil, coordinates[ptr:], nil)
		simpleBounds = append(simpleBounds, pa)
	}
	return simpleBounds
}
