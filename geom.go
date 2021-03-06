package relate

import (
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/node"
)

//Node geometric relation to other context geometries
func IsGeomRelateValid(hull *node.Node, contexts *ctx.ContextGeometries) bool {
	var seg = hull.Segment()
	var lnGeom = hull.Polyline.Geometry()
	var segGeom = seg
	var lnGInter, segGInter bool
	var g *ctx.ContextGeometry

	var bln = true
	var geometries = contexts.DataView()

	for i, n := 0, contexts.Len(); bln && i < n; i++ {
		g = geometries[i]
		lnGInter = lnGeom.Intersects(g.Geom)
		segGInter = segGeom.Intersects(g.Geom)

		bln = !((segGInter && !lnGInter) || (!segGInter && lnGInter) )
	}

	return bln
}
