package relate

import (
	"simplex/pln"
	"github.com/intdxdt/geom"
	"simplex/node"
	"simplex/ctx"
	"simplex/lnr"
)

//direction relate
func IsDirRelateValid(self lnr.Linear, hull *node.Node, ctx *ctx.ContextGeometry) bool {
	var poly = self.Polyline()
	var subpln = poly.SubPolyline(hull.Range)
	var segment = pln.New([]*geom.Point{
		poly.Coordinates[hull.Range.I()],
		poly.Coordinates[hull.Range.J()],
	})

	var lnRelate = DirectionRelate(subpln, ctx.Geom)
	var segRelate = DirectionRelate(segment, ctx.Geom)

	return lnRelate == segRelate
}
