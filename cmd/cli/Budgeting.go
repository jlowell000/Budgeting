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
	mainview "jlowell000.github.io/budgeting/internal/views/main"
)

const (
	CMD_CREATE         = "create"
	CMD_READ           = "read"
	CMD_QUIT           = "quit"
	FLG_ALL            = "all"
	VAR_CONTENT        = "content"
	ENTRYLIST_FILENAME = "./data.json"
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
			Flows:    createTestFlows(),
			Selected: make(map[int]struct{}),
		},
		AccountList: accountlist.AccountListModel{
			Accounts: createTestAccounts(),
			Selected: make(map[int]struct{}),
		},
	}
}

//TODO: below is test data to be removed in later issues

func createTestFlows() []periodicflow.PeriodicFlow {
	return []periodicflow.PeriodicFlow{
		*periodicflow.New(uuid.New(), decimal.NewFromFloat(666.66), period.Weekly, time.Now()),
		*periodicflow.New(uuid.New(), decimal.NewFromFloat(123.66), period.Weekly, time.Now()),
		*periodicflow.New(uuid.New(), decimal.NewFromFloat(542.66), period.Weekly, time.Now()),
		*periodicflow.New(uuid.New(), decimal.NewFromFloat(1366.66), period.Weekly, time.Now()),
	}
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
