package fsm

type Machine struct {
	State   *State
	Roadmap map[string]map[string]string
	States  map[string]*State
}

func (m *Machine) ProcessInp(inp interface{}) {
	upcode := m.State.ProcessInp(inp)
	m.State = m.States[m.Roadmap[m.State.Id][upcode]]
}

type State struct {
	ProcessInp func(inp interface{}) string
	Id         string
}
