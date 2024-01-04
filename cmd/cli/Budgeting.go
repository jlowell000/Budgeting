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
			CreateFlowFunc:  createTestFlow,
			GetFlowListFunc: getTestFlows,
			UpdateFlowFunc:  updateTestFlow,
		},
		AccountList: accountlist.AccountListModel{
			Accounts:           accounts,
			Selected:           make(map[int]struct{}),
			CreateAccountFunc:  createAccountFunc,
			GetAccountListFunc: getTestAccounts,
			UpdateAccountFunc:  updateAccount,
		},
		Account: accountview.AccountModel{
			AddEntry: addBookEntry,
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

func updateTestFlow(
	id uuid.UUID,
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	for i, f := range flows {
		if f.Id == id {
			flows[i] = f.Update(
				name,
				amount,
				period,
			)
			return flows[i]
		}
	}
	return nil
}

func getTestAccounts() []*account.Account {
	return accounts
}

func createTestAccounts() []*account.Account {
	testSize := 10
	testSizeSlice := make([]int, testSize)
	var accounts []*account.Account
	for i := range testSizeSlice {
		testSizeSlice[i] = i
		accounts = append(accounts, createAccount("acc"+fmt.Sprint(i), false))
	}
	return accounts
}

func createAccount(name string, excludable bool) *account.Account {
	amount := decimal.NewFromFloat(111.11)
	testSize := 10
	testSizeSlice := make([]int, testSize)
	var entries []bookentry.BookEntry
	for i := range testSizeSlice {
		testSizeSlice[i] = i
		entries = append(
			entries,
			bookentry.BookEntry{
				Id:        uuid.New(),
				Amount:    amount.Mul(decimal.NewFromInt(int64(i))),
				Timestamp: time.Now(),
			},
		)
	}

	return &account.Account{
		Id:               uuid.New(),
		Name:             name,
		Excludable:       excludable,
		Book:             entries,
		UpdatedTimestamp: time.Now(),
	}
}

func createAccountFunc(name string, excludable bool) account.Account {
	a := &account.Account{
		Id:               uuid.New(),
		Name:             name,
		Excludable:       excludable,
		UpdatedTimestamp: time.Now(),
	}
	accounts = append(accounts, a)
	return *a
}

func updateAccount(
	id uuid.UUID,
	name string,
	excludable bool,
) *account.Account {
	for i, f := range accounts {
		if f.Id == id {
			f.Name = name
			f.Excludable = excludable
			accounts[i] = f
			return f
		}
	}
	return nil
}

func addBookEntry(a *account.Account, amount decimal.Decimal) *account.Account {
	a.Book = append(
		a.Book,
		bookentry.BookEntry{
			Id:        uuid.New(),
			Amount:    amount,
			Timestamp: time.Now(),
		},
	)
	return a
}
