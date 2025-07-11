package user

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"

	"github.com/derhabicht/capa5/pkg/activity"
)

type User struct {
	id       uuid.UUID
	username string
	email    string
	member   *activity.Member
	roles    map[uuid.UUID]mapset.Set[Role]
}

func NewUser(
	id uuid.UUID,
	username string,
	email string,
	member *activity.Member,
	roles map[uuid.UUID]mapset.Set[Role],
) User {
	return User{
		id:       id,
		username: username,
		email:    email,
		member:   member,
		roles:    roles,
	}
}
