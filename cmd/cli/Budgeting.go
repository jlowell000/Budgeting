package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/io"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
	"jlowell000.github.io/budgeting/internal/model/data"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
	"jlowell000.github.io/budgeting/internal/service"
	"jlowell000.github.io/budgeting/internal/service/dataservice"
	"jlowell000.github.io/budgeting/internal/service/periodicflowservice"

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
	d  *data.DataModel
	ds service.DataServiceInterface = &dataservice.DataService{
		Filename:    ENTRYLIST_FILENAME,
		GetDataJSON: io.ReadFromFile,
		PutDataJSON: io.WriteToFile,
	}
	flowService service.PeriodicFlowServiceInterface = &periodicflowservice.PeriodicFlowService{
		Dataservice: ds,
		GetTime:     time.Now,
	}
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() views.AppModel {
	d = ds.GetData()

	return views.AppModel{
		Main: mainview.MainModel{
			Choice:   1,
			Selected: make(map[int]struct{}),
		},
		FlowList: flowlist.FlowListModel{
			Flows:           d.Flows,
			Selected:        make(map[int]struct{}),
			CreateFlowFunc:  createFlow,
			GetFlowListFunc: getTestFlows,
			UpdateFlowFunc:  flowService.UpdatePeriodicFlow,
		},
		AccountList: accountlist.AccountListModel{
			Accounts:           d.Accounts,
			Selected:           make(map[int]struct{}),
			CreateAccountFunc:  createAccount,
			GetAccountListFunc: getAccounts,
			UpdateAccountFunc:  updateAccount,
		},
		Account: accountview.AccountModel{
			AddEntry: addBookEntry,
		},
		SavaDataFunc: func() { d = ds.SaveData(d) },
	}
}

//TODO: below is test data to be removed in later issues

func getTestFlows() []*periodicflow.PeriodicFlow {
	return d.Flows
}

func createFlow(
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	f := periodicflow.New(uuid.New(), name, amount, period, time.Now())
	d.Flows = append(d.Flows, f)
	return f
}

func getAccounts() []*account.Account {
	return d.Accounts
}

func createAccount(name string, excludable bool) *account.Account {
	a := account.New(
		uuid.New(),
		name,
		excludable,
		time.Now(),
	)
	d.Accounts = append(d.Accounts, a)
	return a
}

func updateAccount(
	id uuid.UUID,
	name string,
	excludable bool,
) *account.Account {
	for i, f := range d.Accounts {
		if f.Id == id {
			d.Accounts[i] = f.Update(
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
