package flowlist

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	periodic_flow "jlowell000.github.io/budgeting/internal/model/periodicflow"
	"jlowell000.github.io/budgeting/internal/views/main"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type FlowListModel struct {
	Flows    []periodic_flow.PeriodicFlow // list of flows
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool
}

type Model interface {
	tea.Model
	GetMain() *main.MainModel
	GetFlowList() *FlowListModel
}

func FlowListUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	flowList := m.GetFlowList()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			flowList.Choice++
			if flowList.Choice > len(flowList.Flows)-1 {
				flowList.Choice = len(flowList.Flows) - 1
			}
		case "k", "up":
			flowList.Choice--
			if flowList.Choice < 0 {
				flowList.Choice = 0
			}

		case "b", "backspace":
			main.Chosen = false
		case "enter":
			flowList.Chosen = true
			// return m, nil
		}
	}

	return m, nil
}

func FlowListView(m Model) string {
	flowList := m.GetFlowList()
	c := flowList.Choice
	// The header
	tpl := "Viewing Periodic Flows?\n\n"
	tpl += "%s\n\n"
	tpl += util.Subtle("j/k, up/down: select") + util.Dot +
		util.Subtle("enter: choose") + util.Dot +
		util.Subtle("q, esc: quit")

	flows := ""
	for i, f := range flowList.Flows {
		flows += fmt.Sprintf("%s\n", util.Checkbox(f.String(), c == i))
	}

	return fmt.Sprintf(tpl, flows)
}
