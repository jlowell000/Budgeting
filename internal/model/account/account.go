package account

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
)

/*
Struct for defining a Account.
*/
type Account struct {
	Id               uuid.UUID             `json:"id"`
	Name             string                `json:"name,omitempty"`
	Excludable       bool                  `json:"excludable"`
	Book             []bookentry.BookEntry `json:"book,omitempty"`
	UpdatedTimestamp time.Time             `json:"updated_timestamp,omitempty"`
}

// Returns JSON encoding of Account
func (account *Account) ToJSON() []byte {
	data, err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Parses an Account from JSON
func FromJSON(data []byte) Account {
	var account Account
	json.Unmarshal(data, &account)
	return account
}

/*
Sum the given Accounts' amounts
*/
func Sum(accounts []Account) decimal.Decimal {
	sum := decimal.NewFromInt(0)
	for _, a := range accounts {
		sum = sum.Add(a.GetLatestBookEntry().Amount)
	}
	return sum
}

// Returns JSON encoding of Account
func (account *Account) GetLatestBookEntry() bookentry.BookEntry {

	var bookentry bookentry.BookEntry
	for _, b := range account.Book {
		if b.Timestamp.After(bookentry.Timestamp) {
			bookentry = b
		}
	}
	return bookentry
}

func (a *Account) String() string {
	return string(a.ToJSON())
}
