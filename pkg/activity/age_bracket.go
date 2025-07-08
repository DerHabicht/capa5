package activity

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type AgeBracket int

const (
	Adult AgeBracket = iota
	LateTeen
	MidTeen
	EarlyTeen
)

func ParseAgeBracket(s string) (AgeBracket, error) {
	switch strings.ToLower(s) {
	case "21+":
		return Adult, nil
	case "18-20":
		return LateTeen, nil
	case "14-17":
		return MidTeen, nil
	case "12-13":
		return EarlyTeen, nil
	default:
		return -1, errors.Errorf("invalid age bracket: %s", s)
	}
}

func (b AgeBracket) String() string {
	switch b {
	case Adult:
		return "21+"
	case LateTeen:
		return "18-20"
	case MidTeen:
		return "14-17"
	case EarlyTeen:
		return "12-13"
	default:
		return ""
	}
}

func (b AgeBracket) MarshalJSON() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *AgeBracket) UnmarshalJSON(raw []byte) error {
	val, err := ParseAgeBracket(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*b = val
	return nil
}

func (b AgeBracket) Value() (driver.Value, error) {
	return []byte(b.String()), nil
}

func (b *AgeBracket) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseAgeBracket(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*b = val
	return nil
}
