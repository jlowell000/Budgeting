package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	periodic_flow "jlowell000.github.io/budgeting/internal/model/periodicflow"
)

type FlowListModel struct {
	Flows    []periodic_flow.PeriodicFlow // list of flows
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool
}

func flowListUpdate(msg tea.Msg, m AppModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.FlowList.Choice++
			if m.FlowList.Choice > len(m.FlowList.Flows)-1 {
				m.FlowList.Choice = len(m.FlowList.Flows) - 1
			}
		case "k", "up":
			m.FlowList.Choice--
			if m.FlowList.Choice < 0 {
				m.FlowList.Choice = 0
			}

		case "b", "backspace":
			m.Main.Chosen = false
		case "enter":
			m.FlowList.Chosen = true
			// return m, nil
		}
	}

	return m, nil
}

func flowListView(m AppModel) string {
	c := m.FlowList.Choice
	// The header
	tpl := "Viewing Periodic Flows?\n\n"
	tpl += "%s\n\n"
	tpl += Subtle("j/k, up/down: select") + Dot +
		Subtle("enter: choose") + Dot +
		Subtle("q, esc: quit")

	flows := ""
	for i, f := range m.FlowList.Flows {
		flows += fmt.Sprintf("%s\n", Checkbox(f.String(), c == i))
	}

	return fmt.Sprintf(tpl, flows)
}
