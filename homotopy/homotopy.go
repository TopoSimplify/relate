package homotopy

import (
	"github.com/intdxdt/geom"
	"simplex/ctx"
)

//Homotopy Relate
func Homotopy(coordinates []*geom.Point, contexts *ctx.ContextGeometries) bool {
	coordinates = cloneCoordinates(coordinates)
	var bln = true
	var geometries = contexts.DataView()

	var homo = SimpleBounds(coordinates, false)

	bln = disjointContextRelate(homo, geometries)
	if !bln && homo.IsOdd() {
		homo = SimpleBounds(coordinates, true)
		bln = disjointContextRelate(homo, geometries)
	}
	return bln
}

func disjointContextRelate(homo *SBounds, geometries []*ctx.ContextGeometry) bool {
	var bln = true
	var simple = homo.Intersects.Simple
	var linestring = homo.Polyline
outer:
	for i, n := 0, len(geometries); bln && i < n; i++ {
		g := geometries[i]
		if simple.Intersects(g.Geom) && linestring.Intersects(g.Geom) {
			continue //bln = true : continue
		}
		for _, ply := range homo.Regions {
			if ply != nil && ply.Intersects(g.Geom) {
				bln = false
				continue outer
			}
		}
	}
	return bln
}
