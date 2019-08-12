package connected_test

import (
	"testing"

	c "github.com/d-52/unionfind/connected"
)

func TestNewPopulation(t *testing.T) {
	p := c.New(30)
	if got := p.Size; got != 30 {
		t.Errorf("per.Size = %d, want %d", got, 30)
	}
	if got := p.People[1].ID; got != 1 {
		t.Errorf("p.People[1].ID = %d, want %d", got, 1)
	}
	if got := p.People[29].ID; got != 29 {
		t.Errorf("p.People[1].ID = %d, want %d", got, 29)
	}
}

func TestMakeConnections(t *testing.T) {
	p := c.New(30)
	p.MakeRandomConnections(100)
	p.FindConnectionTime()
}
