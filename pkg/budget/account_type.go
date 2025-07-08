package budget

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type AT int

const (
	Equity AT = iota
	Asset
	Liability
	Income
	Expense
)

func ParseAT(s string) (AT, error) {
	switch strings.ToLower(s) {
	case "equity":
		return Equity, nil
	case "asset":
		return Asset, nil
	case "liability":
		return Liability, nil
	case "income":
		return Income, nil
	case "expense":
		return Expense, nil
	default:
		return -1, errors.Errorf("invalid account type: %s", s)
	}
}

func (at AT) String() string {
	switch at {
	case Equity:
		return "EQUITY"
	case Asset:
		return "ASSET"
	case Liability:
		return "LIABILITY"
	case Income:
		return "INCOME"
	case Expense:
		return "EXPENSE"
	default:
		return ""
	}
}

func (at AT) MarshalJSON() ([]byte, error) {
	return []byte(at.String()), nil
}

func (at *AT) UnmarshalJSON(raw []byte) error {
	val, err := ParseAT(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*at = val
	return nil
}

func (at AT) Value() (driver.Value, error) {
	return []byte(at.String()), nil
}

func (at *AT) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseAT(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*at = val
	return nil
}
