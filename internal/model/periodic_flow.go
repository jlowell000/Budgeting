package model

import (
	// "fmt"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

type PeriodicFlow struct {
	Id                   uuid.UUID
	Name                 string
	Amount               float64
	Period               Period
	WeeklyAmount         float64
	LastUpdatedTimestamp time.Time
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
