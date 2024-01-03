package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"

	"jlowell000.github.io/budgeting/internal/views/accountlist"
	"jlowell000.github.io/budgeting/internal/views/flowlist"
	"jlowell000.github.io/budgeting/internal/views/main"
)

type AppModel struct {
	Main        main.MainModel
	FlowList    flowlist.FlowListModel
	AccountList accountlist.AccountListModel
	Quitting    bool
}

func (m AppModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if m.Main.Chosen == true {
		if m.Main.Choice == 1 {
			return flowlist.FlowListUpdate(msg, &m)
		} else if m.Main.Choice == 2 {
			return accountlist.AccountListUpdate(msg, &m)
		}
	}
	return main.MainUpdate(msg, &m)
}

func (m AppModel) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	if m.Main.Chosen == true {
		if m.Main.Choice == 1 {
			s = flowlist.FlowListView(&m)
		} else if m.Main.Choice == 2 {
			s = accountlist.AccountListView(&m)
		}
	} else {
		s = main.MainView(&m)
	}
	return indent.String("\n"+s+"\n\n", 2)
}

func (m *AppModel) GetMain() *main.MainModel {
	return &m.Main
}

func (m *AppModel) GetFlowList() *flowlist.FlowListModel {
	return &m.FlowList
}

func (m *AppModel) GetAccountList() *accountlist.AccountListModel {
	return &m.AccountList
}
