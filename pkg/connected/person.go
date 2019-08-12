package connected

// Person is social
type Person struct {
	ID          int
	connections map[int]struct{}
}

// Connect accepts person ID
func (p *Person) Connect(i int) {
	if !p.IsConnected(i) {
		p.connections[i] = struct{}{}
	}
}

// IsConnected returns if connected
func (p *Person) IsConnected(i int) bool {
	if _, ok := p.connections[i]; ok {
		return true
	}
	return false
}
