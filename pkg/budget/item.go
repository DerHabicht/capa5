package budget

type Item struct {
	description  string
	amount       int
	transactions []Entry
}

func NewItem(
	description string,
	amount int,
	transactions []Entry,
) Item {
	return Item{
		description: description,
		amount:      amount, transactions: transactions,
	}
}

func (i Item) Description() string {
	return i.description
}

func (i Item) Amount() int {
	return i.amount
}

func (i Item) Transactions() []Entry {
	return i.transactions
}
