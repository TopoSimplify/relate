package relate

import (
	"simplex/ctx"
	"github.com/intdxdt/geom"
	"simplex/relate/homotopy"
)

//Homotopy Relate
func Homotopy(coordinates []*geom.Point, contexts *ctx.ContextGeometries) bool {
	return homotopy.Homotopy(coordinates, contexts)
}
