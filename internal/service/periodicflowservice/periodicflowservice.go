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
	GetId   func() uuid.UUID
}

func (p *PeriodicFlowService) CreatePeriodicFlow(
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	d := p.Dataservice.GetData()
	f := periodicflow.New(p.GetId(), name, amount, period, p.GetTime())
	d.Flows = append(d.Flows, f)
	p.Dataservice.SaveData(d)
	return f
}

func (p *PeriodicFlowService) GetPeriodicFlows() []*periodicflow.PeriodicFlow {
	return p.Dataservice.GetData().Flows
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
