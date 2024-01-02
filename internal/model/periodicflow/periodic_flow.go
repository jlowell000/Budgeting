package periodic_flow

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"jlowell000.github.io/budgeting/internal/model/period"
)

/*
Struct for defining a periodic flow.
A change in amount over a period.
*/
type PeriodicFlow struct {
	Id               uuid.UUID     `json:"id"`
	Name             string        `json:"name,omitempty"`
	Amount           float64       `json:"amount,omitempty"`
	Period           period.Period `json:"period"`
	WeeklyAmount     float64       `json:"weekly_amount,omitempty"`
	UpdatedTimestamp time.Time     `json:"updated_timestamp,omitempty"`
}

/*
Returns PeriodicFlow of PeriodicFlow
*/
func New(
	id uuid.UUID,
	amount float64,
	period period.Period,
	timestamp time.Time,
) *PeriodicFlow {
	return &PeriodicFlow{
		Id:               id,
		Amount:           amount,
		Period:           period,
		WeeklyAmount:     amount * period.WeeklyAmount(),
		UpdatedTimestamp: timestamp,
	}
}

/*
Returns JSON encoding of PeriodicFlow
*/
func (flow *PeriodicFlow) ToJSON() []byte {
	data, err := json.Marshal(flow)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

/*
Parses an PeriodicFlow from JSON
*/
func FromJSON(data []byte) PeriodicFlow {
	var flow PeriodicFlow
	json.Unmarshal(data, &flow)
	return flow
}

/*
Sum the given Periodic Flows' weekly amounts
*/
func Sum(flows []PeriodicFlow) float64 {
	sum := 0.0
	for _, f := range flows {
		sum += f.WeeklyAmount
	}
	return sum
}

/*
Calculate projected change over time
*/
func ProjectedChange(flows []PeriodicFlow, amount float64, period period.Period) float64 {
	return Sum(flows) * period.WeeklyAmount() * amount
}
