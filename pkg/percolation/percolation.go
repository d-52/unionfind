package percolation

import (
	"github.com/d-52/unionfind/wquickunion"
)

// Percolation is a dynamic connection application
type Percolation struct {
	UFUp         *wquickunion.UnionFind
	UFDown       *wquickunion.UnionFind
	OpenSites    [][]bool
	NumOpenSites int
	GridSize     int
}

// Site is the single unit on grid
type Site struct {
	Row, Col int
}

// New returns an initialized percolation struct
func New(size int) *Percolation {
	return new(Percolation).init(size)
}

func (p *Percolation) init(GridSize int) *Percolation {
	p = new(Percolation)
	p.UFUp = wquickunion.New(GridSize*GridSize + 1) // without bottom index
	p.UFDown = wquickunion.New(GridSize*GridSize + 2)
	p.GridSize = GridSize
	p.OpenSites = make([][]bool, GridSize)
	for j := range p.OpenSites {
		p.OpenSites[j] = make([]bool, GridSize)
	}
	return p
}

// calls union on uf s
func (p *Percolation) union(x, y int) {
	p.UFUp.Union(x, y)
	p.UFDown.Union(x, y)
}

// Connect neightbours
func (p *Percolation) connect(s Site) {
	if !s.isValid(p.GridSize) {
		return
	}
	siteIndex := s.ToIndex(p.GridSize)
	// Connect top
	if s.Row > 1 && p.IsOpen(Site{Row: s.Row - 1, Col: s.Col}) {
		p.union(p.xyToIndex(s.Row-1, s.Col), siteIndex)
	}
	// Connect Bottom
	if s.Row < p.GridSize && p.IsOpen(Site{Row: s.Row + 1, Col: s.Col}) {
		p.union(p.xyToIndex(s.Row+1, s.Col), siteIndex)
	}
	// Connect Left
	if s.Col > 1 && p.IsOpen(Site{Row: s.Row, Col: s.Col - 1}) {
		p.union(p.xyToIndex(s.Row, s.Col-1), siteIndex)
	}
	// Connect Right
	if s.Col < p.GridSize && p.IsOpen(Site{Row: s.Row, Col: s.Col + 1}) {
		p.union(p.xyToIndex(s.Row, s.Col+1), siteIndex)
	}
	// Connect Top
	if s.Row == 1 {
		p.union(0, siteIndex)
	}
	// Connect Bottom
	if s.Row == p.GridSize {
		p.UFDown.Union(p.GridSize*p.GridSize+1, siteIndex)
	}
}

// Open opens
func (p *Percolation) Open(s Site) {
	if !s.isValid(p.GridSize) {
		return
	}
	if !p.IsOpen(s) {
		p.NumOpenSites++
		p.OpenSites[s.Row-1][s.Col-1] = true
		p.connect(s)
	}
}

// IsOpen check
func (p Percolation) IsOpen(s Site) bool {
	if !s.isValid(p.GridSize) {
		return false
	}
	return p.OpenSites[s.Row-1][s.Col-1]
}

// IsFull check
func (p Percolation) IsFull(s Site) bool {
	if !s.isValid(p.GridSize) {
		return false
	}
	return p.UFUp.Connected(s.ToIndex(p.GridSize), 0)
}

// Percolates make percolation check
func (p Percolation) Percolates() bool {
	return p.UFDown.Connected(0, p.GridSize*p.GridSize+1)
}

// xyToIndex conversion
func (p Percolation) xyToIndex(row, col int) int {
	return 1 + (row-1)*p.GridSize + (col - 1)
}

// Site Validator
func (s Site) isValid(size int) bool {
	if s.Row <= 0 || s.Row > size || s.Col <= 0 || s.Col > size {
		return false
	}
	return true
}

// ToIndex conversion
func (s Site) ToIndex(GridSize int) int {
	return 1 + (s.Row-1)*GridSize + (s.Col - 1)
}
