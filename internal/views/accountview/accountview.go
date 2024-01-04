package accountview

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
	"jlowell000.github.io/budgeting/internal/views/accountlist"
	"jlowell000.github.io/budgeting/internal/views/mainview"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type AccountModel struct {
	// Account account.Account
}

type Model interface {
	tea.Model
	GetMain() *mainview.MainModel
	GetAccountList() *accountlist.AccountListModel
	// GetForm() *form.FormModel
}

func AccountUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	accountList := m.GetAccountList()
	// accountList.Accounts = accountList.GetAccountListFunc()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// case "j", "down":
		// 	accountList.Choice++
		// 	if accountList.Choice > len(accountList.Accounts)-1 {
		// 		accountList.Choice = len(accountList.Accounts) - 1
		// 	}
		// case "k", "up":
		// 	accountList.Choice--
		// 	if accountList.Choice < 0 {
		// 		accountList.Choice = 0
		// 	}

		case "b":
			main.Choice = 2
		case "enter":
			accountList.Chosen = false
			main.Choice = 2
		}
	}

	return m, nil
}

func AccountView(m Model) string {
	accountList := m.GetAccountList()
	c := accountList.Choice
	account := accountList.Accounts[c]
	// The header
	tpl := "Viewing Accounts\n\n"
	tpl += fmt.Sprintf("%s\n\n", accountlist.DisplayString(account))
	tpl += "%s\n\n"
	tpl += util.Instructions()

	bookEntries := ""
	for _, e := range account.Book {
		bookEntries += fmt.Sprintf("%s\n", entryDisplay(e))
	}

	return fmt.Sprintf(tpl, bookEntries)
}

func entryDisplay(e bookentry.BookEntry) string {
	return "Timestamp: " + e.Timestamp.String() + "; " +
		"Amount: " + e.Amount.String()
}
