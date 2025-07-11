package plan

import (
	"time"

	"github.com/derhabicht/capa5/pkg/user"
)

type CoordinationAction struct {
	To      *user.User
	Action  CA
	Outcome CO
	Date    time.Time
}
