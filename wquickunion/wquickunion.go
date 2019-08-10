package wquickunion

// UnionFind is the data structed implemented
type UnionFind struct {
	parent []int
	size   []int
	count  int
}

// New returns an initialized uf
func New(size int) *UnionFind {
	return new(UnionFind).init(size)
}

func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)
	uf.count = size
	uf.parent = make([]int, size)
	uf.size = make([]int, size)
	for i := 0; i < size; i++ {
		uf.parent[i] = i
		uf.size[i] = i
	}
	return uf
}

// Find returns the component identifier
func (uf *UnionFind) Find(p int) int {
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

// Union merges the component
func (uf *UnionFind) Union(p int, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if uf.size[p] < uf.size[q] {
		uf.parent[proot] = qroot
		uf.size[qroot] += uf.size[proot]
	} else {
		uf.parent[qroot] = proot
		uf.size[proot] += uf.size[qroot]
	}
}

// Connected returns true if two sites are in the same component
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}
