package plan

type Section struct {
	title       string
	content     string
	subsections []Section
}

func NewSection(
	title string,
	content string,
	subsections []Section,
) Section {
	return Section{
		title:       title,
		content:     content,
		subsections: subsections,
	}
}
