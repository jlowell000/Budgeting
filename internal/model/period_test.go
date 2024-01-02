package model

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_P_PERIOD      = Weekly
	TEST_P_PERIOD_NAME = "\"Weekly\""
)

func TestPeriodToJSON(t *testing.T) {
	expected := TEST_P_PERIOD_NAME
	actual, err := json.Marshal(TEST_P_PERIOD)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, string(actual[:]))
}

func TestPeriodFromJSON_data_there(t *testing.T) {
	expected := TEST_P_PERIOD
	var actual Period
	json.Unmarshal([]byte(TEST_P_PERIOD_NAME), &actual)
	assert.Equal(t, expected, actual)
}

func TestPeriodFromJSON_no_data_defaults_to_weekly(t *testing.T) {
	expected := Weekly // is 0th therefore default enum
	var actual Period
	json.Unmarshal([]byte(""), &actual)
	assert.Equal(t, expected, actual)
}
