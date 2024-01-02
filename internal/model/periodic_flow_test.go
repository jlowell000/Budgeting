package model

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_PF_ID      = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_PF_NAME    = "test name"
	TEST_PF_AMOUNT  = 100.99
	TEST_PF_PERIOD  = Weekly
	TEST_PF_TIME    = "2006-01-23T15:04:05Z"
	TEST_PF_CONTENT = "test content"
)

var TEST_PF_JSON = "{\"Id\":\"" + TEST_PF_ID +
	"\",\"Name\":\"" + TEST_PF_NAME +
	"\",\"Amount\":" + fmt.Sprintf("%.2f", TEST_PF_AMOUNT) +
	",\"Period\":\"" + TEST_PF_PERIOD.String() +
	"\",\"WeeklyAmount\":" + fmt.Sprintf("%.2f", TEST_PF_AMOUNT) +
	",\"LastUpdatedTimestamp\":\"" + TEST_PF_TIME + "\"}"

func TestPeriodicFlowToJSON(t *testing.T) {
	id, timestamp := getParsedValues()

	expected := TEST_PF_JSON
	periodicFlow := PeriodicFlow{
		Id:                   id,
		Name:                 TEST_PF_NAME,
		Amount:               TEST_PF_AMOUNT,
		Period:               TEST_PF_PERIOD,
		WeeklyAmount:         TEST_PF_AMOUNT,
		LastUpdatedTimestamp: timestamp,
	}
	actual := string(periodicFlow.ToJSON())

	assert.Equal(t, expected, actual)
}

func TestPeriodicFlowFromJSON_data_there(t *testing.T) {
	id, timestamp := getParsedValues()

	expected := PeriodicFlow{
		Id:                   id,
		Name:                 TEST_PF_NAME,
		Amount:               TEST_PF_AMOUNT,
		Period:               TEST_PF_PERIOD,
		WeeklyAmount:         TEST_PF_AMOUNT,
		LastUpdatedTimestamp: timestamp,
	}
	actual := PeriodicFlowFromJSON([]byte(TEST_PF_JSON))

	assert.Equal(t, expected, actual)
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
