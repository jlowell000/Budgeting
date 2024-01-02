package period

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_PERIOD      = Weekly
	TEST_PERIOD_NAME = "\"Weekly\""
)

func TestPeriodToJSON(t *testing.T) {
	expected := TEST_PERIOD_NAME
	actual, err := json.Marshal(TEST_PERIOD)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, string(actual[:]))
}

func TestPeriodFromJSON_data_there(t *testing.T) {
	expected := TEST_PERIOD
	var actual Period
	json.Unmarshal([]byte(TEST_PERIOD_NAME), &actual)
	assert.Equal(t, expected, actual)
}

func TestPeriodFromJSON_no_data_defaults_to_unknown(t *testing.T) {
	expected := Unknown // is 0th therefore default enum
	var actual Period
	json.Unmarshal([]byte(""), &actual)
	assert.Equal(t, expected, actual)
}

func TestPeriodicFlowToJSON(t *testing.T) {
	for _, p := range Periods {
		expected := periodWeekContract(p)
		actual := p.WeeklyAmount()

		assert.Equal(t, expected, actual)
	}
}

func periodWeekContract(p Period) float64 {
	switch p {
	default:
		return 0.0
	case Unknown:
		return 0.0
	case Daily:
		return 1 / 7
	case Weekly:
		return 1
	case Monthly:
		return 52 / 12
	case Yearly:
		return 52
	}
}
