package budget

import (
	"github.com/google/uuid"
	"github.com/martinlindhe/base36"
)

type Account struct {
	id                 uuid.UUID
	accountType        AT
	code               uint64
	name               string
	transactionEntries []Entry
}

func NewAccount(
	id uuid.UUID,
	accountType AT,
	code string,
	name string,
) Account {
	c := base36.Decode(code)

	return Account{
		id:          id,
		accountType: accountType,
		code:        c,
		name:        name,
	}
}

func (a Account) ID() uuid.UUID {
	return a.id
}

func (a Account) AccountType() AT {
	return a.accountType
}

func (a Account) Code() string {
	return base36.Encode(a.code)
}

func (a Account) Name() string {
	return a.name
}

type Accounts []Account

func (a Accounts) Len() int {
	return len(a)
}

func (a Accounts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Accounts) Less(i, j int) bool {
	return a[i].code < a[j].code
}
