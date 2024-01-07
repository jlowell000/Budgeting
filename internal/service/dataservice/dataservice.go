package dataservice

import (
	"encoding/json"
	"log"

	"jlowell000.github.io/budgeting/internal/model/data"
)

type DataService struct {
	Filename    string
	GetDataJSON func(fileName string) []byte
	PutDataJSON func(data []byte, fileName string)

	dataModel *data.DataModel
}

func (d *DataService) GetData() *data.DataModel {
	if d.dataModel == nil {
		var data data.DataModel
		fileContents := d.GetDataJSON(d.Filename)
		json.Unmarshal(fileContents, &data)
		d.dataModel = &data
	}

	return d.dataModel
}

func (d *DataService) SaveData(dataModel *data.DataModel) *data.DataModel {
	d.dataModel = dataModel
	dataJSON, err := json.Marshal(dataModel)
	if err != nil {
		log.Fatal(err)
	}
	d.PutDataJSON(dataJSON, d.Filename)
	return d.GetData()
}
