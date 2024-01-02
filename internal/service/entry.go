package service

import (
	"time"

	"github.com/google/uuid"
	"jlowell000.github.io/budgeting/internal/io"
	"jlowell000.github.io/budgeting/internal/model"
)

var (
	getEntryListJSON = io.ReadFromFile
	putEntryListJSON = io.WriteToFile

	parseEntryListJSON = model.EntryListFromJSON

	getNewId = uuid.New
	getTime  = time.Now
)

func CreateEntry(content string, filename string) model.Entry {
	list := GetEntryList(filename)
	list.Add(model.Entry{
		Id:        getNewId(),
		Timestamp: getTime(),
		Content:   content,
	})
	SaveEntryListToFile(list, filename)
	return list.GetLatest()
}

func GetLatestEntry(filename string) model.Entry {
	list := GetEntryList(filename)
	return list.GetLatest()
}

func GetEntryList(filename string) model.EntryList {
	data := getEntryListJSON(filename)
	return parseEntryListJSON(data)
}

func SaveEntryListToFile(entryList model.EntryList, filename string) {
	putEntryListJSON(entryList.ToJSON(), filename)
}
