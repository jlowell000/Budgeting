package model

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_ID      = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_TIME    = "2006-01-23T15:04:05Z"
	TEST_CONTENT = "test content"
	TEST_JSON    = "{\"Id\":\"" + TEST_ID + "\",\"Timestamp\":\"" + TEST_TIME + "\",\"Content\":\"" + TEST_CONTENT + "\"}"
)

func TestEntryToJSON(t *testing.T) {
	id, timestamp := getParsedValues()

	expected := TEST_JSON
	entry := Entry{
		Id:        id,
		Timestamp: timestamp,
		Content:   TEST_CONTENT,
	}
	actual := string(entry.ToJSON())

	assert.Equal(t, expected, actual)
}

func TestEntryFromJSON_data_there(t *testing.T) {
	id, timestamp := getParsedValues()

	expected := Entry{
		Id:        id,
		Timestamp: timestamp,
		Content:   TEST_CONTENT,
	}
	actual := EntryFromJSON([]byte(TEST_JSON))

	assert.Equal(t, expected, actual)
}

func TestEntryFromJSON_no_data(t *testing.T) {
	expected := Entry{}
	actual := EntryFromJSON([]byte(""))

	assert.Equal(t, expected, actual)
}

func getParsedValues() (uuid.UUID, time.Time) {
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
