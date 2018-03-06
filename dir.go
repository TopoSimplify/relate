package relate

import (
	"simplex/ctx"
	"simplex/node"
)

//direction relate
func IsDirRelateValid(hull *node.Node, ctx *ctx.ContextGeometries) bool {
	return Homotopy(hull.Coordinates(), ctx)
}
