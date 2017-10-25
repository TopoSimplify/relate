package relate

import (
	"simplex/node"
	"simplex/ctx"
	"simplex/lnr"
)

//geometry relate
func IsGeomRelateValid(self lnr.Linear, hull *node.Node, ctx *ctx.ContextGeometry) bool {
	var seg    =  hull.Segment()
	var subpln = self.Polyline().SubPolyline(hull.Range)

	var ln_geom  = subpln.Geometry
	var seg_geom = seg
	var ctx_geom = ctx.Geom

	var ln_g_inter  = ln_geom.Intersects(ctx_geom)
	var seg_g_inter = seg_geom.Intersects(ctx_geom)

	var bln = true
	if (seg_g_inter && !ln_g_inter)  || (!seg_g_inter && ln_g_inter){
		bln = false
	}
	// both intersects & disjoint
	return bln
}
