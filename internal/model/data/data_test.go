package data

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
)

func Test_FromJSON(t *testing.T) {
	expected := testData()
	actual := FromJSON([]byte(makaDataJSON(expected)))

	assert.Equal(t, expected.ToJSON(), actual.ToJSON())
}

func Test_ToJson(t *testing.T) {
	expected := testData()
	actual := string(expected.ToJSON())
	assert.Equal(t, makaDataJSON(expected), actual, "ToJSON")
}

func testData() *DataModel {
	return &DataModel{
		Flows: []*periodicflow.PeriodicFlow{
			periodicflow.New(uuid.New(), "1", decimal.NewFromFloat(666.66), period.Weekly, time.Now()),
			periodicflow.New(uuid.New(), "2", decimal.NewFromFloat(123.66), period.Weekly, time.Now()),
			periodicflow.New(uuid.New(), "3", decimal.NewFromFloat(542.66), period.Weekly, time.Now()),
			periodicflow.New(uuid.New(), "4", decimal.NewFromFloat(1366.66), period.Weekly, time.Now()),
		},
		Accounts: createTestAccounts(),
	}
}

func createTestAccounts() []*account.Account {
	testSize := 1
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
	testSize := 1
	testSizeSlice := make([]int, testSize)
	var entries []*bookentry.BookEntry
	for i := range testSizeSlice {
		testSizeSlice[i] = i
		entries = append(
			entries,
			&bookentry.BookEntry{
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

func makaDataJSON(
	d *DataModel,
) string {
	var flowsPlaceHolder string
	for i, f := range d.Flows {
		flowsPlaceHolder += makeFlowJSON(f)
		if i+1 < len(d.Flows) {
			flowsPlaceHolder += ","
		}
	}

	var accountsPlaceHolder string
	for i, a := range d.Accounts {
		accountsPlaceHolder += makeAccountJSON(a)
		if i+1 < len(d.Accounts) {
			accountsPlaceHolder += ","
		}
	}
	return fmt.Sprintf(
		"{\"flows\":[%s],\"accounts\":[%s]}",
		flowsPlaceHolder,
		accountsPlaceHolder,
	)

}

func makeFlowJSON(
	p *periodicflow.PeriodicFlow,
) string {
	return "{" +
		"\"id\":\"" + p.Id.String() + "\"," +
		"\"name\":\"" + p.Name + "\"," +
		"\"amount\":\"" + p.Amount.String() + "\"," +
		"\"period\":\"" + p.Period.String() + "\"," +
		"\"monthly_amount\":\"" + p.MonthlyAmount.String() + "\"," +
		"\"updated_timestamp\":" + makeTimeStampJSON(p.UpdatedTimestamp) +
		"}"
}

func makeAccountJSON(
	a *account.Account,
) string {
	var ex string
	if a.Excludable {
		ex = "true"
	} else {
		ex = "false"
	}

	s := "{" +
		"\"id\":\"" + a.Id.String() + "\"," +
		"\"name\":\"" + a.Name + "\"," +
		"\"excludable\":" + ex + "," +
		"\"book\":[%s]," +
		"\"updated_timestamp\":" + makeTimeStampJSON(a.UpdatedTimestamp) +
		"}"

	var entries string
	for i, b := range a.Book {
		entries += makeBookEntryJSON(b)
		if i+1 < len(a.Book) {
			entries += ","
		}
	}
	return fmt.Sprintf(s, entries)
}

func makeBookEntryJSON(
	e *bookentry.BookEntry,
) string {
	return "{" +
		"\"id\":\"" + e.Id.String() + "\"," +
		"\"amount\":\"" + e.Amount.String() + "\"," +
		"\"timestamp\":" + makeTimeStampJSON(e.Timestamp) +
		"}"
}

func makeTimeStampJSON(t time.Time) string {
	dataJSON, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	return string(dataJSON)
}
