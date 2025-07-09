# Budget Models

Budget objects fall into two categories:

1. planned financial actions consisting of:
    - `Category`
    - `Item`
2. actual financial actions consisting of:
    - `Account`
    - `Transaction`
    - `Entry`

The `Budget` object unifies these into a single object that can be associated
with an `Activity`.

```mermaid
---
title: Budget Classes
---

classDiagram
    Budget *-- Account
    Budget *-- Category
    Transaction *-- Entry
    Account o-- Entry
    Account *-- AT
    Category o-- Item
    Entry -- Item
        
    class Account {
        -id                     uuid.UUID
        -accountType            AT
        -code                   uint64
        -name                   string
        -entries                []*Entry
        +NewAccount()           *Account
        +ID()                   uuid.UUID
        +AccountType()          AT
        +Code()                 string
        +Name()                 string
        +Entries()              []*Entry
        +AddEntry(*Entry)       
    }

    class AT {
        Equity
        Asset
        Liability
        Income
        Expense
        ParseAT(string)         AT, error
        String()                string
        MarshallJSON()          []byte, error
        UnmarshallJSON([]byte)  error
        Value()                 driver.Value, error
        Scan(any)               error
    }

    class Budget {
        -id                     uuid.UUID
        -accounts               []Account
        -categories             []Category
        +NewBudget()            Budget
        +ID()                   uuid.UUID
        +Accounts()             []Account
        +Categories()           []Category
        +TotalPlannedIncome()   int
        +TotalActualIncome()    int
        +TotalPlannedExpense()  int
        +TotalActualExpense()   int
    }

    class Category {
        -id             uuid.UUID
        -name           string
        -account        *Account
        -items          []Item
        +NewCategory()  Category
        +ID()           uuid.UUID
        +Name()         string
        +Account()      *Account
        +Items()        []Item
        +Planned()      int
        +Actual()       int
    }

    class Entry {
        -id         uuid.UUID
        -account    *Account
        -debit      int
        -credit     int
        +NewEntry() *Entry
        +ID()       uuid.UUID
        +Account()  *Account
        +Debit()    int
        +Credit()   int
        +AddEntry(*Entry)
    }

    class Item {
        -description    string
        -debit          int
        -credit         int
        -transactions   []*Entry
        +Description()  string
        +Debit()        int
        +Credit()       int
        +Transactions() []*Entry
        +Planned()      int
        +Actual()       int
    }

    class Transaction {
        -id             uuid.UUID
        -date           time.Time
        -description    string
        -entries        []*Entry
        +ID()           uuid.UUID
        +Date()         time.Time
        +Description()  string
        +Entries()      []*Entry
        +AddEntry(*Entry)
    }
```

## Planned Financial Actions Classes

### Category

### Item

## Actual Financial Actions Classes

### Account

### Transaction

### Entry


```mermaid
---
title: Budget ERD
---

erDiagram
BUDGET ||--|{ CATEGORY: "comprises"
BUDGET ||--|{ ACCOUNT: "comprises"
TRANSACTION ||--|{ ENTRY: "comprises"
ACCOUNT ||..|| CATEGORY: "associated with"
ACCOUNT ||..o{ ENTRY: "modified by"
ACCOUNT ||..o{ ITEM: "planned with"
CATEGORY ||--|{ ITEM: "comprises"
ITEM |o..o{ ENTRY: "associated with"

ACCOUNT {
uuid    id              PK
uuid    budget_id       FK
AT      accountType
uint64  code
string  name
}

BUDGET {
uuid    id              PK
uuid    activity_id     FK
}

CATEGORY {
uuid    id          PK
uuid    account_id  FK
string  name
}

ENTRY {
uuid    id              PK
uuid    transaction_id  FK
uuid    item_id         FK
uuid    account_id      FK
int     debit
int     credit
}

ITEM {
uuid    id              PK
uuid    category_id     FK
string  description
int     debit
int     credit
}

TRANSACTION {
uuid    id              PK
string  description
}
```