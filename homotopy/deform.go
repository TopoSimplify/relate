package homotopy

import (
	"fmt"
	"simplex/ctx"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/rtree"
)

func chainDeformation(coordinates []*geom.Point, contexts *ctx.ContextGeometries) *Chain {
	var db = contextDB(contexts)
	var chain = NewChain(coordinates)

	var deformable = true
	var n = chain.size - 1

	for deformable && chain.size > 2 {
		deformable = false
		var link = chain.link

		for link != nil {
			if link.i > 0 && link.i < n {
				if collapseVertex(link, db) {
					remove(link)
					chain.size += -1
					deformable = true
				}
			}
			link = link.next
		}
		printChain(chain)
	}

	return chain
}

func printChain(chain *Chain) {
	var coords []*geom.Point
	var link = chain.link
	for link != nil {
		coords = append(coords, link.Point)
		link = link.next
	}
	fmt.Println(geom.NewLineString(coords).WKT())
}

func contextDB(contexts *ctx.ContextGeometries) *rtree.RTree {
	var db = rtree.NewRTree(4)
	var objects = make([]rtree.BoxObj, 0)
	for _, o := range contexts.DataView() {
		objects = append(objects, o)
	}
	db.Load(objects)
	return db
}

func collapseVertex(v *Vertex, db *rtree.RTree) bool {
	var va, vb, vc = v.prev, v, v.next
	if va == nil || vb == nil || vc == nil {
		return false
	}
	var bln = true
	var a, b, c = va.Point, vb.Point, vc.Point
	var box = a.BBox().ExpandIncludeXY(
		b[0], b[1],
	).ExpandIncludeXY(c[0], c[1])

	var neighbours = db.Search(box)
	if len(neighbours) > 0 {
		var triangle = geom.NewPolygon([]*geom.Point{a, b, c, a})
		if intersectsTriangle(triangle, neighbours) {
			bln = false
		}
	}
	//fmt.Println(geom.NewPolygon([]*geom.Point{a, b, c, a}).WKT())
	return bln
}

//find if intersects simple
func intersectsTriangle(triangle *geom.Polygon, neighbours []*rtree.Node) bool {
	var bln = false
	for _, node := range neighbours {
		n := node.GetItem().(*ctx.ContextGeometry)
		if triangle.Intersects(n.Geom) {
			bln = true
			break
		}
	}
	return bln
}
