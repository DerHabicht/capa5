package budget

import (
	"github.com/google/uuid"
)

type Entry struct {
	id        uuid.UUID
	accountID uuid.UUID
	debit     int
	credit    int
}

func NewSplit(
	id uuid.UUID,
	accountID uuid.UUID,
	debit int,
	credit int,
) Entry {
	return Entry{
		id:        id,
		accountID: accountID,
		debit:     debit,
		credit:    credit,
	}
}

func (s Entry) ID() uuid.UUID {
	return s.id
}

func (s Entry) AccountID() uuid.UUID {
	return s.accountID
}

func (s Entry) Debit() int {
	return s.debit
}

func (s Entry) Credit() int {
	return s.credit
}
