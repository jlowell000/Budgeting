package periodic_flow

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"jlowell000.github.io/budgeting/internal/model/period"
)

const (
	TEST_PF_ID     = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_PF_NAME   = "test name"
	TEST_PF_AMOUNT = 100.99
	TEST_PF_TIME   = "2006-01-23T15:04:05Z"
)

func TestPeriodicFlowToJSON(t *testing.T) {
	id := getUUID(TEST_PF_ID)
	timestamp := getTime(TEST_PF_TIME)
	for _, p := range period.Periods {
		expected := getTestJson(TEST_PF_ID, TEST_PF_NAME, TEST_PF_AMOUNT, p, TEST_PF_AMOUNT, TEST_PF_TIME)
		periodicFlow := PeriodicFlow{
			Id:               id,
			Name:             TEST_PF_NAME,
			Amount:           TEST_PF_AMOUNT,
			Period:           p,
			WeeklyAmount:     TEST_PF_AMOUNT,
			UpdatedTimestamp: timestamp,
		}
		actual := string(periodicFlow.ToJSON())

		assert.Equal(t, expected, actual)
	}
}

func TestPeriodicFlowFromJSON_data_there(t *testing.T) {
	id := getUUID(TEST_PF_ID)
	timestamp := getTime(TEST_PF_TIME)

	for _, p := range period.Periods {
		expected := PeriodicFlow{
			Id:               id,
			Name:             TEST_PF_NAME,
			Amount:           TEST_PF_AMOUNT,
			Period:           p,
			WeeklyAmount:     TEST_PF_AMOUNT,
			UpdatedTimestamp: timestamp,
		}
		actual := PeriodicFlowFromJSON([]byte(
			getTestJson(TEST_PF_ID, TEST_PF_NAME, TEST_PF_AMOUNT, p, TEST_PF_AMOUNT, TEST_PF_TIME),
		))

		assert.Equal(t, expected, actual)
	}
}

func TestPeriodicFlowFromJSON_partial_data_there(t *testing.T) {
	id := getUUID(TEST_PF_ID)
	timestamp := getTime(TEST_PF_TIME)

	for _, p := range period.Periods {
		expected := PeriodicFlow{
			Id:               id,
			Amount:           TEST_PF_AMOUNT,
			Period:           p,
			WeeklyAmount:     TEST_PF_AMOUNT,
			UpdatedTimestamp: timestamp,
		}
		actual := PeriodicFlowFromJSON([]byte(
			getTestJson(TEST_PF_ID, "", TEST_PF_AMOUNT, p, TEST_PF_AMOUNT, TEST_PF_TIME),
		))

		assert.Equal(t, expected, actual)
	}
}

func TestPeriodicFlowFromJSON_no_data(t *testing.T) {
	expected := PeriodicFlow{}
	actual := PeriodicFlowFromJSON([]byte(""))

	assert.Equal(t, expected, actual)
}

func getPFParsedValues() (uuid.UUID, time.Time) {
	id, err1 := uuid.Parse(TEST_PF_ID)
	if err1 != nil {
		log.Fatal(err1)
	}
	timestamp, err2 := time.Parse(time.RFC3339, TEST_PF_TIME)
	if err2 != nil {
		log.Fatal(err2)
	}
	return id, timestamp
}

func getTestJson(
	id string,
	name string,
	amount float64,
	p period.Period,
	weeklyAmount float64,
	time string,
) string {
	return "{\"id\":\"" + id +
		"\",\"name\":\"" + name +
		"\",\"amount\":" + fmt.Sprintf("%.2f", amount) +
		",\"period\":\"" + p.String() +
		"\",\"weekly_amount\":" + fmt.Sprintf("%.2f", weeklyAmount) +
		",\"updated_timestamp\":\"" + time + "\"}"
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
