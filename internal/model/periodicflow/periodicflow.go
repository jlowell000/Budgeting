package periodicflow

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/period"
)

/*
Struct for defining a periodic flow.
A change in amount over a period.
*/
type PeriodicFlow struct {
	Id               uuid.UUID       `json:"id"`
	Name             string          `json:"name,omitempty"`
	Amount           decimal.Decimal `json:"amount,omitempty"`
	Period           period.Period   `json:"period"`
	MonthlyAmount    decimal.Decimal `json:"monthly_amount,omitempty"`
	UpdatedTimestamp time.Time       `json:"updated_timestamp,omitempty"`
}

/*
Returns PeriodicFlow of PeriodicFlow
*/
func New(
	id uuid.UUID,
	name string,
	amount decimal.Decimal,
	period period.Period,
	createTime time.Time,
) *PeriodicFlow {
	return &PeriodicFlow{
		Id:               id,
		Name:             name,
		Amount:           amount,
		Period:           period,
		MonthlyAmount:    amount.Mul(period.MonthlyAmount()),
		UpdatedTimestamp: createTime,
	}
}

func (f *PeriodicFlow) Update(
	name string,
	amount decimal.Decimal,
	period period.Period,
	updateTime time.Time,
) *PeriodicFlow {
	f.Name = name
	f.Amount = amount
	f.Period = period
	f.MonthlyAmount = f.Amount.Mul(f.Period.MonthlyAmount())
	f.UpdatedTimestamp = updateTime
	return f
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
func Sum(flows []*PeriodicFlow) decimal.Decimal {
	sum := decimal.NewFromInt(0)
	for _, f := range flows {
		sum = sum.Add(f.MonthlyAmount)
	}
	return sum
}

/*
Calculate projected change over time
*/
func ProjectedChange(
	flows []*PeriodicFlow,
	amount decimal.Decimal,
	period period.Period,
) decimal.Decimal {
	return Sum(flows).DivRound(period.MonthlyAmount(), 6).Mul(amount)
}

func (p *PeriodicFlow) String() string {
	return string(p.ToJSON())
}
