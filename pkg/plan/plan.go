package plan

import (
	"time"

	"github.com/ag7if/go-files"
	"github.com/google/uuid"
)

type Plan struct {
	id           uuid.UUID
	planNumber   string
	supplement   string
	incident     Incident
	commander    Commander
	unit         Unit
	coverSeals   []files.File
	effective    time.Time
	coordination []CoordinationAction
	sections     []PlanSection
}

func NewPlan(
	id uuid.UUID,
	planNumber string,
	supplement string,
	incident Incident,
	commander Commander,
	unit Unit,
	coverSeals []files.File,
	effective time.Time,
	coordination []CoordinationAction,
	sections []PlanSection,
) Plan {
	return Plan{
		id:           id,
		planNumber:   planNumber,
		supplement:   supplement,
		incident:     incident,
		commander:    commander,
		unit:         unit,
		coverSeals:   coverSeals,
		effective:    effective,
		coordination: coordination,
		sections:     sections,
	}
}

func (p Plan) ID() uuid.UUID {
	return p.id
}

func (p Plan) PlanNumber() string {
	return p.planNumber
}

func (p Plan) Supplement() string {
	return p.supplement
}

func (p Plan) Incident() Incident {
	return p.incident
}

func (p Plan) Commander() Commander {
	return p.commander
}

func (p Plan) Unit() Unit {
	return p.unit
}

func (p Plan) CoverSeals() []files.File {
	return p.coverSeals
}

func (p Plan) Effective() time.Time {
	return p.effective
}

func (p Plan) Coordination() []CoordinationAction {
	return p.coordination
}

func (p Plan) Sections() []PlanSection {
	return p.sections
}
