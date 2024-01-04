package accountlist

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/views/form"
	"jlowell000.github.io/budgeting/internal/views/mainview"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type AccountListModel struct {
	Accounts []account.Account
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool

	/* Tell the model how to Create a flows */
	CreateAccountFunc func(string, decimal.Decimal, period.Period) account.Account
	/* Tell the model how to get list of accounts */
	GetAccountListFunc func() []account.Account
	/* Update FlowList */
	UpdateAccountFunc func(uuid.UUID, string, decimal.Decimal, period.Period) account.Account
}

type Model interface {
	tea.Model
	GetMain() *mainview.MainModel
	GetAccountList() *AccountListModel
	GetForm() *form.FormModel
}

func AccountListUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	accountList := m.GetAccountList()
	accountList.Accounts = accountList.GetAccountListFunc()
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

		case "b":
			main.Chosen = false
		case "enter":
			accountList.Chosen = true
			main.Choice = 4
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
	tpl += util.Instructions()

	accounts := ""
	for i, f := range accountList.Accounts {
		accounts += fmt.Sprintf(
			"%s\n",
			util.Checkbox(DisplayString(f), c == i),
		)
	}

	return fmt.Sprintf(tpl, accounts)
}

func DisplayString(a account.Account) string {
	return "Name: " + a.Name + "; " +
		"Amount: " + a.GetLatestBookEntry().Amount.String() + ";"
}
