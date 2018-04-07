package homotopy

import (
	"simplex/ctx"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/rtree"
)

func chainDeformation(coordinates []*geom.Point, contexts *ctx.ContextGeometries) *Chain {
	var db = contextDB(contexts)
	var chain = NewChain(coordinates)

	var deformable = true

	for deformable && chain.size > 2 {
		deformable = false
		var link = chain.link

		for link != nil {
			if collapseVertex(link, db) {
				remove(link)
				chain.size += -1
				deformable = true
			}
			link = link.next
		}
	}
	return chain
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
		b[geom.X], b[geom.Y],
	).ExpandIncludeXY(
		c[geom.X], c[geom.Y],
	)

	var neighbours = db.Search(box)
	if len(neighbours) > 0 {
		bln = isTriangleCollapsible(a, b, c, neighbours)
	}
	return bln
}

//find if intersects simple
func isTriangleCollapsible(a, b, c *geom.Point, neighbours []*rtree.Node) bool {
	var bln = true
	var coords = []*geom.Point{a, b, c, a}
	var triangle = geom.NewPolygon(coords)
	for _, node := range neighbours {
		n := node.GetItem().(*ctx.ContextGeometry)
		if triangle.Intersects(n.Geom) {
			//[a,b,c,a]->[0,1,2,3]
			ab := geom.NewLineString(coords[0:2]) //[a,b]->[0,1]
			bc := geom.NewLineString(coords[1:3]) //[b,c]->[1,2]
			ac := geom.NewLineString(coords[2:4]) //[c,a]->[2,3]

			bln = (ab.Intersects(n.Geom) || bc.Intersects(n.Geom)) && ac.Intersects(n.Geom)
			if !bln {
				break
			}
		}
	}
	return bln
}
