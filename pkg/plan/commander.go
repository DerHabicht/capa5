package plan

import (
	"github.com/ag7if/cap"
)

type Commander struct {
	commanderTitle cap.Grade
	commanderName  string
	commanderEmail string
	sponsor string
}

func (c Commander) CommanderTitle() cap.Grade {
	return c.commanderTitle
}

func (c Commander) CommanderName() string {
	return c.commanderName
}

func (c Commander) CommanderEmail() string {
	return c.commanderEmail
}

func (c Commander) Sponsor() string {
	return c.sponsor
}