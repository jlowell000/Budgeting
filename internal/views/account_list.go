package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"jlowell000.github.io/budgeting/internal/model/account"
)

type AccountListModel struct {
	Accounts []account.Account // list of flows
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool
}

func accountListUpdate(msg tea.Msg, m AppModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.AccountList.Choice++
			if m.AccountList.Choice > len(m.AccountList.Accounts)-1 {
				m.AccountList.Choice = len(m.AccountList.Accounts) - 1
			}
		case "k", "up":
			m.AccountList.Choice--
			if m.AccountList.Choice < 0 {
				m.AccountList.Choice = 0
			}

		case "b", "backspace":
			m.Main.Chosen = false
		case "enter":
			m.AccountList.Chosen = true
			// return m, nil
		}
	}

	return m, nil
}

func accountListView(m AppModel) string {
	c := m.AccountList.Choice
	// The header
	tpl := "Viewing Accounts\n\n"
	tpl += "%s\n\n"
	tpl += Subtle("j/k, up/down: select") + Dot +
		Subtle("enter: choose") + Dot +
		Subtle("q, esc: quit")

	accounts := ""
	for i, f := range m.AccountList.Accounts {
		accounts += fmt.Sprintf("%s\n", Checkbox(f.String(), c == i))
	}

	return fmt.Sprintf(tpl, accounts)
}
