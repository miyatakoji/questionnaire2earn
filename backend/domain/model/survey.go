package model

import "time"

type Survey struct {
	ID        int64
	StateDate time.Time
	EndDate   time.Time
}
