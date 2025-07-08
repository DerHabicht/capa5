package budget

import (
	"github.com/google/uuid"
)

type Transaction struct {
	id          uuid.UUID
	description string
	splits      []Split
}

func NewTransaction(
	id uuid.UUID,
	description string,
	splits []Split,
) Transaction {
	return Transaction{
		id:          id,
		description: description,
		splits:      splits,
	}
}

func (t Transaction) Id() uuid.UUID {
	return t.id
}

func (t Transaction) Description() string {
	return t.description
}

func (t Transaction) Splits() []Split {
	return t.splits
}
