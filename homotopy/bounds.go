package homotopy

import "github.com/intdxdt/geom"

type SBounds struct {
	Regions    []*geom.Polygon
	Intersects *Intersects
	Polyline   *geom.LineString
}

func (sb *SBounds) RegionCount() int {
	return len(sb.Regions)
}
func (sb *SBounds) IsEven() bool {
	return sb.Intersects.IsEven()
}

func (sb *SBounds) IsOdd() bool {
	return !sb.IsEven()
}

