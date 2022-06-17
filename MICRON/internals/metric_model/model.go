package metric_model

import (
	"time"
)

type ID interface{}

type Metrics struct {
	Id 			ID
	Timestamp	time.Time
	Name		string
	Value		int
}