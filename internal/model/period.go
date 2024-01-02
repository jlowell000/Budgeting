package model

import (
	"strings"
)

//go:generate stringer -type Period
type Period int

const (
	Unknown Period = iota
	Weekly
	Monthly
	Yearly
	Daily
)

var Periods = [5]Period{
	Unknown,
	Weekly,
	Monthly,
	Yearly,
	Daily,
}

func (p Period) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *Period) UnmarshalText(text []byte) error {
	*p = PeriodFromText(string(text))
	return nil
}

func PeriodFromText(text string) Period {
	switch strings.ToLower(text) {
	default:
		return Unknown
	case "weekly":
		return Weekly
	case "monthly":
		return Monthly
	case "yearly":
		return Yearly
	case "daily":
		return Daily
	}
}
