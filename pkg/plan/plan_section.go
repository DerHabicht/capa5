package plan

import (
	"fmt"
	"strings"
)

type PlanSection struct {
	level       int
	title       string
	content     string
	subsections []PlanSection
}

func NewPlanSection(
	level int,
	title string,
	content string,
) PlanSection {
	return PlanSection{
		level:   level,
		title:   title,
		content: content,
	}
}

func ParsePlanSection(s string) (PlanSection, error) {
	// (#+)\s+(\w+)$\n+(.*)
	panic("ParsePlanSection() not implemented")
}

func (ps PlanSection) Level() int {
	return ps.level
}

func (ps PlanSection) Title() string {
	return ps.title
}

func (ps PlanSection) Content() string {
	return ps.content
}

func (ps PlanSection) Subsections() []PlanSection {
	return ps.subsections
}

func (ps PlanSection) String() string {
	content := ps.content

	for _, s := range ps.subsections {
		content += s.String()
	}

	return fmt.Sprintf("%s %s\n\n%s", strings.Repeat("#", ps.level), ps.title, content)
}
