package quickfind_test

import (
	"testing"

	. "github.com/d-52/unionfind/quickfind"
)

var intSmall = 10
var intMedium = 1000
var intLarge = 100000

func TestNew(t *testing.T) {
	uf := New(intSmall)
	if len(uf.ID) != intSmall {
		t.Errorf("New(5) = lenght expected to be %d", intSmall)
	}
}

func TestFind(t *testing.T) {
	uf := New(intSmall)
	cases := []int{0, 5, 9}
	for _, c := range cases {
		got := uf.Find(c)
		if got != c {
			t.Errorf("Find(%d) = %d; want %d", c, got, c)
		}
	}
}

func TestUnion(t *testing.T) {
	uf := New(intSmall)
	uf.Union(0, 1)
	got := uf.Find(0)
	if got != 1 {
		t.Errorf("Union(0, 1) = Should change #%d to 1; got %d", 0, got)
	}
}

func TestConnected(t *testing.T) {
	uf := New(intSmall)
	uf.Union(0, 1)
	got := uf.Connected(0, 1)
	if !got {
		t.Errorf("Connected(0, 1) = True; got %t", got)
	}
}

func BenchmarkNewSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(intSmall)
	}
}

func BenchmarkNewMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(intMedium)
	}
}

func BenchmarkNewLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(intLarge)
	}
}

func BenchmarkFindSmall(b *testing.B) {
	uf := New(intSmall)
	for i := 0; i < b.N; i++ {
		uf.Find(i % 10)
	}
}

func BenchmarkFindMedium(b *testing.B) {
	uf := New(intMedium)
	for i := 0; i < b.N; i++ {
		uf.Find(i % 100)
	}
}

func BenchmarkFindLarge(b *testing.B) {
	uf := New(intLarge)
	for i := 0; i < b.N; i++ {
		uf.Find(i % 1000)
	}
}

func BenchmarkUnionSmall(b *testing.B) {
	uf := New(intSmall)
	for i := 0; i < b.N; i++ {
		uf.Union(i%10, (i+2)%10)
	}
}

func BenchmarkUnionMedium(b *testing.B) {
	uf := New(intMedium)
	for i := 0; i < b.N; i++ {
		uf.Union(i%100, (i+2)%100)
	}
}

func BenchmarkUnionLarge(b *testing.B) {
	uf := New(intLarge)
	for i := 0; i < b.N; i++ {
		uf.Union(i%1000, (i+2)%1000)
	}
}

func BenchmarkConnectedSmall(b *testing.B) {
	uf := New(intSmall)
	for i := 0; i < b.N; i++ {
		uf.Connected(i%10, (i+2)%10)
	}
}

func BenchmarkConnectedMedium(b *testing.B) {
	uf := New(intMedium)
	for i := 0; i < b.N; i++ {
		uf.Connected(i%100, (i+2)%100)
	}
}

func BenchmarkConnectedLarge(b *testing.B) {
	uf := New(intLarge)
	for i := 0; i < b.N; i++ {
		uf.Connected(i%1000, (i+2)%1000)
	}
}
