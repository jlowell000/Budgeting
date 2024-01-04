package accountlist

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/views/form"
	"jlowell000.github.io/budgeting/internal/views/mainview"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type AccountListModel struct {
	Accounts []*account.Account
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool

	/* Tell the model how to Create a flows */
	CreateAccountFunc func(string, bool) account.Account
	/* Tell the model how to get list of accounts */
	GetAccountListFunc func() []*account.Account
	/* Update FlowList */
	UpdateAccountFunc func(uuid.UUID, string, bool) *account.Account
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
	form := m.GetForm()
	checkFormForNewData(accountList, form)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			accountList.Choice++
			if accountList.Choice > len(accountList.Accounts)-1 {
				accountList.Choice = len(accountList.Accounts) - 1
			}
		case "up":
			accountList.Choice--
			if accountList.Choice < 0 {
				accountList.Choice = 0
			}

		case "e":
			accountList.Chosen = true
			c := accountList.Choice
			form.LastScreen = 2
			form.Inputs = createFormInputs(
				accountList.Accounts[c].Name,
				accountList.Accounts[c].Excludable,
			)
			main.Choice = 3

		case "n":
			form.LastScreen = 2
			form.Inputs = createFormInputs("", false)
			main.Choice = 3

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

func DisplayString(a *account.Account) string {
	str := "Name: " + a.Name + util.Dot +
		"Amount: " + a.GetLatestBookEntry().Amount.String() + util.Dot +
		"Updated: " + util.TimeFormat(a.UpdatedTimestamp)

	if a.Excludable {
		str += util.Dot + "Excluded"
	}
	return str
}

func createFormInputs(
	name string,
	excludeable bool,
) []textinput.Model {
	inputs := make([]textinput.Model, 2)
	var t textinput.Model

	var excludeString string
	if excludeable {
		excludeString = "Y"
	} else {
		excludeString = "N"
	}

	for i := range inputs {
		t = textinput.New()
		t.Cursor.Style = util.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.SetValue(name)
		case 1:
			t.Placeholder = "Exclude"
			t.SetValue(excludeString)
			// TODO valideate as Y/N
		}

		inputs[i] = t
	}

	return inputs
}

func checkFormForNewData(
	accountList *AccountListModel,
	form *form.FormModel,
) bool {
	if form.Submitted {
		if !accountList.Chosen {
			accountList.CreateAccountFunc(
				form.Inputs[0].Value(),
				form.Inputs[1].Value() == "Y",
			)
		} else {
			accountList.UpdateAccountFunc(
				accountList.Accounts[accountList.Choice].Id,
				form.Inputs[0].Value(),
				form.Inputs[1].Value() == "Y",
			)
		}
		accountList.Chosen = false
		form.ResetForm()
		return true
	}
	accountList.Chosen = false
	form.ResetForm()
	return false
}
