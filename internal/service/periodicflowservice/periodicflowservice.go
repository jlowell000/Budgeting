package periodicflowservice

import (
	"cmp"
	"slices"
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

func (p *PeriodicFlowService) Create(
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	d := p.Dataservice.GetData()
	f := periodicflow.New(p.GetId(), name, amount, period, p.GetTime())
	d.Flows = append(d.Flows, f)
	slices.SortFunc(d.Flows, compareFlowId)
	p.Dataservice.SaveData(d)
	return f
}

func (p *PeriodicFlowService) Get(id uuid.UUID) *periodicflow.PeriodicFlow {
	d := p.Dataservice.GetData()
	i := findPeriodicFlow(id, d.Flows)
	return d.Flows[i]
}

func (p *PeriodicFlowService) GetAll() []*periodicflow.PeriodicFlow {
	d := p.Dataservice.GetData()
	slices.SortFunc(d.Flows, compareFlowId)
	return d.Flows
}

func (p *PeriodicFlowService) GetAllSortedByDate() []*periodicflow.PeriodicFlow {
	f := slices.Clone(p.Dataservice.GetData().Flows)
	slices.SortFunc(f, compareFlowTime)
	return f
}

func (p *PeriodicFlowService) Update(
	id uuid.UUID,
	name string,
	amount decimal.Decimal,
	period period.Period,
) *periodicflow.PeriodicFlow {
	p.Get(id).Update(
		name,
		amount,
		period,
		p.GetTime(),
	)
	p.Dataservice.SaveData(p.Dataservice.GetData())
	return p.Get(id)

}

func (p *PeriodicFlowService) Delete(id uuid.UUID) {
	d := p.Dataservice.GetData()
	d.Flows = slices.DeleteFunc(
		d.Flows,
		func(f *periodicflow.PeriodicFlow) bool { return f.Id == id },
	)
	p.Dataservice.SaveData(d)
}

func sort(flows []*periodicflow.PeriodicFlow) []*periodicflow.PeriodicFlow {

	return flows
}

func findPeriodicFlow(id uuid.UUID, flows []*periodicflow.PeriodicFlow) int {
	slices.SortFunc(
		flows,
		compareFlowId,
	)
	n, _ := slices.BinarySearchFunc(
		flows,
		&periodicflow.PeriodicFlow{Id: id},
		compareFlowId,
	)
	return n
}

func compareFlowId(a, b *periodicflow.PeriodicFlow) int {
	return cmp.Compare(a.Id.String(), b.Id.String())
}

func compareFlowTime(a, b *periodicflow.PeriodicFlow) int {
	return cmp.Compare(a.Id.String(), b.Id.String())
}
