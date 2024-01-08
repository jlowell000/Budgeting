package period

import (
	"strings"

	"github.com/shopspring/decimal"
)

/*
Enum to define what the period of a periodic flow is.
*/
//go:generate stringer -type Period
type Period int

const (
	Unknown Period = iota
	Weekly
	Monthly
	Yearly
	Daily
)

/*
List of all Periods
*/
var Periods = [5]Period{
	Unknown,
	Weekly,
	Monthly,
	Yearly,
	Daily,
}

var PeriodStrings = [5]string{
	Unknown.String(),
	Weekly.String(),
	Monthly.String(),
	Yearly.String(),
	Daily.String(),
}

/*
Convert Period to []byte for JSON. To fulfil TextMarshaller interface.
*/
func (p Period) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

/*
Convert String to Period. To fulfil TextUnmarshaller interface.
*/
func (p *Period) UnmarshalText(text []byte) error {
	*p = PeriodFromText(string(text))
	return nil
}

/*
Get the Period for a given string.
*/
func PeriodFromText(text string) Period {
	switch strings.ToLower(text) {
	default:
		return Unknown
	case "daily":
		return Daily
	case "weekly":
		return Weekly
	case "monthly":
		return Monthly
	case "yearly":
		return Yearly
	}
}

/*
Get the Period for a given string.
*/
func (p *Period) WeeklyAmount() decimal.Decimal {
	switch *p {
	default:
		return decimal.NewFromInt(1)
	case Unknown:
		return decimal.NewFromInt(1)
	case Daily:
		return decimal.NewFromInt(1).Div(decimal.NewFromInt(7))
	case Weekly:
		return decimal.NewFromInt(1)
	case Monthly:
		return decimal.NewFromInt(52).Div(decimal.NewFromInt(12))
	case Yearly:
		return decimal.NewFromInt(52)
	}
}
