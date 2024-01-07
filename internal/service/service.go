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
	Create(string, decimal.Decimal, period.Period) *periodicflow.PeriodicFlow
	Get(uuid.UUID) *periodicflow.PeriodicFlow
	GetAll() []*periodicflow.PeriodicFlow
	GetAllSortedByDate() []*periodicflow.PeriodicFlow
	Update(uuid.UUID, string, decimal.Decimal, period.Period) *periodicflow.PeriodicFlow
	Delete(uuid.UUID)
}
