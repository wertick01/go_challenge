package models

import (
	"time"
)

type ID interface{}

type All struct {
	Fl *Filters
	Mtr *Metrics
}

type Filters struct {
	Begin time.Time
	End time.Time
	Duration time.Duration
	Name string
}

type Metrics struct {
	Id 			ID
	Timestamp	time.Time
	Name		string
	Value		int
}

