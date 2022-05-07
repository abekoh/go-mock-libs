//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=repository_gomock.go
//go:generate moq -out=repository_moq.go -pkg=$GOPACKAGE -stub . Repository
package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (User, error)
	Save(ctx context.Context, user User) error
}
