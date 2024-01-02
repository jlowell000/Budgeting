package service

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"jlowell000.github.io/budgeting/internal/model"
)

const (
	TEST_ID        = "16cfd708-db6d-42fd-8ad1-55316690520c"
	TEST_TIME      = "2006-01-23T15:04:05Z"
	TEST_CONTENT   = "test content"
	TEST_FILE_NAME = "test_file.txt"
	TEST_JSON      = "{\"Id\":\"" + TEST_ID + "\",\"Timestamp\":\"" + TEST_TIME + "\",\"Content\":\"" + TEST_CONTENT + "\"}"
	TEST_LIST_JSON = "{\"Entries\":[" + TEST_JSON + "," + TEST_JSON + "]}"
)

var (
	numGetEntryListJSON   int = 0
	numPutEntryListJSON   int = 0
	numParseEntryListJSON int = 0
	numGetNewId           int = 0
	numGetTime            int = 0
)

func TestCreateEntry(t *testing.T) {
	id, timestamp := getParsedValues()
	setUpMocking()
	expected := model.Entry{
		Id:        id,
		Timestamp: timestamp,
		Content:   TEST_CONTENT,
	}

	actual := CreateEntry(TEST_CONTENT, TEST_FILE_NAME)

	assert.Equal(t, expected, actual)
	assert.Equal(t, 1, numGetEntryListJSON, "GetEntryListJSON")
	assert.Equal(t, 1, numParseEntryListJSON, "ParseEntryListJSON")
	assert.Equal(t, 1, numGetNewId, "GetNewId")
	assert.Equal(t, 1, numGetTime, "GetTime")
	assert.Equal(t, 1, numPutEntryListJSON, "PutEntryListJSON")
}

func TestGetLatestEntry(t *testing.T) {
	id, timestamp := getParsedValues()
	setUpMocking()
	expected := model.Entry{
		Id:        id,
		Timestamp: timestamp,
		Content:   TEST_CONTENT,
	}

	actual := GetLatestEntry(TEST_FILE_NAME)

	assert.Equal(t, expected, actual)
	assert.Equal(t, 1, numGetEntryListJSON, "GetEntryListJSON")
	assert.Equal(t, 1, numParseEntryListJSON, "ParseEntryListJSON")
	assert.Equal(t, 0, numGetNewId, "GetNewId")
	assert.Equal(t, 0, numGetTime, "GetTime")
	assert.Equal(t, 0, numPutEntryListJSON, "PutEntryListJSON")

}

func TestGetEntryList(t *testing.T) {
	id, timestamp := getParsedValues()
	setUpMocking()
	expected := model.EntryList{
		Entries: []model.Entry{
			{
				Id:        id,
				Timestamp: timestamp,
				Content:   TEST_CONTENT,
			},
			{
				Id:        id,
				Timestamp: timestamp,
				Content:   TEST_CONTENT,
			},
		},
	}

	actual := GetEntryList(TEST_FILE_NAME)

	assert.Equal(t, expected, actual)
	assert.Equal(t, 1, numGetEntryListJSON, "GetEntryListJSON")
	assert.Equal(t, 1, numParseEntryListJSON, "ParseEntryListJSON")
	assert.Equal(t, 0, numGetNewId, "GetNewId")
	assert.Equal(t, 0, numGetTime, "GetTime")
	assert.Equal(t, 0, numPutEntryListJSON, "PutEntryListJSON")

}

func TestSaveEntryListToFile(t *testing.T) {
	id, timestamp := getParsedValues()
	setUpMocking()

	SaveEntryListToFile(model.EntryList{
		Entries: []model.Entry{
			{
				Id:        id,
				Timestamp: timestamp,
				Content:   TEST_CONTENT,
			},
			{
				Id:        id,
				Timestamp: timestamp,
				Content:   TEST_CONTENT,
			},
		},
	}, TEST_FILE_NAME)

	assert.Equal(t, 0, numGetEntryListJSON, "GetEntryListJSON")
	assert.Equal(t, 0, numParseEntryListJSON, "ParseEntryListJSON")
	assert.Equal(t, 0, numGetNewId, "GetNewId")
	assert.Equal(t, 0, numGetTime, "GetTime")
	assert.Equal(t, 1, numPutEntryListJSON, "PutEntryListJSON")

}

func setUpMocking() {
	id, timestamp := getParsedValues()
	getEntryListJSON = func(fileName string) []byte {
		numGetEntryListJSON++
		return []byte(TEST_LIST_JSON)
	}
	putEntryListJSON = func(data []byte, fileName string) {
		numPutEntryListJSON++
		return
	}

	parseEntryListJSON = func(data []byte) model.EntryList {
		numParseEntryListJSON++
		return model.EntryList{
			Entries: []model.Entry{
				{
					Id:        id,
					Timestamp: timestamp,
					Content:   TEST_CONTENT,
				},
				{
					Id:        id,
					Timestamp: timestamp,
					Content:   TEST_CONTENT,
				},
			},
		}
	}

	getNewId = func() uuid.UUID {
		numGetNewId++
		return id
	}
	getTime = func() time.Time {
		numGetTime++
		return timestamp
	}

	numGetEntryListJSON = 0
	numPutEntryListJSON = 0
	numParseEntryListJSON = 0
	numGetNewId = 0
	numGetTime = 0
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
