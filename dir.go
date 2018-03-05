package relate

import (
	"simplex/ctx"
	"simplex/node"
)

//direction relate
func IsDirRelateValid(hull *node.Node, ctx *ctx.ContextGeometry) bool {
	var subpln = hull.Polyline
	var segment = hull.SegmentAsPolyline()

	var lnRelate = QuadRelate(subpln, ctx.Geom)
	var segRelate = QuadRelate(segment, ctx.Geom)

	return lnRelate == segRelate
}
