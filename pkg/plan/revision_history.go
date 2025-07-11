package plan

import "time"

type RevisionHistoryEntry struct {
	revision uint
	date     time.Time
	authors  []string
	notes    string
}

func NewRevisionHistoryEntry(
	revsion uint,
	date time.Time,
	authors []string,
	notes string,
) RevisionHistoryEntry {
	return RevisionHistoryEntry{
		revision: revision,
		date:    date,
		authors: authors,
		notes:   notes,
	}
}
