package relate

import (
	"simplex/ctx"
	"simplex/node"
)

//direction relate
func IsDirRelateValid(hull *node.Node, ctx *ctx.ContextGeometry) bool {
	return Homotopy(hull.Coordinates(), ctx.Geom)
}
