package period

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/shopspring/decimal"
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

func Test_MonthlyAmount(t *testing.T) {
	for _, p := range Periods {
		expected := periodMonthContract(p)
		actual := p.MonthlyAmount()
		assert.Equal(t, expected, actual)
	}
}

func TestPeriodStrings(t *testing.T) {

	for i, s := range PeriodStrings {
		expected := i
		actual := int(PeriodFromText(s))

		assert.Equal(t, expected, actual)
	}
}

func periodMonthContract(p Period) decimal.Decimal {
	switch p {
	default:
		return decimal.NewFromInt(1)
	case Unknown:
		return decimal.NewFromInt(1)
	case Daily:
		return decimal.NewFromInt(30)
	case Weekly:
		return decimal.NewFromInt(4)
	case Monthly:
		return decimal.NewFromInt(1)
	case Yearly:
		return decimal.NewFromInt(1).Div(decimal.NewFromInt(12))
	}
}
