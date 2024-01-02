package bookentry

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

/*
Struct for defining a BookEntry.
*/
type BookEntry struct {
	Id        uuid.UUID `json:"id"`
	Amount    float64   `json:"amount,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// Returns JSON encoding of BookEntry
func (bookEntry *BookEntry) ToJSON() []byte {
	data, err := json.Marshal(bookEntry)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Parses an BookEntry from JSON
func FromJSON(data []byte) BookEntry {
	var bookEntry BookEntry
	json.Unmarshal(data, &bookEntry)
	return bookEntry
}