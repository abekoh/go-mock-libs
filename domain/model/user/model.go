//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../../gomock/domain/model/$GOPACKAGE/$GOFILE
package user

import (
	"context"
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
	birthday types.Birthday
}

func NewUser(name Name, birthday types.Birthday) User {
	return User{
		id:       uuid.New(),
		name:     name,
		birthday: birthday,
	}
}

func NewUserWithID(id uuid.UUID, name Name, birthday types.Birthday) User {
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

func (u User) Birthday() types.Birthday {
	return u.birthday
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
