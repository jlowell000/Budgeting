package dataservice

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

const (
	TEST_FILE_NAME = "test.json"
	TEST_ID        = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_TIME      = "2006-01-23T15:04:05Z"
)

var (
	getDataJSONCount int = 0
	putDataJSONCount int = 0
	putDataJSONValue string
	getNewIdCount    int = 0
	getTimeCount     int = 0
)

func TestGetDataFromFile(t *testing.T) {
	expected := testData()
	setUpMocking(t, expected)
	actual := GetDataFromFile(TEST_FILE_NAME)

	assert.Equal(t, expected.ToJSON(), actual.ToJSON())
	assert.Equal(t, 1, getDataJSONCount, "getDataJSON")
	assert.Equal(t, 0, putDataJSONCount, "putDataJSONCount")
	assert.Equal(t, 0, getNewIdCount, "getNewIdCount")
	assert.Equal(t, 0, getTimeCount, "getTimeCount")

}

func TestSaveEntryListToFile(t *testing.T) {
	expected := testData()
	setUpMocking(t, expected)
	SaveDataToFile(expected, TEST_FILE_NAME)

	assert.Equal(t, 0, getDataJSONCount, "getDataJSONCount")
	assert.Equal(t, 1, putDataJSONCount, "putDataJSONCount")
	assert.Equal(t, makaDataJSON(expected), putDataJSONValue, "putDataJSONCount")
	assert.Equal(t, 0, getNewIdCount, "getNewIdCount")
	assert.Equal(t, 0, getTimeCount, "getTimeCount")
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

func setUpMocking(t *testing.T, data *DataModel) {
	getDataJSON = func(fileName string) []byte {
		getDataJSONCount++
		return data.ToJSON()
	}
	putDataJSON = func(data []byte, fileName string) {
		putDataJSONCount++
		putDataJSONValue = string(data)
		return
	}

	getNewId = func() uuid.UUID {
		getNewIdCount++
		id, err1 := uuid.Parse(TEST_ID)
		if err1 != nil {
			log.Fatal(err1)
		}
		return id
	}
	getTime = func() time.Time {
		getTimeCount++
		timestamp, err2 := time.Parse(time.RFC3339, TEST_TIME)
		if err2 != nil {
			log.Fatal(err2)
		}
		return timestamp
	}

	getDataJSONCount = 0
	putDataJSONCount = 0
	putDataJSONValue = ""
	getNewIdCount = 0
	getTimeCount = 0
}

func createTestFlow() []*periodicflow.PeriodicFlow {
	testSize := 1
	testSizeSlice := make([]int, testSize)
	var flows []*periodicflow.PeriodicFlow
	for i := range testSizeSlice {
		testSizeSlice[i] = i
		flows = append(
			flows,
			periodicflow.New(
				uuid.New(),
				"flow_"+fmt.Sprint(i),
				decimal.NewFromFloat(111.11).Mul(decimal.NewFromInt(int64(i))),
				period.Weekly,
				time.Now(),
			),
		)
	}
	return flows
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
		"\"weekly_amount\":\"" + p.WeeklyAmount.String() + "\"," +
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
