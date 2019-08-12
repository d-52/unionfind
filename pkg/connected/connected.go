package connected

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/d-52/unionfind/wquickunion"
)

var ()

// Population has many People
type Population struct {
	People map[int]*Person
	conn   *wquickunion.UnionFind
	Size   int
}

// New returns an initialized Group
func New(Size int) *Population {
	return new(Population).init(Size)
}

func (g *Population) init(Size int) *Population {
	g = new(Population)
	g.Size = Size
	g.People = make(map[int]*Person)
	g.conn = wquickunion.New(Size)
	for i := 0; i < g.Size; i++ {
		g.People[i] = &Person{ID: i}
		g.People[i].connections = make(map[int]struct{})
	}
	return g
}

// MakeRandomConnections for testing
// Logs connections into a file
func (g *Population) MakeRandomConnections(count int) {
	file, _ := os.OpenFile("connections.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	l := log.New(file, "", 0)

	defer file.Close()

	if count > g.Size*g.Size {
		return
	}

	rand.Seed(42)
	t := time.Date(1990, 1, 1, 00, 00, 00, 00, time.UTC)

	for i := 0; i < count; {
		p1 := rand.Intn(g.Size)
		p2 := rand.Intn(g.Size)
		if p1 != p2 && !g.People[p1].IsConnected(p2) {
			g.People[p1].Connect(p2)
			log := fmt.Sprintf("%s %d %d", t.Format("2006-01-02T15:04:05"), p1, p2)
			l.Println(log)
			t = t.Add(time.Hour*6 + time.Minute*37 + time.Second*42)
			i++
		}
	}
}

// FindConnectionTime reads logs and finds the first time
// everyone is connected
func (g *Population) FindConnectionTime() {

	file, _ := os.OpenFile("connections.log", os.O_RDONLY, 0644)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		t := s[0]
		p, _ := strconv.Atoi(s[1])
		q, _ := strconv.Atoi(s[2])
		g.conn.Union(p, q)
		if g.areAllConnected() {
			fmt.Println(t)
			return
		}
	}
}

func (g Population) areAllConnected() bool {
	for i := 1; i < g.Size; i++ {
		if g.conn.Connected(0, i) == false {
			fmt.Println("yo")
			return false
		}
	}

	return true
}
