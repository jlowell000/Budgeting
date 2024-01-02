package model

import (
	// "fmt"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	Id        uuid.UUID
	Timestamp time.Time
	Content   string
}

// Returns JSON encoding of Entry
func (entry *Entry) ToJSON() []byte {
	data, err := json.Marshal(entry)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Parses an Entry from JSON
func EntryFromJSON(data []byte) Entry {
	var entry Entry
	json.Unmarshal(data, &entry)
	return entry
}
