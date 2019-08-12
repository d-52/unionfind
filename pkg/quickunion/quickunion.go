package quickunion

// UnionFind is the data structed implemented
type UnionFind struct {
	ID []int
}

// New returns an initialized uf
func New(size int) *UnionFind {
	return new(UnionFind).init(size)
}

func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)
	uf.ID = make([]int, size)
	for i := 0; i < size; i++ {
		uf.ID[i] = i
	}
	return uf
}

// Find returns the component identifier
func (uf *UnionFind) Find(p int) int {
	return uf.ID[p]
}

// Union merges the component
func (uf *UnionFind) Union(p int, q int) {
	pID := uf.root(p)
	qID := uf.root(q)
	uf.ID[pID] = qID

}

// Connected returns true if two sites are in the same component
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.root(p) == uf.root(q)
}

// find returns the root component
func (uf *UnionFind) root(i int) int {
	for i != uf.ID[i] {
		i = uf.ID[i]
	}
	return i
}
