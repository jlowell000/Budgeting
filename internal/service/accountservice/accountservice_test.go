package accountservice

import (
	"slices"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
	"jlowell000.github.io/budgeting/internal/model/data"
)

var (
	getDataCount  int = 0
	saveDataCount int = 0
	testTime          = time.UnixMilli(0)
	testTime2         = time.UnixMilli(1000000)
	testId            = uuid.New()
	testData      *data.DataModel

	subject = AccountService{
		Dataservice: &mockDataService{},
		GetTime:     func() time.Time { return testTime2 },
		GetId:       func() uuid.UUID { return testId },
	}
)

func Test_Create(t *testing.T) {
	resetData()
	subject.Delete(testId)
	expected := account.New(testId, "testCreate", true, testTime2)
	expectedList := append(slices.Clone(testData.Accounts), expected)
	slices.SortFunc(expectedList, compareAccountId)
	actual := subject.Create("testCreate", true)
	actualList := testData.Accounts

	assert.Equal(t, 3, getDataCount, "getDataCount")
	assert.Equal(t, 2, saveDataCount, "saveDataCount")
	assert.Equal(t, *expected, *actual, "not equal object")
	assert.Equal(t, findAccount(testId, expectedList), findAccount(testId, actualList), "same index")
}

func Test_GetAll(t *testing.T) {
	resetData()
	expected := slices.Clone(testData.Accounts)
	slices.SortFunc(expected, compareAccountId)
	actual := subject.GetAll()

	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "not equal object")
}
func Test_GetAllSortedByDate(t *testing.T) {
	resetData()
	expected := slices.Clone(testData.Accounts)
	slices.SortFunc(expected, compareAccountTime)
	actual := subject.GetAllSortedByDate()

	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "not equal object")
}

func Test_Update(t *testing.T) {
	resetData()
	expected := account.New(testId, "testEdited", true, testTime2)
	actual := subject.Update(testId, "testEdited", true)

	assert.Equal(t, 3, getDataCount, "getDataCount")
	assert.Equal(t, 1, saveDataCount, "saveDataCount")
	assert.Equal(t, *expected, *actual, "not equal object")
}

func Test_AddBookEntry(t *testing.T) {
	resetData()
	amount := decimal.NewFromInt(123)
	expected := account.New(testId, "test1", false, testTime2)
	expected.Book = []*bookentry.BookEntry{
		bookentry.New(testId, amount, testTime2),
	}
	actual := subject.AddBookEntry(testId, amount)

	assert.Equal(t, 3, getDataCount, "getDataCount")
	assert.Equal(t, 1, saveDataCount, "saveDataCount")
	assert.Equal(t, *expected, *actual, "not equal object")
}

func Test_Delete(t *testing.T) {
	resetData()
	expected := slices.DeleteFunc(
		slices.Clone(testData.Accounts),
		func(f *account.Account) bool { return f.Id == testId },
	)
	slices.SortFunc(expected, compareAccountId)
	subject.Delete(testId)

	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 1, saveDataCount, "saveDataCount")
	assert.True(t, !slices.ContainsFunc(testData.Accounts, func(f *account.Account) bool { return f.Id == testId }), "no longer contains")
}

type mockDataService struct{}

func resetData() {
	getDataCount = 0
	saveDataCount = 0
	testData = &data.DataModel{
		Accounts: []*account.Account{
			account.New(uuid.New(), "testA", false, time.UnixMilli(5000)),
			account.New(uuid.New(), "testB", false, time.UnixMilli(4000)),
			account.New(testId, "test1", false, testTime),
			account.New(uuid.New(), "testC", false, time.UnixMilli(10000)),
			account.New(uuid.New(), "testD", false, time.UnixMilli(1000)),
		},
	}

}
func (m *mockDataService) GetData() *data.DataModel {
	getDataCount++
	return testData
}
func (m *mockDataService) SaveData(newD *data.DataModel) *data.DataModel {
	saveDataCount++
	testData = newD
	return testData
}
