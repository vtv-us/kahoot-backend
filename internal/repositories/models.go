package repositories

import (
	"github.com/vtv-us/kahoot-backend/internal/entities"
)

type User struct {
	entities.User
}

type Group struct {
	entities.Group
}

type UserGroup struct {
	entities.UserGroup
}
