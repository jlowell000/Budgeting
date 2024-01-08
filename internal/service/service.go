package service

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/model/data"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
)

type DataServiceInterface interface {
	GetData() *data.DataModel
	SaveData(*data.DataModel) *data.DataModel
}

type PeriodicFlowServiceInterface interface {
	Create(name string, amount decimal.Decimal, period period.Period) *periodicflow.PeriodicFlow
	Get(id uuid.UUID) *periodicflow.PeriodicFlow
	GetAll() []*periodicflow.PeriodicFlow
	GetAllSortedByDate() []*periodicflow.PeriodicFlow
	GetTotalWeeklyInflow() decimal.Decimal
	GetTotalWeeklyOutflow() decimal.Decimal
	GetTotalWeeklyFlow() decimal.Decimal
	Update(id uuid.UUID, name string, amount decimal.Decimal, period period.Period) *periodicflow.PeriodicFlow
	Delete(id uuid.UUID)
}

type AccountServiceInterface interface {
	Create(name string, excludable bool) *account.Account
	Get(id uuid.UUID) *account.Account
	GetAll() []*account.Account
	GetAllSortedByDate() []*account.Account
	GetTotal(exclude bool) decimal.Decimal
	Update(id uuid.UUID, name string, excludable bool) *account.Account
	AddBookEntry(id uuid.UUID, amount decimal.Decimal) *account.Account
	Delete(id uuid.UUID)
}
