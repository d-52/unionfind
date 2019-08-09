package percolation

import (
	"github.com/d-52/unionfind/quickfind"
)

// Percolation is a dynamic connection application
type Percolation struct {
	ufup         *quickfind.UnionFind
	ufdown       *quickfind.UnionFind
	openSites    [][]bool
	NumOpenSites int
	gridSize     int
}

// Site is the single unit on grid
type Site struct {
	Row, Col int
}

// New returns an initialized percolation struct
func New(size int) *Percolation {
	return new(Percolation).init(size)
}

func (p *Percolation) init(gridSize int) *Percolation {
	p = new(Percolation)
	p.ufup = quickfind.New(gridSize*gridSize + 1) // without bottom index
	p.ufdown = quickfind.New(gridSize*gridSize + 2)
	p.gridSize = gridSize
	p.openSites = make([][]bool, gridSize)
	for j := range p.openSites {
		p.openSites[j] = make([]bool, gridSize)
	}
	return p
}

// Connect neightbours
func (p *Percolation) connect(s Site) {
	if !s.isValid(p.gridSize) {
		return
	}
	siteIndex := s.toIndex(p.gridSize)
	// Connect top
	if s.Row > 1 && p.IsOpen(Site{Row: s.Row - 1, Col: s.Col}) {
		p.ufup.Union(p.xyToIndex(s.Row-1, s.Col), siteIndex)
		p.ufdown.Union(p.xyToIndex(s.Row-1, s.Col), siteIndex)
	}
	// Connect Bottom
	if s.Row < p.gridSize && p.IsOpen(Site{Row: s.Row + 1, Col: s.Col}) {
		p.ufup.Union(p.xyToIndex(s.Row+1, s.Col), siteIndex)
		p.ufdown.Union(p.xyToIndex(s.Row+1, s.Col), siteIndex)
	}
	// Connect Left
	if s.Col > 1 && p.IsOpen(Site{Row: s.Row, Col: s.Col - 1}) {
		p.ufup.Union(p.xyToIndex(s.Row, s.Col-1), siteIndex)
		p.ufdown.Union(p.xyToIndex(s.Row, s.Col-1), siteIndex)
	}
	// Connect Right
	if s.Col < p.gridSize && p.IsOpen(Site{Row: s.Row, Col: s.Col + 1}) {
		p.ufup.Union(p.xyToIndex(s.Row, s.Col+1), siteIndex)
		p.ufdown.Union(p.xyToIndex(s.Row, s.Col+1), siteIndex)
	}
	// Connect Top
	if s.Row == 1 {
		p.ufup.Union(0, siteIndex)
		p.ufdown.Union(0, siteIndex)
	}
	// Connect Bottom
	if s.Row == p.gridSize {
		p.ufdown.Union(p.gridSize*p.gridSize+1, siteIndex)
	}
}

// Open opens
func (p *Percolation) Open(s Site) {
	if !s.isValid(p.gridSize) {
		return
	}
	if !p.IsOpen(s) {
		p.NumOpenSites++
		p.openSites[s.Row-1][s.Col-1] = true
		p.connect(s)
	}
}

// IsOpen check
func (p Percolation) IsOpen(s Site) bool {
	if !s.isValid(p.gridSize) {
		return false
	}
	return p.openSites[s.Row-1][s.Col-1]
}

// IsFull check
func (p Percolation) IsFull(s Site) bool {
	if !s.isValid(p.gridSize) {
		return false
	}
	return p.ufup.Connected(s.toIndex(p.gridSize), 0)
}

// Percolates make percolation check
func (p Percolation) Percolates() bool {
	return p.ufdown.Connected(0, p.gridSize*p.gridSize+1)
}

// xyToIndex conversion
func (p Percolation) xyToIndex(row, col int) int {
	return 1 + (row-1)*p.gridSize + (col - 1)
}

// Site Validator
func (s Site) isValid(size int) bool {
	if s.Row <= 0 || s.Row > size || s.Col <= 0 || s.Col > size {
		return false
	}
	return true
}

// toIndex conversion
func (s Site) toIndex(gridSize int) int {
	return 1 + (s.Row-1)*gridSize + (s.Col - 1)
}
