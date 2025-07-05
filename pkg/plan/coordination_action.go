package plan

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type CA int

const (
	APPROVE CA = iota
	COORD
	INFO
	SIGN
	REVIEW
	POC
	LOG
)

func ParseCA(s string) (CA, error) {
	switch strings.ToLower(s) {
	case "approve":
		return APPROVE, nil
	case "coord":
		return COORD, nil
	case "info":
		return INFO, nil
	case "sign":
		return SIGN, nil
	case "review":
		return REVIEW, nil
	case "poc":
		return POC, nil
	case "log":
		return LOG, nil
	default:
		return -1, errors.Errorf("invalid coordination action: %s", s)
	}
}

func (ca CA) String() string {
	switch ca {
	case APPROVE:
		return "APPROVE"
	case COORD:
		return "COORD"
	case INFO:
		return "INFO"
	case SIGN:
		return "SIGN"
	case REVIEW:
		return "REVIEW"
	case POC:
		return "POC"
	case LOG:
		return "LOG"
	}

	return ""
}

func (ca CA) MarshalJSON() ([]byte, error) {
	return []byte(ca.String()), nil
}

func (ca *CA) UnmarshalJSON(raw []byte) error {
	val, err := ParseCA(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	ca = &val
	return nil
}

func (ca CA) Value() (driver.Value, error) {
	return []byte(ca.String()), nil
}

func (ca *CA) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseCA(s)
	if err != nil {
		return errors.WithStack(err)
	}

	ca = &val
	return nil
}
