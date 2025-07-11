package plan

import (
	"time"
)

type Plan interface {
	PlanNumber() string
	Title() string
	Effective() time.Time
	Sections() []Section
	Annexes() []Section
}
