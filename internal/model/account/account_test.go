package account

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
)

const (
	TEST_ID     = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_NAME   = "test name"
	TEST_FLAG   = false
	TEST_AMOUNT = 100.99
	TEST_TIME   = "2006-01-23T15:04:05Z"
)

func TestAccountToJSON(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	expected := getTestJson(TEST_ID, TEST_NAME, TEST_FLAG, TEST_AMOUNT, TEST_TIME)
	periodicFlow := Account{
		Id:         id,
		Name:       TEST_NAME,
		Excludable: TEST_FLAG,
		Book: []bookentry.BookEntry{
			{
				Id:        id,
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
		},
		UpdatedTimestamp: timestamp,
	}
	actual := string(periodicFlow.ToJSON())

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_data_there(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	expected := Account{
		Id:         id,
		Name:       TEST_NAME,
		Excludable: TEST_FLAG,
		Book: []bookentry.BookEntry{
			{
				Id:        id,
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
		},
		UpdatedTimestamp: timestamp,
	}
	actual := FromJSON([]byte(
		getTestJson(TEST_ID, TEST_NAME, TEST_FLAG, TEST_AMOUNT, TEST_TIME),
	))

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_partial_data_there(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)

	expected := Account{
		Id: id,
		Book: []bookentry.BookEntry{
			{
				Id:        id,
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
		},
		UpdatedTimestamp: timestamp,
	}
	actual := FromJSON([]byte(
		getTestJson(TEST_ID, "", TEST_FLAG, TEST_AMOUNT, TEST_TIME),
	))

	assert.Equal(t, expected, actual)
}

func TestAccountFromJSON_no_data(t *testing.T) {
	expected := Account{}
	actual := FromJSON([]byte(""))

	assert.Equal(t, expected, actual)
}

func Test_GetLatestBookEntry(t *testing.T) {
	id := getUUID(TEST_ID)
	timestamp := getTime(TEST_TIME)
	expected := bookentry.BookEntry{
		Id:        uuid.New(),
		Amount:    666.6,
		Timestamp: time.Now(),
	}
	account := Account{
		Id: id,
		Book: []bookentry.BookEntry{
			{
				Id:        uuid.New(),
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
			{
				Id:        uuid.New(),
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
			expected,
			{
				Id:        uuid.New(),
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
			{
				Id:        uuid.New(),
				Amount:    TEST_AMOUNT,
				Timestamp: timestamp,
			},
		},
		UpdatedTimestamp: timestamp,
	}
	actual := account.GetLatestBookEntry()

	assert.Equal(t, expected, actual)
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
	excludable bool,
	amount float64,
	time string,
) string {
	return "{\"id\":\"" + id + "\"," +
		"\"name\":\"" + name + "\"," +
		"\"excludable\":" + fmt.Sprintf("%t", excludable) + "," +
		"\"book\":[" +
		"{\"id\":\"" + id +
		"\",\"amount\":" + fmt.Sprintf("%.2f", amount) +
		",\"timestamp\":\"" + time + "\"}" +
		"]," +
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
