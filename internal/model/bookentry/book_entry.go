package bookentry

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

/*
Struct for defining a BookEntry.
*/
type BookEntry struct {
	Id        uuid.UUID       `json:"id"`
	Amount    decimal.Decimal `json:"amount,omitempty"`
	Timestamp time.Time       `json:"timestamp,omitempty"`
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

func (b *BookEntry) String() string {
	return string(b.ToJSON())
}

/*
Rate of change between two entries.

Returns in units of `Money / Millisecond`
*/
func RateOfChange(a BookEntry, b BookEntry) decimal.Decimal {
	timeDiff := b.Timestamp.UnixMilli() - a.Timestamp.UnixMilli()
	return (b.Amount.Sub(a.Amount)).Div(decimal.NewFromInt(timeDiff))
}
