package billet

import (
	"github.com/google/uuid"

	"github.com/derhabicht/capa5/pkg/activity"
)

type Room struct {
	id         uuid.UUID
	identifier string
	bunks      []Bunk
	ageBracket activity.AgeBracket
}

func NewRoom(
	id uuid.UUID,
	number string,
	bunks []Bunk,
	ageBracket activity.AgeBracket,
) Room {
	return Room{
		id:         id,
		identifier: number,
		bunks:      bunks,
		ageBracket: ageBracket,
	}
}

func (r Room) ID() uuid.UUID {
	return r.id
}

func (r Room) Identifier() string {
	return r.identifier
}

func (r Room) Bunks() []Bunk {
	return r.bunks
}

func (r Room) AgeBracket() activity.AgeBracket {
	return r.ageBracket
}
