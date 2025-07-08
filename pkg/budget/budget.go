package budget

import (
	"github.com/google/uuid"
)

type Budget struct {
	id         uuid.UUID
	accounts   []Account
	categories []Category
}

func NewBudget(
	id uuid.UUID,
	accounts []Account,
	categories []Category,
) Budget {
	return Budget{
		id:         id,
		accounts:   accounts,
		categories: categories,
	}
}

func (b Budget) ID() uuid.UUID {
	return b.id
}

func (b Budget) Accounts() []Account {
	return b.accounts
}

func (b Budget) Categories() []Category {
	return b.categories
}
