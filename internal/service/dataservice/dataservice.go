package dataservice

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"jlowell000.github.io/budgeting/internal/io"
	"jlowell000.github.io/budgeting/internal/model/account"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
)

type DataModel struct {
	Flows    []*periodicflow.PeriodicFlow `json:"flows"`
	Accounts []*account.Account           `json:"accounts"`
}

var (
	getDataJSON = io.ReadFromFile
	putDataJSON = io.WriteToFile

	getNewId = uuid.New
	getTime  = time.Now
)

func GetDataFromFile(filename string) *DataModel {
	var data DataModel
	fileContents := getDataJSON(filename)
	json.Unmarshal(fileContents, &data)
	return &data
}

func SaveDataToFile(data *DataModel, filename string) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	putDataJSON(dataJSON, filename)
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
