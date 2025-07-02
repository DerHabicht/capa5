package plan

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type CA int

const (
	Approve CA = iota
	Coord
	Info
	Sign
	Review
	POC
	Log
)

func ParseCA(s string) (CA, error) {
	switch strings.ToLower(s) {
	case "approve":
		return Approve, nil
	case "coord":
		return Coord, nil
	case "info":
		return Info, nil
	case "sign":
		return Sign, nil
	case "review":
		return Review, nil
	case "poc":
		return POC, nil
	case "log":
		return Log, nil
	default:
		return -1, errors.Errorf("invalid coordination action: %s", s)
	}
}

func (ca CA) String() string {
	switch ca {
	case Approve:
		return "APPROVE"
	case Coord:
		return "COORD"
	case Info:
		return "INFO"
	case Sign:
		return "SIGN"
	case Review:
		return "REVIEW"
	case POC:
		return "POC"
	case Log:
		return "LOG"
	default:
		return ""
	}
}

func (ca CA) MarshalJSON() ([]byte, error) {
	return []byte(ca.String()), nil
}

func (ca *CA) UnmarshalJSON(raw []byte) error {
	val, err := ParseCA(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*ca = val
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

	*ca = val
	return nil
}
