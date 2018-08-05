package relate

import (
	"github.com/TopoSimplify/pln"
			"github.com/intdxdt/geom"
	)


func polyln(wkt string) *pln.Polyline {
	return pln.New(geom.NewLineStringFromWKT(wkt).Coordinates)
}

