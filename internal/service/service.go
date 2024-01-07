package service

import (
	"jlowell000.github.io/budgeting/internal/model/data"
)

type DataserviceInterface interface {
	GetDataFromFile(string) *data.DataModel
	SaveDataToFile(*data.DataModel, string)
}
