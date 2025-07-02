package plan

import (
	"time"
)

type Incident struct {
	number       string
	incidentType string
	name         string
	location     string
	start        time.Time
	end          time.Time
}

func NewIncident(
	number string,
	incidentType string,
	name string,
	location string,
	start time.Time,
	end time.Time,

) Incident {
	return Incident{
		number:       number,
		incidentType: incidentType,
		name:         name,
		location:     location,
		start:        start,
		end:          end,
	}
}

func (i Incident) Number() string {
	return i.number
}

func (i Incident) IncidentType() string {
	return i.incidentType
}

func (i Incident) Name() string {
	return i.name
}

func (i Incident) Location() string {
	return i.location
}

func (i Incident) Start() time.Time {
	return i.start
}

func (i Incident) End() time.Time {
	return i.end
}
