package homotopy

import "github.com/intdxdt/geom"

type Inter struct {
	Point *geom.Point
	I     int
	J     int
}

type Intersects struct {
	Origin     *geom.Point
	Intersects []*Inter
}

func (o *Intersects) Len() int {
	return len(o.Intersects)
}

func (o *Intersects) Swap(i, j int) {
	o.Intersects[i], o.Intersects[j] = o.Intersects[j], o.Intersects[i]
}

func (o *Intersects) Less(i, j int) bool {
	var a, b = o.Intersects[i].Point, o.Intersects[j].Point
	return o.Origin.Magnitude(a) < o.Origin.Magnitude(b)
}
