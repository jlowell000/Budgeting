package bookentry

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_ID   = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_NAME = "test name"
	TEST_TIME = "2006-01-23T15:04:05Z"
)

var TEST_AMOUNT = decimal.NewFromFloat(100.99)

func TestAccountToJSON(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	expected := getTestJson(TEST_ID, TEST_AMOUNT, TEST_TIME)
	bookEntry := BookEntry{
		Id:        id,
		Amount:    TEST_AMOUNT,
		Timestamp: timestamp,
	}
	actual := string(bookEntry.ToJSON())

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_data_there(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	expected := BookEntry{
		Id:        id,
		Amount:    TEST_AMOUNT,
		Timestamp: timestamp,
	}
	actual := FromJSON([]byte(
		getTestJson(TEST_ID, TEST_AMOUNT, TEST_TIME),
	))

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_partial_data_there(t *testing.T) {
	id := getUUID(TEST_ID)

	expected := BookEntry{
		Id:     id,
		Amount: TEST_AMOUNT,
	}
	actual := FromJSON([]byte(
		getTestJson(TEST_ID, TEST_AMOUNT, ""),
	))

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_no_data(t *testing.T) {
	expected := BookEntry{}
	actual := FromJSON([]byte(""))

	assert.Equal(t, expected, actual)
}

func TestRateOfChange(t *testing.T) {
	testSize := 100
	testSizeSlice := make([]bool, testSize)
	a := &BookEntry{
		Amount:    decimal.NewFromInt(0),
		Timestamp: time.UnixMilli(0),
	}

	for i := range testSizeSlice {
		for j := range testSizeSlice {
			di := decimal.NewFromInt(int64(i + 1))
			dj := int64(j + 1)
			expected := di.Div(decimal.NewFromInt(dj))
			actual := RateOfChange(
				a,
				&BookEntry{
					Amount:    di,
					Timestamp: time.UnixMilli(dj),
				},
			)

			assert.True(t, expected.Equals(actual))
		}
	}

	for i := range testSizeSlice {
		for j := range testSizeSlice {
			di := decimal.NewFromInt(int64(i + 1))
			dj := int64(testSize - j + 1)
			expected := di.Div(decimal.NewFromInt(dj))
			actual := RateOfChange(
				a,
				&BookEntry{
					Amount:    di,
					Timestamp: time.UnixMilli(dj),
				},
			)

			assert.True(t, expected.Equals(actual))
		}
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
	amount decimal.Decimal,
	time string,
) string {
	return "{\"id\":\"" + id + "\"," +
		"\"amount\":\"" + amount.String() + "\"," +
		"\"timestamp\":\"" + time + "\"}"
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
