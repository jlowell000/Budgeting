package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_LIST_JSON = "{\"Entries\":[" + TEST_JSON + "," + TEST_JSON + "]}"
)

func TestEntryListToJSON(t *testing.T) {
	id, timestamp := getParsedValues()

	expected := TEST_LIST_JSON
	entry := EntryList{
		Entries: []Entry{
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

	actual := string(entry.ToJSON())

	assert.Equal(t, expected, actual)
}

func TestEntryListFromJSON_data_exists(t *testing.T) {
	id, timestamp := getParsedValues()

	expected := EntryList{
		Entries: []Entry{
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
	actual := EntryListFromJSON([]byte(TEST_LIST_JSON))

	assert.Equal(t, expected, actual)
}

func TestEntryListFromJSON_no_data(t *testing.T) {
	expected := EntryList{}
	actual := EntryListFromJSON([]byte(""))

	assert.Equal(t, expected, actual)
}

func TestAdd(t *testing.T) {
	id, timestamp := getParsedValues()
	expected := EntryList{
		Entries: []Entry{
			{
				Id:        id,
				Timestamp: timestamp.Add(5),
				Content:   TEST_CONTENT,
			},
			{
				Id:        id,
				Timestamp: timestamp,
				Content:   TEST_CONTENT,
			},
		},
	}

	actual := EntryList{
		Entries: []Entry{
			{
				Id:        id,
				Timestamp: timestamp,
				Content:   TEST_CONTENT,
			},
		},
	}
	actual.Add(Entry{
		Id:        id,
		Timestamp: timestamp.Add(5),
		Content:   TEST_CONTENT,
	})

	assert.Equal(t, expected, actual)
}
