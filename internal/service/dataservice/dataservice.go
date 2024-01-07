package dataservice

import (
	"encoding/json"
	"log"

	"jlowell000.github.io/budgeting/internal/model/data"
)

type DataService struct {
	GetDataJSON func(fileName string) []byte
	PutDataJSON func(data []byte, fileName string)
}

func (d *DataService) GetDataFromFile(filename string) *data.DataModel {
	var data data.DataModel
	fileContents := d.GetDataJSON(filename)
	json.Unmarshal(fileContents, &data)
	return &data
}

func (d *DataService) SaveDataToFile(data *data.DataModel, filename string) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	d.PutDataJSON(dataJSON, filename)
}
