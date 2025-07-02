package org

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ActivityUnit struct {
	id           uuid.UUID
	name         string
	superiorUnit *ActivityUnit
	assignments  []*Assignment
}

func NewActivityUnit(
	id uuid.UUID,
	name string,
	superiorUnit *ActivityUnit,
	assignments []*Assignment, // NOTE: Use an ordering field when creating the database representation of this to preserve index
) *ActivityUnit {
	return &ActivityUnit{
		id:           id,
		name:         name,
		superiorUnit: superiorUnit,
		assignments:  assignments,
	}
}

func (a *ActivityUnit) ID() uuid.UUID {
	return a.id
}

func (a *ActivityUnit) Name() string {
	return a.name
}

func (a *ActivityUnit) SuperiorUnit() *ActivityUnit {
	return a.superiorUnit
}

/*
// I don't know if I actually need this:
func (a *ActivityUnit) Assignments() []*Assignment {
	return a.assignments
}
*/

func (a *ActivityUnit) Leader() (*Assignment, error) {
	if len(a.assignments) > 0 {
		return a.assignments[0], nil
	}

	return nil, errors.Errorf("no leader assignment set for %s (%s)", a.name, a.id)
}

func (a *ActivityUnit) Assignment(idx int) (*Assignment, error) {
	if idx < len(a.assignments) {
		return a.assignments[idx], nil
	}

	return nil, errors.Errorf("index too big for %s (%s) assignment list", a.name, a.id)
}