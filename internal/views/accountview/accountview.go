package accountview

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
	"jlowell000.github.io/budgeting/internal/service"
	"jlowell000.github.io/budgeting/internal/views/accountlist"
	"jlowell000.github.io/budgeting/internal/views/form"
	"jlowell000.github.io/budgeting/internal/views/mainview"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type AccountModel struct {
	AccountService service.AccountServiceInterface
}

type Model interface {
	tea.Model
	GetMain() *mainview.MainModel
	GetAccountView() *AccountModel
	GetAccountList() *accountlist.AccountListModel
	GetForm() *form.FormModel
}

func AccountUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	accountView := m.GetAccountView()
	accountList := m.GetAccountList()
	form := m.GetForm()
	checkFormForNewData(accountView, accountList, form)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "n":
			form.LastScreen = 4
			form.Inputs = createFormInputs()
			main.Choice = 3
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
	accountView := m.GetAccountView()
	accountList := m.GetAccountList()
	account := accountView.AccountService.Get(accountList.ChoiceId)
	// The header
	tpl := "Viewing Accounts\n\n"
	tpl += fmt.Sprintf("%s\n\n", accountlist.DisplayString(account))
	tpl += "\nId: " + account.Id.String() + "\n"
	tpl += "%s\n\n"
	tpl += util.Instructions()

	bookEntries := ""
	for _, e := range account.Book {
		bookEntries += fmt.Sprintf("%s\n", entryDisplay(e))
	}

	return fmt.Sprintf(tpl, bookEntries)
}

func entryDisplay(e *bookentry.BookEntry) string {
	str := ""
	if e != nil {
		str += "Timestamp: " + util.TimeFormat(e.Timestamp) + util.Dot +
			"Amount: " + e.Amount.String()
	}
	return str
}

func createFormInputs() []textinput.Model {
	inputs := make([]textinput.Model, 1)
	var t textinput.Model
	for i := range inputs {
		t = textinput.New()
		t.Cursor.Style = util.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Amount"
			t.Validate = util.IsMoneyNumber
		}

		inputs[i] = t
	}

	return inputs
}

func checkFormForNewData(
	account *AccountModel,
	accountList *accountlist.AccountListModel,
	form *form.FormModel,
) bool {
	if form.Submitted {
		d, _ := decimal.NewFromString(form.Inputs[0].Value())
		account.AccountService.AddBookEntry(accountList.ChoiceId, d)
		form.ResetForm()
		return true
	}
	return false
}
