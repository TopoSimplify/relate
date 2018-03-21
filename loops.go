package relate

import (
	"sort"
	"simplex/ctx"
	"simplex/seg"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/rtree"
	"fmt"
)

//Homotopy Relate
func HomotopySimpleLoops(coordinates []*geom.Point, contexts *ctx.ContextGeometries) bool {
	var bln = true
	var n = len(coordinates) - 1
	var g *ctx.ContextGeometry
	var linestring = geom.NewLineString(coordinates)
	var geometries = contexts.DataView()
	var segment = geom.NewSegment(coordinates[0], coordinates[n])
	var backLoops []*geom.Polygon
	var forwardLoops = SimpleLoops(coordinates, false)

	for idx, n := 0, contexts.Len(); bln && idx < n; idx++ {
		g = geometries[idx]
		if segment.Intersects(g.Geom) && linestring.Intersects(g.Geom) {
			continue //bln = true : continue
		}

		bln = disjointLoopRelate(g, forwardLoops)
		if !bln {
			if len(backLoops) == 0 {
				backLoops = SimpleLoops(coordinates, true)
			}
			bln = disjointLoopRelate(g, backLoops)
		}
	}
	return bln
}

func disjointLoopRelate(g *ctx.ContextGeometry, loops []*geom.Polygon) bool {
	var bln = true
	for i, n := 0, len(loops); bln && i < n; i++ {
		bln = !loops[i].Intersects(g.Geom)
	}
	return bln
}

func SimpleLoops(coordinates []*geom.Point, reverse bool) []*geom.Polygon {
	var s *seg.Seg
	var neighbours []*seg.Seg
	var db = rtree.NewRTree(8)

	coordinates = makeRing(coordinates)
	if reverse {
		reverseCoordinates(coordinates)
	}

	var segments = lineSegments(coordinates)
	var iterator = segments[:]
	var loops []*geom.Polygon

	fmt.Println("# segments : ", len(segments))
	var NSegs = 0
	for len(iterator) > 0 {
		NSegs += 1
		s = iterator[0]
		iterator = iterator[1:]

		fmt.Println(s.WKT())

		neighbours = searchIntersects(db, s)
		db.Insert(s)

		for _, o := range neighbours {
			if o.J == s.I {
				continue
			}
			var intersects = s.Segment.Intersection(o.Segment)
			if len(intersects) == 0 {
				continue
			}

			pt := intersects[0]

			if len(iterator) == 0 && coordinates[len(coordinates)-1].Equals2D(pt) {
				continue
			}

			//coordinates of loop
			lp := append([]*geom.Point{pt}, coordinates[o.J: s.J]...)
			if len(lp) > 2 {
				loops = append(loops, geom.NewPolygon(lp))
			}

			rmSegments(db, segments, o.I, s.J)

			var oa = seg.NewSeg(coordinates[o.I], pt, o.I+0, o.I+1)
			var sb = seg.NewSeg(pt, coordinates[s.J], o.I+1, o.I+2)
			var dimA, dimB = !oa.A.Equals2D(oa.B), !sb.A.Equals2D(sb.B)

			//new coordinates and segments
			coordinates = coordinates[:o.I]
			if !dimA && !dimB {
				coordinates = append(coordinates, oa.A)
				iterator = iterator[:0]
			} else if !dimA && dimB {
				coordinates = append(coordinates, sb.A, sb.B)
				iterator = []*seg.Seg{sb}
			} else if dimA && !dimB {
				coordinates = append(coordinates, oa.A, oa.B)
				iterator = []*seg.Seg{oa}
			} else {
				coordinates = append(coordinates, oa.A, oa.B, sb.B)
				iterator = []*seg.Seg{oa, sb}
			}
			//coordinates, segments, iterator = updateChain(coordinates, segments, iterator, o.I, s.J, sb.J)
			var index = len(coordinates) - 1
			for _, ss := range segments[s.J:] {
				ss.I, ss.J, index = index, index+1, index+1
				iterator = append(iterator, ss)
				coordinates = append(coordinates, ss.B)
			}
			segments = append(segments[:o.I:o.I], iterator...)
			break
		}
	}

	if len(coordinates) > 2 {
		loops = append(loops, geom.NewPolygon(coordinates))
	}

	fmt.Println("# -runs- segments : ", NSegs)
	return loops
}

type SegKnn struct {
	*seg.Seg
	Offset float64
}

type Segments []*SegKnn

func (s Segments) Len() int {
	return len(s)
}
func (s Segments) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Segments) Less(i, j int) bool {
	return s[i].Offset < s[j].Offset
}

func makeRing(coordinates []*geom.Point) []*geom.Point {
	var n = len(coordinates)
	var coords = make([]*geom.Point, n, n+1)
	copy(coords, coordinates)
	if n > 1 && !isRing(coords) {
		coords = append(coords, coords[0].Clone())
	}
	return coords
}

func reverseCoordinates(coordinates []*geom.Point) []*geom.Point {
	for i, j := 0, len(coordinates)-1; i < j; i, j = i+1, j-1 {
		coordinates[i], coordinates[j] = coordinates[j], coordinates[i]
	}
	return coordinates
}

//ring : P0 == Pn
func isRing(coordinates []*geom.Point) bool {
	if len(coordinates) < 2 {
		return false
	}
	return coordinates[0].Equals2D(coordinates[len(coordinates)-1])
}

func lineSegments(coordinates []*geom.Point) []*seg.Seg {
	var i, j int
	var segments = make([]*seg.Seg, 0, len(coordinates)-1)
	for i = 0; i < len(coordinates)-1; i++ {
		j = i + 1
		segments = append(segments, seg.NewSeg(coordinates[i], coordinates[j], i, j))
	}
	return segments
}

func searchIntersects(db *rtree.RTree, segment *seg.Seg) []*seg.Seg {
	var neighbours []*SegKnn
	var ptA = segment.A
	for _, n := range db.Search(segment.BBox()) {
		s := n.GetItem().(*seg.Seg)
		if s.Intersects(segment.Segment) {
			neighbours = append(neighbours, &SegKnn{Seg: s, Offset: s.Distance(ptA)})
		}
	}
	if len(neighbours) > 0 {
		sort.Sort(Segments(neighbours))
	}
	var knn = make([]*seg.Seg, 0, len(neighbours))
	for _, s := range neighbours {
		knn = append(knn, s.Seg)
	}
	return knn
}

func rmSegments(db *rtree.RTree, segments []*seg.Seg, i, j int) {
	for _, r := range segments[i:j] {
		db.Remove(r)
	}
}

func updateChain(
	coordinates []*geom.Point, segments, iterator []*seg.Seg,
	startIndex, restIndex int, index int,
) ([]*geom.Point, []*seg.Seg, []*seg.Seg) {

	var j = index
	for _, ss := range segments[restIndex:] {
		ss.I, ss.J, j = j, j+1, j+1
		iterator = append(iterator, ss)
		coordinates = append(coordinates, ss.B)
	}
	segments = append(segments[:startIndex:startIndex], iterator...)
	return coordinates, segments, iterator
}

func subSegments(a, pt, b *geom.Point, i int) (*seg.Seg, *seg.Seg) {
	var oa = seg.NewSeg(a, pt, i, i+1)
	var sb = seg.NewSeg(pt, b, oa.J, oa.J+1)
	return oa, sb
}
