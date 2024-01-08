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
	"jlowell000.github.io/budgeting/internal/model/data"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
)

const (
	TEST_FILE_NAME = "test.json"
)

var (
	getDataJSONCount int = 0
	putDataJSONCount int = 0
	testTime             = time.Now()
	testId               = uuid.New()
	putDataJSONValue string
	subject          DataService
)

func Test_GetData(t *testing.T) {
	expected := testData("get")
	testData := testData("get")
	setUpMocking(t, testData)
	actual := subject.GetData()

	assert.Equal(t, expected.ToJSON(), actual.ToJSON())
	assert.Equal(t, 1, getDataJSONCount, "getDataJSON")
	assert.Equal(t, 0, putDataJSONCount, "putDataJSONCount")
}

func Test_SaveData(t *testing.T) {
	expected := testData("save")
	setUpMocking(t, expected)
	testDataModel := testData("save")
	subject.SaveData(testDataModel)

	assert.Equal(t, 0, getDataJSONCount, "getDataJSONCount")
	assert.Equal(t, 1, putDataJSONCount, "putDataJSONCount")
	assert.Equal(t, makaDataJSON(expected), putDataJSONValue, "json not equal")
}

func testData(mod string) *data.DataModel {
	return &data.DataModel{
		Flows: []*periodicflow.PeriodicFlow{
			periodicflow.New(testId, "1"+mod, decimal.NewFromFloat(666.66), period.Weekly, testTime),
			periodicflow.New(testId, "2"+mod, decimal.NewFromFloat(123.66), period.Weekly, testTime),
			periodicflow.New(testId, "3"+mod, decimal.NewFromFloat(542.66), period.Weekly, testTime),
			periodicflow.New(testId, "4"+mod, decimal.NewFromFloat(1366.66), period.Weekly, testTime),
		},
		Accounts: createTestAccounts(),
	}
}

func setUpMocking(t *testing.T, data *data.DataModel) {
	subject = DataService{
		Filename: TEST_FILE_NAME,
		GetDataJSON: func(fileName string) []byte {
			getDataJSONCount++
			return data.ToJSON()
		},
		PutDataJSON: func(data []byte, fileName string) {
			putDataJSONCount++
			putDataJSONValue = string(data)
			return
		},
	}
	getDataJSONCount = 0
	putDataJSONCount = 0
	putDataJSONValue = ""
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
				testId,
				"flow_"+fmt.Sprint(i),
				decimal.NewFromFloat(111.11).Mul(decimal.NewFromInt(int64(i))),
				period.Weekly,
				testTime,
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
				Id:        testId,
				Amount:    amount.Mul(decimal.NewFromInt(int64(i))),
				Timestamp: testTime,
			},
		)
	}

	return &account.Account{
		Id:               testId,
		Name:             name,
		Excludable:       excludable,
		Book:             entries,
		UpdatedTimestamp: testTime,
	}
}

func makaDataJSON(
	d *data.DataModel,
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
