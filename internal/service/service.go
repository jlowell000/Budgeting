package service

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/data"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
)

type DataServiceInterface interface {
	GetData() *data.DataModel
	SaveData(*data.DataModel) *data.DataModel
}

type PeriodicFlowServiceInterface interface {
	UpdatePeriodicFlow(
		id uuid.UUID,
		name string,
		amount decimal.Decimal,
		period period.Period,
	) *periodicflow.PeriodicFlow
}
