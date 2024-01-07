package data

import (
	"encoding/json"
	"log"

	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
)

type DataModel struct {
	Flows    []*periodicflow.PeriodicFlow `json:"flows"`
	Accounts []*account.Account           `json:"accounts"`
}

func FromJSON(data []byte) DataModel {
	var d DataModel
	json.Unmarshal(data, &d)
	return d
}

func (d *DataModel) ToJSON() []byte {
	data, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (d *DataModel) String() string {
	return string(d.ToJSON())
}
