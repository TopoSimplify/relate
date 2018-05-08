package relate

import (
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/homotopy"
	"github.com/intdxdt/geom"
)

//Homotopy Relate
func Homotopy(coordinates []*geom.Point, contexts *ctx.ContextGeometries) bool {
	return homotopy.Homotopy(coordinates, contexts)
}
