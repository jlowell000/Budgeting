package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
	"jlowell000.github.io/budgeting/internal/service/dataservice"

	"jlowell000.github.io/budgeting/internal/views"
	"jlowell000.github.io/budgeting/internal/views/accountlist"
	"jlowell000.github.io/budgeting/internal/views/accountview"
	"jlowell000.github.io/budgeting/internal/views/flowlist"
	"jlowell000.github.io/budgeting/internal/views/mainview"
)

const (
	CMD_CREATE         = "create"
	CMD_READ           = "read"
	CMD_QUIT           = "quit"
	FLG_ALL            = "all"
	VAR_CONTENT        = "content"
	ENTRYLIST_FILENAME = "./data.json"
)

var (
	data *dataservice.DataModel
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() views.AppModel {
	data = dataservice.GetDataFromFile(ENTRYLIST_FILENAME)

	return views.AppModel{
		Main: mainview.MainModel{
			Choice:   1,
			Selected: make(map[int]struct{}),
		},
		FlowList: flowlist.FlowListModel{
			Flows:           data.Flows,
			Selected:        make(map[int]struct{}),
			CreateFlowFunc:  createFlow,
			GetFlowListFunc: getTestFlows,
			UpdateFlowFunc:  updateTestFlow,
		},
		AccountList: accountlist.AccountListModel{
			Accounts:           data.Accounts,
			Selected:           make(map[int]struct{}),
			CreateAccountFunc:  createAccount,
			GetAccountListFunc: getAccounts,
			UpdateAccountFunc:  updateAccount,
		},
		Account: accountview.AccountModel{
			AddEntry: addBookEntry,
		},
		SavaDataFunc: func() { dataservice.SaveDataToFile(data, ENTRYLIST_FILENAME) },
	}
}

//TODO: below is test data to be removed in later issues

func getTestFlows() []*periodicflow.PeriodicFlow {
	return data.Flows
}

func createFlow(
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	f := periodicflow.New(uuid.New(), name, amount, period, time.Now())
	data.Flows = append(data.Flows, f)
	return f
}

func updateTestFlow(
	id uuid.UUID,
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	for i, f := range data.Flows {
		if f.Id == id {
			data.Flows[i] = f.Update(
				name,
				amount,
				period,
				time.Now(),
			)
			return data.Flows[i]
		}
	}
	return nil
}

func getAccounts() []*account.Account {
	return data.Accounts
}

func createAccount(name string, excludable bool) *account.Account {
	a := account.New(
		uuid.New(),
		name,
		excludable,
		time.Now(),
	)
	data.Accounts = append(data.Accounts, a)
	return a
}

func updateAccount(
	id uuid.UUID,
	name string,
	excludable bool,
) *account.Account {
	for i, f := range data.Accounts {
		if f.Id == id {
			data.Accounts[i] = f.Update(
				name,
				excludable,
				time.Now(),
			)
			return f
		}
	}
	return nil
}

func addBookEntry(a *account.Account, amount decimal.Decimal) *account.Account {
	a.Book = append(
		a.Book,
		bookentry.New(
			uuid.New(),
			amount,
			time.Now(),
		),
	)
	return a
}
