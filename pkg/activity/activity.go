package activity

import (
	"time"

	"github.com/google/uuid"

	"github.com/derhabicht/capa5/pkg/budget"
)

type Activity struct {
	id       uuid.UUID
	name     string
	start    time.Time
	end      time.Time
	rootUnit Unit
	budget   budget.Budget
}
