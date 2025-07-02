package plan

import (
	"time"
)

type CoordinationAction struct {
	To string
	Action CA
	Outcome CO
	Date time.Time
}