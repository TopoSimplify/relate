package homotopy

import (
	"github.com/intdxdt/geom"
)

func SimpleBounds(coordinates []*geom.Point, reverse bool, simpleBounds ...*SBounds) *SBounds {
	var sb = &SBounds{}

	if len(simpleBounds) > 0 {
		sb = simpleBounds[0]
	} else {
		coordinates = cloneCoordinates(coordinates)
		sb.Intersects = NewIntersects(coordinates)
		sb.Polyline = geom.NewLineString(cloneCoordinates(coordinates)) //todo:void clone
	}

	if sb.IsEven() {
		sb.Regions = evenBoundedFaces(coordinates, sb.Intersects)
	} else {
		var n = sb.Intersects.Len() - 1
		if reverse {
			n = 0
		}
		sb.Regions = oddBoundedFaces(coordinates, sb.Intersects.Intersects[n])
	}
	return sb
}

func boundedPolygon(a *geom.Point, coordinates []*geom.Point, b *geom.Point) *geom.Polygon {
	var bounded = make([]*geom.Point, 0, len(coordinates)+2)
	if a != nil {
		bounded = append(bounded, a)
	}
	bounded = append(bounded, coordinates...)
	if b != nil {
		bounded = append(bounded, b)
	}
	if len(bounded) == 2 && bounded[0].Equals2D(bounded[0]) {
		//plotting program will not plot ring of same point < 4 vertices
		bounded = append(bounded, bounded[0], bounded[0])
	}
	return geom.NewPolygon(bounded)
}
