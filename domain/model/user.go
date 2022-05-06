package model

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	name     Name
	birthday Birthday
}

func NewUser(name Name, birthday Birthday) User {
	return User{
		id:       uuid.New(),
		name:     name,
		birthday: birthday,
	}
}

func NewUserWithID(id uuid.UUID, name Name, birthday Birthday) User {
	return User{
		id:       id,
		name:     name,
		birthday: birthday,
	}
}

type UserList []User

func NewUserList(users ...User) UserList {
	return UserList(users)
}

type UserRepository interface {
	Get(ctx context.Context, id uuid.UUID) (User, error)
	GetAll(ctx context.Context) (UserList, error)
	Save(ctx context.Context, user User) error
}
