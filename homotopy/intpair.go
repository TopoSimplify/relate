package homotopy

type IntPair struct {
	a *Inter
	b *Inter
}

func NewIntPair(a, b *Inter) *IntPair {
	if b.I < a.I {
		a, b = b, a
	}
	return &IntPair{a: a, b: b}
}

type IntPairs []*IntPair

func (o IntPairs) Len() int {
	return len(o)
}

func (o IntPairs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o IntPairs) Less(i, j int) bool {
	return o[i].a.I < o[j].a.I
}

func compactPairs(ranges IntPairs) IntPairs {
	dict := make(map[[2]int]*IntPair, ranges.Len())
	for _, o := range ranges {
		dict[[2]int{o.a.I, o.b.J}] = o
	}
	for _, o := range ranges {
		var key = [2]int{o.a.I, o.b.J}
		if _, ok := dict[key]; ok {
			var rm [][2]int
			for dk := range dict {
				if isWithinRange(dk, key) {
					rm = append(rm, key)
				}
			}
			for _, k := range rm {
				delete(dict, k)
			}
		}
	}
	var comp IntPairs
	for _, v := range dict {
		comp = append(comp, v)
	}
	return comp
}

func isWithinRange(self, other [2]int) bool {
	var i, j = self[0], self[1]
	var a, b = other[0], other[1]
	return i < a && b < j
}
