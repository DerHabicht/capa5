package org

import (
	"github.com/ag7if/cap"
	"github.com/google/uuid"
)

type Assignment struct {
	id uuid.UUID
	title string
	memberType cap.MemberType
	minGrade cap.Grade
	maxGrade cap.Grade
	assignee *Member
}

func NewAssignment(
	id uuid.UUID,
	title string,
	memberType cap.MemberType,
	minGrade cap.Grade,
	maxGrade cap.Grade,
	assignee *Member,
) *Assignment {
	return &Assignment{
		id: id,
		title: title,
		memberType: memberType,
		minGrade: minGrade,
		maxGrade: maxGrade,
		assignee: assignee,
	}
}

func (a *Assignment) ID() uuid.UUID {
	return a.id
}

func (a *Assignment) Title() string {
	return a.title
}

func (a *Assignment) MemberType() cap.MemberType {
	return a.memberType
}

func (a *Assignment) MinGrade() cap.Grade {
	return a.minGrade
}

func (a *Assignment) MaxGrade() cap.Grade {
	return a.maxGrade
}

func (a *Assignment) Assignee() *Member {
	return a.assignee
}

func (a *Assignment) AssignMember(member *Member) {
	a.assignee = member
}