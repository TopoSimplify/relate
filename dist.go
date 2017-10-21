package relate

import (
	"simplex/node"
	"simplex/ctx"
	"simplex/lnr"
)

//distance relate
func IsDistRelateValid(self lnr.Linear, hull *node.Node, ctx *ctx.CtxGeom) bool {
	var mindist = self.Options().MinDist
	var seg     =  hull.Segment()
	var ln_geom = hull.SubPolyline().Geometry

	var seg_geom = seg
	var ctx_geom = ctx.Geom

	var _or = ln_geom.Distance(ctx_geom) // original relate
	var dr  = seg_geom.Distance(ctx_geom) // new relate

	bln := dr >= mindist
	if (!bln) && _or < mindist {//if not bln and _or <= mindist:
		//if original violates constraint, then simple can
		// >= than original or <= original, either way should be true
		// [original & simple] <= mindist, then simple cannot be  simple >= mindist no matter
		// how many vertices introduced
		bln = true
	}
	return bln
}

