package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"

	"jlowell000.github.io/budgeting/internal/views/accountlist"
	"jlowell000.github.io/budgeting/internal/views/accountview"
	"jlowell000.github.io/budgeting/internal/views/flowlist"
	"jlowell000.github.io/budgeting/internal/views/form"
	"jlowell000.github.io/budgeting/internal/views/mainview"
)

type AppModel struct {
	Main        mainview.MainModel
	FlowList    flowlist.FlowListModel
	FlowForm    form.FormModel
	AccountList accountlist.AccountListModel
	Account     accountview.AccountModel
	Quitting    bool

	SavaDataFunc func()
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

	if m.Main.Chosen == true {
		switch m.Main.Choice {
		case 1:
			return flowlist.FlowListUpdate(msg, &m)
		case 2:
			return accountlist.AccountListUpdate(msg, &m)
		case 3:
			return form.FormUpdate(msg, &m)
		case 4:
			return accountview.AccountUpdate(msg, &m)
		}
	}
	return mainview.MainUpdate(msg, &m)
}

func (m AppModel) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	if m.Main.Chosen == true {
		switch m.Main.Choice {
		case 1:
			s = flowlist.FlowListView(&m)
		case 2:
			s = accountlist.AccountListView(&m)
		case 3:
			s = form.FormView(&m)
		case 4:
			s = accountview.AccountView(&m)
		}
	} else {
		s = mainview.MainView(&m)
	}
	return indent.String("\n"+s+"\n\n", 2)
}

func (m *AppModel) UpdateInputs(msg tea.Msg) tea.Cmd {
	if m.Main.Chosen == true {
		if m.Main.Choice == 3 {
			cmds := make([]tea.Cmd, len(m.FlowForm.Inputs))

			for i := range m.FlowForm.Inputs {
				m.FlowForm.Inputs[i], cmds[i] = m.FlowForm.Inputs[i].Update(msg)
			}

			return tea.Batch(cmds...)
		}
	}
	return nil
}

func (m *AppModel) GetMain() *mainview.MainModel {
	return &m.Main
}

func (m *AppModel) GetFlowList() *flowlist.FlowListModel {
	return &m.FlowList
}

func (m *AppModel) GetForm() *form.FormModel {
	return &m.FlowForm
}

func (m *AppModel) GetAccountList() *accountlist.AccountListModel {
	return &m.AccountList
}

func (m *AppModel) GetAccountView() *accountview.AccountModel {
	return &m.Account
}

func (m *AppModel) SavaData() {
	m.SavaDataFunc()
}
