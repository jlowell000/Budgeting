package accountlist

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/views/main"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type AccountListModel struct {
	Accounts []account.Account // list of flows
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool
}

type Model interface {
	tea.Model
	GetMain() *main.MainModel
	GetAccountList() *AccountListModel
}

func AccountListUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	accountList := m.GetAccountList()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			accountList.Choice++
			if accountList.Choice > len(accountList.Accounts)-1 {
				accountList.Choice = len(accountList.Accounts) - 1
			}
		case "k", "up":
			accountList.Choice--
			if accountList.Choice < 0 {
				accountList.Choice = 0
			}

		case "b", "backspace":
			main.Chosen = false
		case "enter":
			accountList.Chosen = true
			// return m, nil
		}
	}

	return m, nil
}

func AccountListView(m Model) string {
	accountList := m.GetAccountList()
	c := accountList.Choice
	// The header
	tpl := "Viewing Accounts\n\n"
	tpl += "%s\n\n"
	tpl += util.Subtle("j/k, up/down: select") + util.Dot +
		util.Subtle("enter: choose") + util.Dot +
		util.Subtle("q, esc: quit")

	accounts := ""
	for i, f := range accountList.Accounts {
		accounts += fmt.Sprintf("%s\n", util.Checkbox(f.String(), c == i))
	}

	return fmt.Sprintf(tpl, accounts)
}
