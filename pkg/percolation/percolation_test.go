package percolation_test

import (
	"testing"

	p "github.com/d-52/unionfind/percolation"
)

var intSmall = 10
var intMedium = 100
var intLarge = 1000

func TestNew(t *testing.T) {
	per := p.New(intSmall)
	if per.GridSize != intSmall {
		t.Errorf("New(5) = lenght expected to be %d", intSmall)
	}
}
func TestOpen(t *testing.T) {
	per := p.New(intSmall)

	site := p.Site{1, 2}
	per.Open(site)

	site = p.Site{10, 10}
	per.Open(site)

	site = p.Site{-1, -1}
	per.Open(site)

	site = p.Site{11, 11}
	per.Open(site)
}
func TestIsOpen(t *testing.T) {
	per := p.New(intSmall)

	if got := per.IsOpen(p.Site{0, 0}); got != false {
		t.Errorf("per.IsOpen({0, 0}) = %t, want %t", got, false)
	}
	if got := per.IsOpen(p.Site{1, 2}); got != false {
		t.Errorf("per.IsOpen({1, 2}) = %t, want %t", got, false)
	}
	if got := per.IsOpen(p.Site{21, 22}); got != false {
		t.Errorf("per.IsOpen({1, 2}) = %t, want %t", got, false)
	}

	per.Open(p.Site{-1, -2})
	if got := per.IsOpen(p.Site{-1, -2}); got != false {
		t.Errorf("per.IsOpen({1, 2}) = %t, want %t", got, false)
	}

	per.Open(p.Site{1, 2})
	if got := per.IsOpen(p.Site{0, 0}); got != false {
		t.Errorf("per.IsOpen({0, 0}) = %t, want %t", got, false)
	}
	if got := per.IsOpen(p.Site{1, 2}); got != true {
		t.Errorf("per.IsOpen({1, 2}) = %t, want %t", got, true)
	}

	per.Open(p.Site{1, 1})
	if got := per.IsOpen(p.Site{1, 1}); got != true {
		t.Errorf("per.IsOpen({1, 1}) = %t, want %t", got, true)
	}
	per.Open(p.Site{10, 10})
	if got := per.IsOpen(p.Site{10, 10}); got != true {
		t.Errorf("per.IsOpen({10, 10}) = %t, want %t", got, true)
	}
}

func TestUnion(t *testing.T) {
	per := p.New(intSmall)
	per.Open(p.Site{5, 5})
	per.Open(p.Site{5, 6})
	if got := per.UFUp.Connected(
		p.Site{5, 5}.ToIndex(intSmall),
		p.Site{5, 6}.ToIndex(intSmall)); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
	per.Open(p.Site{4, 5})
	if got := per.UFUp.Connected(
		p.Site{5, 5}.ToIndex(intSmall),
		p.Site{4, 5}.ToIndex(intSmall)); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
	per.Open(p.Site{5, 6})
	if got := per.UFUp.Connected(
		p.Site{5, 5}.ToIndex(intSmall),
		p.Site{5, 6}.ToIndex(intSmall)); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
	per.Open(p.Site{5, 4})
	if got := per.UFUp.Connected(
		p.Site{5, 5}.ToIndex(intSmall),
		p.Site{5, 4}.ToIndex(intSmall)); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
	per.Open(p.Site{10, 5})
	if got := per.UFDown.Connected(
		p.Site{10, 5}.ToIndex(intSmall),
		101); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
}

func TestPercolates(t *testing.T) {
	per := p.New(intSmall)

	if got := per.Percolates(); got != false {
		t.Errorf("per.Percolates() = %t, want %t", got, false)
	}

	for i := 1; i < 11; i++ {
		per.Open(p.Site{i, 1})
	}

	if got := per.Percolates(); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}

	per = p.New(intSmall)

	if got := per.Percolates(); got != false {
		t.Errorf("per.Percolates() = %t, want %t", got, false)
	}

	for i := 6; i < 11; i++ {
		per.Open(p.Site{i, 1})
	}

	if got := per.Percolates(); got != false {
		t.Errorf("per.Percolates() = %t, want %t", got, false)
	}

	for i := 1; i < 6; i++ {
		per.Open(p.Site{i, 1})
	}
	if got := per.Percolates(); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
}
