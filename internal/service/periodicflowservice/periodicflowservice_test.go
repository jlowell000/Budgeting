package periodicflowservice

import (
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
	}
)

func Test_UpdatePeriodicFlow(t *testing.T) {
	resetData()
	expected := periodicflow.New(testId, "testEdited", decimal.NewFromFloat(111.11), period.Monthly, testTime2)
	actual := subject.UpdatePeriodicFlow(testId, "testEdited", decimal.NewFromFloat(111.11), period.Monthly)

	assert.Equal(t, 1, getDataCount, "getDataCount")
	assert.Equal(t, 1, saveDataCount, "saveDataCount")
	assert.Equal(t, *expected, *actual, "not equal object")
}

type mockDataService struct{}

func resetData() {
	getDataCount = 0
	saveDataCount = 0
	testData = &data.DataModel{
		Flows: []*periodicflow.PeriodicFlow{
			periodicflow.New(testId, "test1", decimal.NewFromFloat(666.66), period.Weekly, testTime),
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
