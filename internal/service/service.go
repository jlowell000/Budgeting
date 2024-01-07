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
	Update(id uuid.UUID, name string, amount decimal.Decimal, period period.Period) *periodicflow.PeriodicFlow
	Delete(id uuid.UUID)
}

type AccountServiceInterface interface {
	Create(name string, excludable bool) *account.Account
	Get(id uuid.UUID) *account.Account
	GetAll() []*account.Account
	GetAllSortedByDate() []*account.Account
	Update(id uuid.UUID, name string, excludable bool) *account.Account
	Delete(id uuid.UUID)
}
