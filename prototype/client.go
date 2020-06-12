package prototype

type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{
		showcase: make(map[string]Product),
	}
}

func (m *Manager) Register(name string, proto Product) {
	m.showcase[name] = proto
}

func (m *Manager) Create(protoName string) Product {
	p := m.showcase[protoName]
	return p.CreateClone()
}
