package user

import (
	"errors"
	"fmt"

	"github.com/abekoh/go-mock-libs/domain/types"
	"github.com/google/uuid"
)

type Name struct {
	first string
	last  string
}

func NewName(first, last string) (Name, error) {
	if len(first) == 0 || len(last) == 0 {
		return Name{}, errors.New("invalid name")
	}
	return Name{first, last}, nil
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

func (n Name) String() string {
	return n.FullName()
}

type User struct {
	id       uuid.UUID
	name     Name
	birthday types.Date
}

func NewUser(name Name, birthday types.Date) User {
	return User{
		id:       uuid.New(),
		name:     name,
		birthday: birthday,
	}
}

func NewUserWithID(id uuid.UUID, name Name, birthday types.Date) User {
	return User{
		id:       id,
		name:     name,
		birthday: birthday,
	}
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Name() Name {
	return u.name
}

func (u User) Birthday() types.Date {
	return u.birthday
}
