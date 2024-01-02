package io

import (
	"log"
	"os"
)

// Reads file into memory
func ReadFromFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("No file found adding")
		return []byte{}
	}
	return data
}

// bytes to files
func WriteToFile(data []byte, fileName string) {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
