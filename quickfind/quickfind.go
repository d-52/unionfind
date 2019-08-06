package quickfind

// UnionFind is the data structed implemented
type UnionFind struct {
	id []int
}

// New returns an initialized uf
func New(size int) *UnionFind {
	return new(UnionFind).init(size)
}

func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)
	for i := 0; i < size; i++ {
		uf.id = append(uf.id, i)
	}
	return uf
}

// Find returns the component identifier
func (uf *UnionFind) Find(p int) int {
	return uf.id[p]
}

// Union merges the component
func (uf *UnionFind) Union(p int, q int) {
	pID := uf.id[p]
	qID := uf.id[q]
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
}

// Connected returns true if two sites are in the same component
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.id[p] == uf.id[q]
}
