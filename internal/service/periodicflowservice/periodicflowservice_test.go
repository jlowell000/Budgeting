package periodicflowservice

import (
	"slices"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"jlowell000.github.io/budgeting/internal/model/data"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
)

var (
	getDataCount  int = 0
	saveDataCount int = 0
	testTime          = time.UnixMilli(0)
	testTime2         = time.UnixMilli(1000000)
	testId            = uuid.New()
	testData      *data.DataModel

	subject = PeriodicFlowService{
		Dataservice: &mockDataService{},
		GetTime:     func() time.Time { return testTime2 },
		GetId:       func() uuid.UUID { return testId },
	}
)

func Test_Create(t *testing.T) {
	resetData()
	subject.Delete(testId)
	expected := periodicflow.New(testId, "testCreate", decimal.NewFromFloat(111.11), period.Monthly, testTime2)
	expectedList := append(slices.Clone(testData.Flows), expected)
	slices.SortFunc(expectedList, compareFlowId)
	actual := subject.Create("testCreate", decimal.NewFromFloat(111.11), period.Monthly)
	actualList := testData.Flows

	assert.Equal(t, 3, getDataCount, "getDataCount")
	assert.Equal(t, 2, saveDataCount, "saveDataCount")
	assert.Equal(t, *expected, *actual, "not equal object")
	assert.Equal(t, findPeriodicFlow(testId, expectedList), findPeriodicFlow(testId, actualList), "same index")
}

func Test_GetAll(t *testing.T) {
	resetData()
	expected := slices.Clone(testData.Flows)
	slices.SortFunc(expected, compareFlowId)
	actual := subject.GetAll()

	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "not equal object")
}
func Test_GetAllSortedByDate(t *testing.T) {
	resetData()
	expected := slices.Clone(testData.Flows)
	slices.SortFunc(expected, compareFlowTime)
	actual := subject.GetAllSortedByDate()

	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "not equal object")
}

func Test_GetTotalInflow(t *testing.T) {
	resetData()
	expected := decimal.NewFromFloat(1666.98)
	actual := subject.GetTotalWeeklyInflow()
	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "equality")
}
func Test_GetTotalOutflow(t *testing.T) {
	resetData()
	expected := decimal.NewFromFloat(-334.32)
	actual := subject.GetTotalWeeklyOutflow()
	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "equality")
}

func Test_GetTotalFlow(t *testing.T) {
	resetData()
	expected := decimal.NewFromFloat(1332.66)
	actual := subject.GetTotalWeeklyFlow()
	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "equality")
}

func Test_GetProjectedTotalInflow(t *testing.T) {
	resetData()
	a := decimal.NewFromInt(6)
	p := period.Monthly
	expected := decimal.NewFromFloat(1666.98).Mul(a).Mul(p.WeeklyAmount())
	actual := subject.GetProjectedTotalInflow(a, p)
	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "equality")
}
func Test_GetProjectedTotalWeeklyOutflow(t *testing.T) {
	resetData()
	a := decimal.NewFromInt(6)
	p := period.Monthly
	expected := decimal.NewFromFloat(-334.32).Mul(a).Mul(p.WeeklyAmount())
	actual := subject.GetProjectedTotalOutflow(a, p)
	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "equality")
}

func Test_GetProjectedTotalWeeklyFlow(t *testing.T) {
	resetData()
	a := decimal.NewFromInt(6)
	p := period.Monthly
	expected := decimal.NewFromFloat(1332.66).Mul(a).Mul(p.WeeklyAmount())
	actual := subject.GetProjectedTotalFlow(a, p)
	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 0, saveDataCount, "saveDataCount")
	assert.Equal(t, expected, actual, "equality")
}

func Test_Update(t *testing.T) {
	resetData()
	expected := periodicflow.New(testId, "testEdited", decimal.NewFromFloat(111.11), period.Monthly, testTime2)
	actual := subject.Update(testId, "testEdited", decimal.NewFromFloat(111.11), period.Monthly)

	assert.Equal(t, 3, getDataCount, "getDataCount")
	assert.Equal(t, 1, saveDataCount, "saveDataCount")
	assert.Equal(t, *expected, *actual, "not equal object")
}

func Test_Delete(t *testing.T) {
	resetData()
	expected := slices.DeleteFunc(
		slices.Clone(testData.Flows),
		func(f *periodicflow.PeriodicFlow) bool { return f.Id == testId },
	)
	slices.SortFunc(expected, compareFlowId)
	subject.Delete(testId)

	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 1, saveDataCount, "saveDataCount")
	assert.True(t, !slices.ContainsFunc(testData.Flows, func(f *periodicflow.PeriodicFlow) bool { return f.Id == testId }), "no longer contains")
}

type mockDataService struct{}

func resetData() {
	getDataCount = 0
	saveDataCount = 0
	testData = &data.DataModel{
		Flows: []*periodicflow.PeriodicFlow{
			periodicflow.New(uuid.New(), "testA", decimal.NewFromFloat(-111.66), period.Weekly, time.UnixMilli(5000)),
			periodicflow.New(uuid.New(), "testB", decimal.NewFromFloat(-222.66), period.Weekly, time.UnixMilli(4000)),
			periodicflow.New(testId, "test1", decimal.NewFromFloat(666.66), period.Weekly, testTime),
			periodicflow.New(uuid.New(), "testC", decimal.NewFromFloat(444.66), period.Weekly, time.UnixMilli(10000)),
			periodicflow.New(uuid.New(), "testD", decimal.NewFromFloat(555.66), period.Weekly, time.UnixMilli(1000)),
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
