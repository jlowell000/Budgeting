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

/*
Sum the given Accounts' amounts excludes all Excludable accounts
*/
func SumExclusion(accounts []Account) decimal.Decimal {
	sum := decimal.NewFromInt(0)
	for _, a := range accounts {
		if !a.Excludable {
			sum = sum.Add(a.GetLatestBookEntry().Amount)
		}
	}
	return sum
}

// Returns Latest Book entry for Account
func (account *Account) GetLatestBookEntry() bookentry.BookEntry {

	var bookentry bookentry.BookEntry
	for _, b := range account.Book {
		if bookentry.Timestamp.Before(b.Timestamp) {
			bookentry = b
		}
	}
	return bookentry
}

// Returns Earliest Book entry for Account
func (account *Account) GetEarliestBookEntry() bookentry.BookEntry {
	var bookentry bookentry.BookEntry
	if len(account.Book) > 0 {
		bookentry = account.Book[0]
		for _, b := range account.Book {
			if bookentry.Timestamp.After(b.Timestamp) {
				bookentry = b
			}
		}
	}
	return bookentry
}

func (account *Account) GetBookEndEntries() (bookentry.BookEntry, bookentry.BookEntry) {
	var first bookentry.BookEntry
	var second bookentry.BookEntry

	if len(account.Book) > 0 {
		first = account.Book[0]
		second = account.Book[0]
		for _, b := range account.Book {
			if first.Timestamp.After(b.Timestamp) {
				first = b
			}
			if second.Timestamp.Before(b.Timestamp) {
				second = b
			}
		}
	}
	return first, second
}

func (account *Account) RateOfChange() decimal.Decimal {
	if len(account.Book) < 2 {
		return decimal.NewFromInt(0)
	} else {
		a, b := account.GetBookEndEntries()
		return bookentry.RateOfChange(a, b)
	}
}

func (a *Account) String() string {
	return string(a.ToJSON())
}
