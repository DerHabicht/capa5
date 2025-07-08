package billet

import (
	"github.com/google/uuid"
)

type Facility struct {
	id        uuid.UUID
	name      string
	address   string
	city      string
	state     string
	zip       uint
	buildings []Building
}

func NewFacility(
	id uuid.UUID,
	name string,
	address string,
	city string,
	state string,
	zip uint,
	buildings []Building,
) Facility {
	return Facility{
		id:        id,
		name:      name,
		address:   address,
		city:      city,
		state:     state,
		zip:       zip,
		buildings: buildings,
	}
}

func (f Facility) ID() uuid.UUID {
	return f.id
}

func (f Facility) Name() string {
	return f.name
}

func (f Facility) Address() string {
	return f.address
}

func (f Facility) City() string {
	return f.city
}

func (f Facility) State() string {
	return f.state
}

func (f Facility) Zip() uint {
	return f.zip
}

func (f Facility) Buildings() []Building {
	return f.buildings
}
