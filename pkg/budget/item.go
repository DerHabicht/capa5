package budget

type Item struct {
	description  string
	amount       int
	transactions []Split
}

func NewItem(
	description string,
	amount int,
	transactions []Split,
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

func (i Item) Transactions() []Split {
	return i.transactions
}
