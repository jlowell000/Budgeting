package model

import (
	// "fmt"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

type PeriodicFlow struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name,omitempty"`
	Amount           float64   `json:"amount,omitempty"`
	Period           Period    `json:"period"`
	WeeklyAmount     float64   `json:"weekly_amount,omitempty"`
	UpdatedTimestamp time.Time `json:"updated_timestamp,omitempty"`
}

// Returns JSON encoding of PeriodicFlow
func (flow *PeriodicFlow) ToJSON() []byte {
	data, err := json.Marshal(flow)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Parses an PeriodicFlow from JSON
func PeriodicFlowFromJSON(data []byte) PeriodicFlow {
	var flow PeriodicFlow
	json.Unmarshal(data, &flow)
	return flow
}
