package billet

import (
	"github.com/google/uuid"
)

type Building struct {
	id         uuid.UUID
	identifier string
	rooms      []Room
}

func NewBuilding(
	id uuid.UUID,
	identifier string,
	rooms []Room,
) Building {
	return Building{
		id:         id,
		identifier: identifier,
		rooms:      rooms,
	}
}

func (b Building) SetId(id uuid.UUID) {
	b.id = id
}

func (b Building) SetIdentifier(identifier string) {
	b.identifier = identifier
}

func (b Building) SetRooms(rooms []Room) {
	b.rooms = rooms
}
