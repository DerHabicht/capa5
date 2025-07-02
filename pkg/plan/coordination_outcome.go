package plan

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type CO int

const (
	Concur CO = iota
	NonConcur
	Comment
)

func ParseCO(s string) (CO, error) {
	switch strings.ToLower(s) {
	case "concur":
		return Concur, nil
	case "nonconcur":
		return NonConcur, nil
	case "comment":
		return Comment, nil
	default:
		return -1, errors.Errorf("invalid coord outcome: %s", s)
	}
}

func (co CO) String() string {
	switch co {
	case Concur:
		return "CONCUR"
	case NonConcur:
		return "NONCONCUR"
	case Comment:
		return "COMMENT"
	default:
		return ""
	}
}

func (co CO) MarshalJSON() ([]byte, error) {
	return []byte(co.String()), nil
}

func (co *CO) UnmarshalJSON(raw []byte) error {
	val, err := ParseCO(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*co = val
	return nil
}

func (co CO) Value() (driver.Value, error) {
	return []byte(co.String()), nil
}

func (co *CO) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseCO(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*co = val
	return nil
}
