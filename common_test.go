package relate

import (
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/intdxdt/geom"
)


func polyln(wkt string) *pln.Polyline {
	return pln.New(geom.NewLineStringFromWKT(wkt).Coordinates())
}
//hull geom
func hullGeom(coords []geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0]
	}
	return g
}

func create_hulls(indxs [][]int, coords []geom.Point) []*node.Node {
	var poly  = pln.New(coords)
	var hulls = make([]*node.Node, 0)
	for _, o := range indxs {
		var r = rng.Range(o[0], o[1])
		var n = node.New(poly.SubCoordinates(r), r , hullGeom)
		hulls = append(hulls, n )
	}
	return hulls
}

