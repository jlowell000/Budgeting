package periodicflowservice

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
	"jlowell000.github.io/budgeting/internal/service"
)

type PeriodicFlowService struct {
	Dataservice service.DataServiceInterface

	GetTime func() time.Time
}

func (p *PeriodicFlowService) UpdatePeriodicFlow(
	id uuid.UUID,
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	d := p.Dataservice.GetData()
	for i, f := range d.Flows {
		if f.Id == id {
			d.Flows[i] = f.Update(
				name,
				amount,
				period,
				p.GetTime(),
			)
			p.Dataservice.SaveData(d)
			return d.Flows[i]
		}
	}
	return nil
}
