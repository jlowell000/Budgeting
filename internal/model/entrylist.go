package model

import (
	"encoding/json"
	"log"
	"sort"
)

type EntryList struct {
	Entries []Entry
}

// appends entry to entry list
func (entryList *EntryList) Add(entry Entry) {
	entryList.Entries = append(entryList.Entries, entry)
	sort.Slice(entryList.Entries, func(i, j int) bool {
		return entryList.Entries[i].Timestamp.After(entryList.Entries[j].Timestamp)
	})
}

func (entryList *EntryList) GetLatest() Entry {
	return entryList.Entries[0]
}

// Returns JSON encoding of Entry
func (entryList *EntryList) ToJSON() []byte {
	data, err := json.Marshal(entryList)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Parses an Entry from JSON
func EntryListFromJSON(data []byte) EntryList {
	var entryList EntryList
	json.Unmarshal(data, &entryList)
	return entryList
}
