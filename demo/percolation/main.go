package main

import (
	"fmt"
	"math/rand"
	"time"

	"syscall/js"

	p "github.com/d-52/unionfind/applications/percolation"
)

type dom struct {
	doc  js.Value
	view js.Value
}

type percolationView struct {
	ctx     js.Value
	el      dom
	per     *p.Percolation
	seed    int
	playing bool
	done    chan bool
}

func main() {
	v := &percolationView{}
	p := p.New(16)
	v.per = p
	rand.Seed(42)
	v.start()
}

func (v *percolationView) start() {
	v.done = make(chan bool, 0)

	doc := js.Global().Get("document")
	v.el.view = doc.Call("getElementById", "percolationView")
	v.buildDOM()
	v.simulate()
	<-v.done
}

func (v *percolationView) buildDOM() {
	viewHTML := ""
	for i := 1; i <= 16; i++ {
		rowHTML := ""
		for j := 1; j <= 16; j++ {
			rowHTML += fmt.Sprintf(`<div class="cell" id="cell_%d_%d"></div>`, j, i)
		}
		viewHTML += fmt.Sprintf(`<div class="row">%s</div>`, rowHTML)
	}
	v.el.view.Set("innerHTML", viewHTML)
}

func (v *percolationView) next() {
	row := rand.Intn(16) + 1
	col := rand.Intn(16) + 1

	v.per.Open(p.Site{Row: row, Col: col})
	v.openCell(row, col)

	if v.per.IsFull(p.Site{Row: row, Col: col}) {
		v.updatePercolation()
	}
}

func (v *percolationView) openCell(row int, col int) {
	id := fmt.Sprintf("cell_%d_%d", row, col)
	cell := js.Global().Get("document").Call("getElementById", id)
	cell.Get("classList").Call("add", "open")
	opensites := js.Global().Get("document").Call("getElementById", "opensites")
	opensites.Set("innerHTML", v.per.NumOpenSites)
}

func (v *percolationView) updatePercolation() {
	js.Global().Get("console").Call("log", "updating full")
	site := p.Site{}
	for i := 1; i <= 16; i++ {
		for j := 1; j <= 16; j++ {
			site.Row, site.Col = i, j
			if v.per.IsFull(site) {
				id := fmt.Sprintf("cell_%d_%d", i, j)
				cell := js.Global().Get("document").Call("getElementById", id)
				cell.Get("classList").Call("remove", "open")
				cell.Get("classList").Call("add", "full")
			}
		}
	}
	if v.per.Percolates() {
		tag := js.Global().Get("document").Call("getElementById", "percolates")
		tag.Set("innerHTML", "True")
		v.done <- true
	}
}

func (v *percolationView) simulate() {
	go func() {
		for {
			next := time.After(time.Second / 5)
			select {
			case <-v.done:
				return
			case <-next:
				v.next()
			}
		}
	}()
}
