package activity

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Unit struct {
	id           uuid.UUID
	name         string
	superiorUnit *Unit
	assignments  []*Assignment
}

func NewActivityUnit(
	id uuid.UUID,
	name string,
	superiorUnit *Unit,
	assignments []*Assignment, // NOTE: Use an ordering field when creating the database representation of this to preserve index
) *Unit {
	return &Unit{
		id:           id,
		name:         name,
		superiorUnit: superiorUnit,
		assignments:  assignments,
	}
}

func (a *Unit) ID() uuid.UUID {
	return a.id
}

func (a *Unit) Name() string {
	return a.name
}

func (a *Unit) SuperiorUnit() *Unit {
	return a.superiorUnit
}

/*
// I don't know if I actually need this:
func (a *Unit) Assignments() []*Assignment {
	return a.assignments
}
*/

func (a *Unit) Leader() (*Assignment, error) {
	if len(a.assignments) > 0 {
		return a.assignments[0], nil
	}

	return nil, errors.Errorf("no leader assignment set for %s (%s)", a.name, a.id)
}

func (a *Unit) Assignment(idx int) (*Assignment, error) {
	if idx < len(a.assignments) {
		return a.assignments[idx], nil
	}

	return nil, errors.Errorf("index too big for %s (%s) assignment list", a.name, a.id)
}
