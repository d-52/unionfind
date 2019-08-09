package percolation

import "testing"

var intSmall = 10
var intMedium = 100
var intLarge = 1000

func TestNew(t *testing.T) {
	per := New(intSmall)
	if per.gridSize != intSmall {
		t.Errorf("New(5) = lenght expected to be %d", intSmall)
	}
}
func TestOpen(t *testing.T) {
	per := New(intSmall)

	site := Site{1, 2}
	per.Open(site)

	site = Site{10, 10}
	per.Open(site)

	site = Site{-1, -1}
	per.Open(site)

	site = Site{11, 11}
	per.Open(site)
}
func TestIsOpen(t *testing.T) {
	per := New(intSmall)

	if got := per.IsOpen(Site{0, 0}); got != false {
		t.Errorf("per.IsOpen({0, 0}) = %t, want %t", got, false)
	}
	if got := per.IsOpen(Site{1, 2}); got != false {
		t.Errorf("per.IsOpen({1, 2}) = %t, want %t", got, false)
	}
	if got := per.IsOpen(Site{21, 22}); got != false {
		t.Errorf("per.IsOpen({1, 2}) = %t, want %t", got, false)
	}

	per.Open(Site{-1, -2})
	if got := per.IsOpen(Site{-1, -2}); got != false {
		t.Errorf("per.IsOpen({1, 2}) = %t, want %t", got, false)
	}

	per.Open(Site{1, 2})
	if got := per.IsOpen(Site{0, 0}); got != false {
		t.Errorf("per.IsOpen({0, 0}) = %t, want %t", got, false)
	}
	if got := per.IsOpen(Site{1, 2}); got != true {
		t.Errorf("per.IsOpen({1, 2}) = %t, want %t", got, true)
	}

	per.Open(Site{1, 1})
	if got := per.IsOpen(Site{1, 1}); got != true {
		t.Errorf("per.IsOpen({1, 1}) = %t, want %t", got, true)
	}
	per.Open(Site{10, 10})
	if got := per.IsOpen(Site{10, 10}); got != true {
		t.Errorf("per.IsOpen({10, 10}) = %t, want %t", got, true)
	}
}

func TestUnion(t *testing.T) {
	per := New(intSmall)
	per.Open(Site{5, 5})
	per.Open(Site{5, 6})
	if got := per.ufup.Connected(
		Site{5, 5}.toIndex(intSmall),
		Site{5, 6}.toIndex(intSmall)); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
	per.Open(Site{4, 5})
	if got := per.ufup.Connected(
		Site{5, 5}.toIndex(intSmall),
		Site{4, 5}.toIndex(intSmall)); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
	per.Open(Site{5, 6})
	if got := per.ufup.Connected(
		Site{5, 5}.toIndex(intSmall),
		Site{5, 6}.toIndex(intSmall)); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
	per.Open(Site{5, 4})
	if got := per.ufup.Connected(
		Site{5, 5}.toIndex(intSmall),
		Site{5, 4}.toIndex(intSmall)); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
	per.Open(Site{10, 5})
	if got := per.ufdown.Connected(
		Site{10, 5}.toIndex(intSmall),
		101); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
}

func TestPercolates(t *testing.T) {
	per := New(intSmall)

	if got := per.Percolates(); got != false {
		t.Errorf("per.Percolates() = %t, want %t", got, false)
	}

	for i := 1; i < 11; i++ {
		per.Open(Site{i, 1})
	}

	if got := per.Percolates(); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}

	per = New(intSmall)

	if got := per.Percolates(); got != false {
		t.Errorf("per.Percolates() = %t, want %t", got, false)
	}

	for i := 6; i < 11; i++ {
		per.Open(Site{i, 1})
	}

	if got := per.Percolates(); got != false {
		t.Errorf("per.Percolates() = %t, want %t", got, false)
	}

	for i := 1; i < 6; i++ {
		per.Open(Site{i, 1})
	}
	if got := per.Percolates(); got != true {
		t.Errorf("per.Percolates() = %t, want %t", got, true)
	}
}
