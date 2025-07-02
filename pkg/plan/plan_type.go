package plan

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type PT int

const (
	STRATPLAN PT = iota
	BPLAN
	CONPLAN
	OPLAN
	OPORD
	FRAGORD
)

func ParsePT(s string) (PT, error) {
	switch strings.ToLower(s) {
	case "strategic plan":
		fallthrough
	case "stratplan":
		return STRATPLAN, nil
	case "base plan":
		fallthrough
	case "bplan":
		return BPLAN, nil
	case "concept plan":
		fallthrough
	case "conplan":
		return CONPLAN, nil
	case "operation plan":
		fallthrough
	case "oplan":
		return OPLAN, nil
	case "operation order":
		fallthrough
	case "opord":
		return OPORD, nil
	case "fragmentary operation order":
		fallthrough
	case "fragord":
		return FRAGORD, nil
	default:
		return -1, errors.Errorf("unrecognized plan type: %s", s)
	}
}

func (pt PT) String() string {
	switch pt {
	case STRATPLAN:
		return "STRATPLAN"
	case BPLAN:
		return "BPLAN"
	case CONPLAN:
		return "CONPLAN"
	case OPLAN:
		return "OPLAN"
	case OPORD:
		return "OPORD"
	case FRAGORD:
		return "FRAGORD"
	default:
		return ""
	}
}

func (pt PT) FullName() string {
	switch pt {
	case STRATPLAN:
		return "Strategic Plan"
	case BPLAN:
		return "Base Plan"
	case CONPLAN:
		return "Concept Plan"
	case OPLAN:
		return "Operation Plan"
	case OPORD:
		return "Operation Order"
	case FRAGORD:
		return "Fragmentary Operation Order"
	default:
		return ""
	}
}

func (pt PT) MarshalJSON() ([]byte, error) {
	return []byte(pt.String()), nil
}

func (pt *PT) UnmarshalJSON(raw []byte) error {
	val, err := ParsePT(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*pt = val
	return nil
}

func (pt PT) Value() (driver.Value, error) {
	return []byte(pt.String()), nil
}

func (pt *PT) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParsePT(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*pt = val
	return nil
}
