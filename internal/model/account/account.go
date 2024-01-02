package account

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"jlowell000.github.io/budgeting/internal/model/bookentry"
)

/*
Struct for defining a Account.
*/
type Account struct {
	Id               uuid.UUID             `json:"id"`
	Name             string                `json:"name,omitempty"`
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
