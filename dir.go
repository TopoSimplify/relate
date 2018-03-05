package relate

import (
	"simplex/ctx"
	"simplex/node"
	"simplex/lnr"
)

//direction relate
func IsDirRelateValid(self lnr.Polygonal, hull *node.Node, ctx *ctx.ContextGeometry) bool {
	return Homotopy(hull.Coordinates(), ctx.Geom)
}
