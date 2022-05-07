//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../../gomock/domain/model/$GOPACKAGE/$GOFILE
package user

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Get(ctx context.Context, id uuid.UUID) (User, error)
	GetAll(ctx context.Context) (UserList, error)
	Save(ctx context.Context, user User) error
}