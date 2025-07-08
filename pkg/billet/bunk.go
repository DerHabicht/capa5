package billet

import (
	"github.com/google/uuid"

	"github.com/derhabicht/capa5/pkg/activity"
)

type Bunk struct {
	id       uuid.UUID
	assignee activity.Member
}

func NewBunk(
	id uuid.UUID,
	assignee activity.Member,
) Bunk {
	return Bunk{
		id:       id,
		assignee: assignee,
	}
}

func (b Bunk) ID() uuid.UUID {
	return b.id
}

func (b Bunk) Assignee() activity.Member {
	return b.assignee
}
