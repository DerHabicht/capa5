package plan

import (
	"time"

	"github.com/ag7if/cap"
	"github.com/ag7if/go-files"
	"github.com/google/uuid"
)

type Plan struct {
	id             uuid.UUID
	planNumber     string
	supplement     string
	incidentNumber string
	incidentType   string
	incidentName   string
	coverSeals     []files.File
	effective      time.Time
	location       string
	incidentStart  time.Time
	incidentEnd    time.Time
	commanderTitle cap.Grade
	commanderName  string
	commanderEmail string
	sponsor string
}
