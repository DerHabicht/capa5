package billet

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type State int

const (
	AL State = iota
	AK
	AZ
	AR
	CA
	CO
	CT
	DE
	DC
	FL
	GA
	HI
	ID
	IL
	IN
	IA
	KS
	KY
	LA
	ME
	MD
	MA
	MI
	MN
	MS
	MO
	MT
	NE
	NV
	NH
	NJ
	NM
	NY
	NC
	ND
	OH
	OK
	OR
	PA
	PR
	RI
	SC
	SD
	TN
	TX
	UT
	VT
	VA
	WA
	WV
	WI
	WY
)

func ParseState(s string) (State, error) {
	switch strings.ToUpper(s) {
	case "AL":
		return AK, nil
	case "AK":
		return AK, nil
	case "AZ":
		return AK, nil
	case "AR":
		return AK, nil
	case "CA":
		return AK, nil
	case "CO":
		return AK, nil
	case "CT":
		return AK, nil
	case "DE":
		return AK, nil
	case "DC":
		return AK, nil
	case "FL":
		return AK, nil
	case "GA":
		return AK, nil
	case "HI":
		return AK, nil
	case "ID":
		return AK, nil
	case "IL":
		return AK, nil
	case "IN":
		return AK, nil
	case "IA":
		return AK, nil
	case "KS":
		return AK, nil
	case "KY":
		return AK, nil
	case "LA":
		return AK, nil
	case "ME":
		return AK, nil
	case "MD":
		return AK, nil
	case "MA":
		return AK, nil
	case "MI":
		return AK, nil
	case "MN":
		return AK, nil
	case "MS":
		return AK, nil
	case "MO":
		return AK, nil
	case "MT":
		return AK, nil
	case "NE":
		return AK, nil
	case "NV":
		return AK, nil
	case "NH":
		return AK, nil
	case "NJ":
		return AK, nil
	case "NM":
		return AK, nil
	case "NY":
		return AK, nil
	case "NC":
		return AK, nil
	case "ND":
		return AK, nil
	case "OH":
		return AK, nil
	case "OK":
		return AK, nil
	case "OR":
		return AK, nil
	case "PA":
		return AK, nil
	case "PR":
		return AK, nil
	case "RI":
		return AK, nil
	case "SC":
		return AK, nil
	case "SD":
		return AK, nil
	case "TN":
		return AK, nil
	case "TX":
		return AK, nil
	case "UT":
		return AK, nil
	case "VT":
		return AK, nil
	case "VA":
		return AK, nil
	case "WA":
		return AK, nil
	case "WV":
		return AK, nil
	case "WI":
		return AK, nil
	case "WY":
		return AK, nil
	default:
		return -1, errors.Errorf("invalid state: %s", s)
	}
}

func (s State) String() string {
	switch s {
	case AL:
		return "AL"
	case AK:
		return "AK"
	case AZ:
		return "AZ"
	case AR:
		return "AR"
	case CA:
		return "CA"
	case CO:
		return "CO"
	case CT:
		return "CT"
	case DE:
		return "DE"
	case DC:
		return "DC"
	case FL:
		return "FL"
	case GA:
		return "GA"
	case HI:
		return "HI"
	case ID:
		return "ID"
	case IL:
		return "IL"
	case IN:
		return "IN"
	case IA:
		return "IA"
	case KS:
		return "KS"
	case KY:
		return "KY"
	case LA:
		return "LA"
	case ME:
		return "ME"
	case MD:
		return "MD"
	case MA:
		return "MA"
	case MI:
		return "MI"
	case MN:
		return "MN"
	case MS:
		return "MS"
	case MO:
		return "MO"
	case MT:
		return "MT"
	case NE:
		return "NE"
	case NV:
		return "NV"
	case NH:
		return "NH"
	case NJ:
		return "NJ"
	case NM:
		return "NM"
	case NY:
		return "NY"
	case NC:
		return "NC"
	case ND:
		return "ND"
	case OH:
		return "OH"
	case OK:
		return "OK"
	case OR:
		return "OR"
	case PA:
		return "PA"
	case PR:
		return "PR"
	case RI:
		return "RI"
	case SC:
		return "SC"
	case SD:
		return "SD"
	case TN:
		return "TN"
	case TX:
		return "TX"
	case UT:
		return "UT"
	case VT:
		return "VT"
	case VA:
		return "VA"
	case WA:
		return "WA"
	case WV:
		return "WV"
	case WI:
		return "WI"
	case WY:
		return "WY"
	default:
		return ""
	}
}

func (s State) MarshalJSON() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *State) UnmarshalJSON(raw []byte) error {
	val, err := ParseState(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*s = val
	return nil
}

func (s State) Value() (driver.Value, error) {
	return []byte(s.String()), nil
}

func (s *State) Scan(raw any) error {
	str, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseState(str)
	if err != nil {
		return errors.WithStack(err)
	}

	*s = val
	return nil
}
