package model

type Period int

const (
	Weekly Period = iota
	Monthly
	Yearly
	Daily
)
