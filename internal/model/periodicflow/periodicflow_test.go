package periodicflow

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"jlowell000.github.io/budgeting/internal/model/period"
)

const (
	TEST_ID   = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_NAME = "test name"
	TEST_TIME = "2006-01-23T15:04:05Z"
)

var TEST_AMOUNT = decimal.NewFromFloat(666.66)

func TestPeriodicFlowToJSON(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)
	for _, p := range period.Periods {
		expected := getTestJson(TEST_ID, TEST_NAME, TEST_AMOUNT, p, TEST_AMOUNT, TEST_TIME)
		periodicFlow := PeriodicFlow{
			Id:               id,
			Name:             TEST_NAME,
			Amount:           TEST_AMOUNT,
			Period:           p,
			WeeklyAmount:     TEST_AMOUNT,
			UpdatedTimestamp: timestamp,
		}
		actual := string(periodicFlow.ToJSON())

		assert.Equal(t, expected, actual)
	}
}

func TestPeriodicFlowFromJSON_data_there(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	for _, p := range period.Periods {
		expected := PeriodicFlow{
			Id:               id,
			Name:             TEST_NAME,
			Amount:           TEST_AMOUNT,
			Period:           p,
			WeeklyAmount:     TEST_AMOUNT,
			UpdatedTimestamp: timestamp,
		}
		actual := FromJSON([]byte(
			getTestJson(TEST_ID, TEST_NAME, TEST_AMOUNT, p, TEST_AMOUNT, TEST_TIME),
		))

		assert.Equal(t, expected, actual)
	}
}

func TestPeriodicFlowFromJSON_partial_data_there(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	for _, p := range period.Periods {
		expected := PeriodicFlow{
			Id:               id,
			Amount:           TEST_AMOUNT,
			Period:           p,
			WeeklyAmount:     TEST_AMOUNT,
			UpdatedTimestamp: timestamp,
		}
		actual := FromJSON([]byte(
			getTestJson(TEST_ID, "", TEST_AMOUNT, p, TEST_AMOUNT, TEST_TIME),
		))

		assert.Equal(t, expected, actual)
	}
}

func TestPeriodicFlowFromJSON_no_data(t *testing.T) {
	expected := PeriodicFlow{}
	actual := FromJSON([]byte(""))

	assert.Equal(t, expected, actual)
}

func TestPeriodicFlowFConstructor_properly_sets_weekly_amount(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	for _, p := range period.Periods {
		expected := PeriodicFlow{
			Id:               id,
			Name:             TEST_NAME,
			Amount:           TEST_AMOUNT,
			Period:           p,
			WeeklyAmount:     TEST_AMOUNT.Mul(p.WeeklyAmount()),
			UpdatedTimestamp: timestamp,
		}
		actual := *New(id, TEST_NAME, TEST_AMOUNT, p, timestamp)

		assert.Equal(t, expected, actual)
	}
}

func TestPeriodicFlow_Sum_different_periods(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	var expected decimal.Decimal
	var flows []PeriodicFlow
	for _, p := range period.Periods {
		expected = expected.Add(TEST_AMOUNT.Mul(p.WeeklyAmount()))
		flows = append(flows, *New(id, TEST_NAME, TEST_AMOUNT, p, timestamp))
	}
	actual := Sum(flows)
	assert.Equal(t, expected, actual)
}

func TestPeriodicFlow_Projected_change_different_periods(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)
	projectAmount := decimal.NewFromFloat(30.0)
	projectPeriod := period.Monthly

	var expected decimal.Decimal
	var flows []PeriodicFlow
	for _, p := range period.Periods {
		expected = expected.Add(TEST_AMOUNT.Mul(p.WeeklyAmount()))
		flows = append(flows, *New(id, TEST_NAME, TEST_AMOUNT, p, timestamp))
	}
	expected = expected.Mul(projectAmount).Mul(projectPeriod.WeeklyAmount())
	actual := ProjectedChange(flows, projectAmount, projectPeriod)
	assert.Equal(t, expected, actual)
}

func TestPeriodicFlow_Update(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)
	expected := New(id, TEST_NAME, TEST_AMOUNT, period.Daily, timestamp)
	actual := New(id, TEST_NAME, TEST_AMOUNT, period.Daily, timestamp)
	assert.Equal(t, expected, actual)
	for _, p := range period.Periods {
		expected = New(id, TEST_NAME, TEST_AMOUNT, p, timestamp)
		actual = actual.Update(actual.Name, actual.Amount, p, timestamp)
		assert.Equal(t, expected, actual)
	}
}

func getPFParsedValues() (uuid.UUID, time.Time) {
	id, err1 := uuid.Parse(TEST_ID)
	if err1 != nil {
		log.Fatal(err1)
	}
	timestamp, err2 := time.Parse(time.RFC3339, TEST_TIME)
	if err2 != nil {
		log.Fatal(err2)
	}
	return id, timestamp
}

func getTestJson(
	id string,
	name string,
	amount decimal.Decimal,
	p period.Period,
	weeklyAmount decimal.Decimal,
	time string,
) string {
	return "{\"id\":\"" + id + "\"," +
		"\"name\":\"" + name + "\"," +
		"\"amount\":\"" + amount.String() + "\"," +
		"\"period\":\"" + p.String() + "\"," +
		"\"weekly_amount\":\"" + weeklyAmount.String() + "\"," +
		"\"updated_timestamp\":\"" + time + "\"}"
}

func getUUID(s string) uuid.UUID {
	id, err1 := uuid.Parse(s)
	if err1 != nil {
		log.Fatal(err1)
	}
	return id
}

func getTime(s string) time.Time {
	timestamp, err2 := time.Parse(time.RFC3339, s)
	if err2 != nil {
		log.Fatal(err2)
	}
	return timestamp
}
