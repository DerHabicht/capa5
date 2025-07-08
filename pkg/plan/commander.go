package plan

import (
	"github.com/ag7if/cap/v2"
)

type Commander struct {
	title   cap.Grade
	name    string
	email   string
	sponsor string
}

func (c Commander) Title() cap.Grade {
	return c.title
}

func (c Commander) Name() string {
	return c.name
}

func (c Commander) Email() string {
	return c.email
}

func (c Commander) Sponsor() string {
	return c.sponsor
}
