package accountservice

import (
	"cmp"
	"slices"
	"time"

	"github.com/google/uuid"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/service"
)

type AccountService struct {
	Dataservice service.DataServiceInterface

	GetTime func() time.Time
	GetId   func() uuid.UUID
}

func (a *AccountService) Create(name string, excludable bool) *account.Account {
	d := a.Dataservice.GetData()
	newId := a.GetId()
	newA := account.New(
		newId,
		name,
		excludable,
		a.GetTime(),
	)
	d.Accounts = append(d.Accounts, newA)
	slices.SortFunc(d.Accounts, compareAccountId)
	a.Dataservice.SaveData(d)
	return a.Get(newId)
}

func (a *AccountService) Get(id uuid.UUID) *account.Account {
	d := a.Dataservice.GetData()
	i := findAccount(id, d.Accounts)
	return d.Accounts[i]
}

func (a *AccountService) GetAll() []*account.Account {
	d := a.Dataservice.GetData()
	slices.SortFunc(d.Accounts, compareAccountId)
	return d.Accounts
}

func (a *AccountService) GetAllSortedByDate() []*account.Account {
	l := slices.Clone(a.Dataservice.GetData().Accounts)
	slices.SortFunc(l, compareAccountTime)
	return l
}

func (a *AccountService) Update(id uuid.UUID, name string, excludable bool) *account.Account {
	a.Get(id).Update(
		name,
		excludable,
		a.GetTime(),
	)
	a.Dataservice.SaveData(a.Dataservice.GetData())
	return a.Get(id)
}

func (a *AccountService) Delete(id uuid.UUID) {
	d := a.Dataservice.GetData()
	d.Accounts = slices.DeleteFunc(
		d.Accounts,
		func(c *account.Account) bool { return c.Id == id },
	)
	a.Dataservice.SaveData(d)
}

func findAccount(id uuid.UUID, accounts []*account.Account) int {
	slices.SortFunc(accounts, compareAccountId)
	n, _ := slices.BinarySearchFunc(
		accounts,
		&account.Account{Id: id},
		compareAccountId,
	)
	return n
}

func compareAccountId(a, b *account.Account) int {
	return cmp.Compare(a.Id.String(), b.Id.String())
}

func compareAccountTime(a, b *account.Account) int {
	return cmp.Compare(b.UpdatedTimestamp.UnixMilli(), a.UpdatedTimestamp.UnixMilli())
}
