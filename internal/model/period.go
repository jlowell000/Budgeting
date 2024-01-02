package model

//go:generate stringer -type Period
type Period int

const (
	Weekly Period = iota
	Monthly
	Yearly
	Daily
)

func (b Period) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}
