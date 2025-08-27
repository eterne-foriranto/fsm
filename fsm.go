package fsm

type Machine struct {
	State   *State
	Roadmap map[string]map[string]string
	States  map[string]*State
}

func (m *Machine) ProcessInp(inp interface{}) {
	state := m.State
	if state != nil {
		processFun := state.ProcessInp
		if processFun != nil {
			upcode := processFun(inp)
			m.State = m.States[m.Roadmap[state.Id][upcode]]
		}
	}
}

type State struct {
	ProcessInp func(inp interface{}) string
	Id         string
}
