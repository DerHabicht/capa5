package org

import (
	"fmt"

	"github.com/ag7if/cap"
)

type Member struct {
	capid     uint
	grade     cap.Grade
	lastName  string
	firstName string
}

func NewMember(
	capid uint,
	grade cap.Grade,
	lastName string,
	firstName string,
) Member {
	return Member{
		capid:     capid,
		grade:     grade,
		lastName:  lastName,
		firstName: firstName,
	}
}

func (m Member) CAPID() uint {
	return m.capid
}

func (m Member) Grade() cap.Grade {
	return m.grade
}

func (m Member) LastName() string {
	return m.lastName
}

func (m Member) FirstName() string {
	return m.firstName
}

func (m Member) String() string {
	return fmt.Sprintf("%s %s %s", m.grade, m.firstName, m.lastName)
}
