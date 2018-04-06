package homotopy

import (
	"github.com/intdxdt/geom"
	"simplex/ctx"
)

//Homotopy Relate
func Homotopy(coordinates []*geom.Point, contexts *ctx.ContextGeometries) bool {
	var bln bool
	var ch = chainDeformation(coordinates, contexts)
	var n = len(coordinates)-1
	if ch.size == 2 {
		var a , b = ch.link.Point, ch.link.next.Point
		bln = coordinates[0].Equals2D(a) && coordinates[n].Equals2D(b)
	}
	return bln
}
