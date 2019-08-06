package percolation

import (
	"fmt"

	"github.com/d-52/unionfind/quickfind"
)

// Percolation is a dynamic connection application
type Percolation struct {
	uf           *quickfind.UnionFind
	openSites    []bool
	numOpenSites int
	size         int
	vtop         int
	vbtm         int
}

// New returns an initialized percolation struct
func New(size int) *Percolation {
	return new(Percolation).init(size)
}

func (p *Percolation) init(size int) *Percolation {
	p = new(Percolation)
	p.uf = quickfind.New(size)
	p.size = size
	p.vtop = size * size
	p.vbtm = size*size + 1
	p.openSites = make([]bool, size)
	return p
}

// IsOpen check
func (p Percolation) IsOpen(row int, col int) (bool, error) {
	err := p.assertCoords(row, col)
	if err != nil {
		return false, err
	}
	return p.openSites[p.xyTo1D(row, col)], err
}

// Connect neightbours
func (p *Percolation) connectNeighbours(row int, col int) {
	// Connect top
	if open, _ := p.IsOpen(row-1, col); row > 1 && open {
		p.uf.Union(p.xyTo1D(row-1, col), p.xyTo1D(row, col))
	}
	// Connect Bottom
	if open, _ := p.IsOpen(row+1, col); row < p.size && open {
		p.uf.Union(p.xyTo1D(row+1, col), p.xyTo1D(row, col))
	}
	// Connect Left
	if open, _ := p.IsOpen(row, col-1); col > 1 && open {
		p.uf.Union(p.xyTo1D(row, col-1), p.xyTo1D(row, col))
	}
	// Connect Right
	if open, _ := p.IsOpen(row, col+1); col < p.size && open {
		p.uf.Union(p.xyTo1D(row, col+1), p.xyTo1D(row, col))
	}
}

// Connect virtuals
func (p *Percolation) connectVirtuals(row int, col int) {
	// vtop
	if row == 1 {
		p.uf.Union(p.vtop, p.xyTo1D(row, col))
	}
	// vbtm
	if row == p.size {
		if p.uf.Connected(p.vtop, p.xyTo1D(row, col)) {
			p.uf.Union(p.vbtm, p.xyTo1D(row, col))
		}
	}
}

// Open opens
func (p *Percolation) Open(row int, col int) {
	if p.assertCoords(row, col) != nil {
		return
	}
	if open, _ := p.IsOpen(row, col); open {
		p.numOpenSites++
		p.openSites[p.xyTo1D(row, col)] = true
		p.connectVirtuals(row, col)
		p.connectNeighbours(row, col)
	}
}

// IsFull check
func (p Percolation) IsFull(row int, col int) (bool, error) {
	err := p.assertCoords(row, col)
	if err != nil {
		return false, err
	}
	return p.uf.Connected(p.xyTo1D(row, col), p.vtop), err
}

// Percolates make percolation check
func (p Percolation) Percolates() bool {
	return p.uf.Connected(p.vtop, p.vbtm)
}

// InvalidCoordError is returned on invalid coords
type InvalidCoordError struct {
	Name  string
	Value int
}

func (e *InvalidCoordError) Error() string {
	return fmt.Sprintf("%v index %d out of bounds",
		e.Name, e.Value)
}

func (p Percolation) assertCoords(row, col int) error {
	if row <= 0 || row > p.size {
		return &InvalidCoordError{
			Name:  "row",
			Value: row,
		}
	}
	if col <= 0 || col > p.size {
		return &InvalidCoordError{
			Name:  "col",
			Value: col,
		}
	}
	return nil
}

func (p Percolation) xyTo1D(row int, col int) int {
	return (row-1)*p.size + (col - 1)
}
