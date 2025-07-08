package budget

import (
	"github.com/google/uuid"
)

type Category struct {
	id      uuid.UUID
	name    string
	account Account
	items   []Item
}

func NewCategory(
	id uuid.UUID,
	name string,
	account Account,
	items []Item,
) Category {
	return Category{
		id:      id,
		name:    name,
		account: account,
		items:   items,
	}
}

func (c Category) ID() uuid.UUID {
	return c.id
}

func (c Category) Name() string {
	return c.name
}

func (c Category) Account() Account {
	return c.account
}

func (c Category) Items() []Item {
	return c.items
}
