package model

import "time"

type Questionnaire struct {
	ID        int64
	StateDate time.Time
	EndDate   time.Time
}
