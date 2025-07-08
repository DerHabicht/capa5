package budget

import (
	"github.com/google/uuid"
)

type Split struct {
	id            uuid.UUID
	fromAccountID uuid.UUID
	toAccountID   uuid.UUID
	transfer      int
}

func NewSplit(
	id uuid.UUID,
	fromAccountID uuid.UUID,
	toAccountID uuid.UUID,
	transfer int,
) Split {
	return Split{
		id:            id,
		fromAccountID: fromAccountID,
		toAccountID:   toAccountID,
		transfer:      transfer,
	}
}

func (s Split) ID() uuid.UUID {
	return s.id
}

func (s Split) FromAccountID() uuid.UUID {
	return s.fromAccountID
}

func (s Split) ToAccountID() uuid.UUID {
	return s.toAccountID
}

func (s Split) Transfer() int {
	return s.transfer
}
