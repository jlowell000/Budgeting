package io

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_VALUE     = "test value"
	TEST_FILE_NAME = "test_file.txt"
)

func TestReadToFile(t *testing.T) {
	os.Remove(TEST_FILE_NAME)
	err := os.WriteFile(TEST_FILE_NAME, []byte(TEST_VALUE), 0644)
	if err != nil {
		log.Fatal(err)
	}

	data := ReadFromFile(TEST_FILE_NAME)
	assert.Equal(t, []byte(TEST_VALUE), data)
	os.Remove(TEST_FILE_NAME)

}

func TestWriteToFile(t *testing.T) {
	os.Remove(TEST_FILE_NAME)
	WriteToFile([]byte(TEST_VALUE), TEST_FILE_NAME)

	data, err := os.ReadFile(TEST_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, []byte(TEST_VALUE), data)
	os.Remove(TEST_FILE_NAME)
}
