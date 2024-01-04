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

	"jlowell000.github.io/budgeting/internal/views"
	"jlowell000.github.io/budgeting/internal/views/accountlist"
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
	flows = []*periodicflow.PeriodicFlow{
		periodicflow.New(uuid.New(), "1", decimal.NewFromFloat(666.66), period.Weekly, time.Now()),
		periodicflow.New(uuid.New(), "2", decimal.NewFromFloat(123.66), period.Weekly, time.Now()),
		periodicflow.New(uuid.New(), "3", decimal.NewFromFloat(542.66), period.Weekly, time.Now()),
		periodicflow.New(uuid.New(), "4", decimal.NewFromFloat(1366.66), period.Weekly, time.Now()),
	}
	accounts = createTestAccounts()
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() views.AppModel {
	return views.AppModel{
		Main: mainview.MainModel{
			Choice:   1,
			Selected: make(map[int]struct{}),
		},
		FlowList: flowlist.FlowListModel{
			Flows:           flows,
			Selected:        make(map[int]struct{}),
			GetFlowListFunc: getTestFlows,
			CreateFlowFunc:  createTestFlow,
		},
		AccountList: accountlist.AccountListModel{
			Accounts:           accounts,
			Selected:           make(map[int]struct{}),
			GetAccountListFunc: getTestAccounts,
		},
	}
}

//TODO: below is test data to be removed in later issues

func getTestFlows() []*periodicflow.PeriodicFlow {
	return flows
}

func createTestFlow(
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	f := periodicflow.New(uuid.New(), name, amount, period, time.Now())
	flows = append(flows, f)
	return f
}

func getTestAccounts() []account.Account {
	return accounts
}

func createTestAccounts() []account.Account {
	amount := decimal.NewFromFloat(666.66)
	testSize := 10
	testSizeSlice := make([]int, testSize)
	var accounts []account.Account
	for i := range testSizeSlice {
		testSizeSlice[i] = i
		accounts = append(accounts, createAccount(amount, false))
	}
	return accounts
}

func createAccount(amount decimal.Decimal, excludable bool) account.Account {
	return account.Account{
		Id:         uuid.New(),
		Excludable: excludable,
		Book: []bookentry.BookEntry{
			{
				Id:        uuid.New(),
				Amount:    amount,
				Timestamp: time.Now(),
			},
		},
		UpdatedTimestamp: time.Now(),
	}
}
