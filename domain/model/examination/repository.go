//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=repository_gomock.go
//go:generate moq -out=repository_moq.go -pkg=$GOPACKAGE -stub . Repository
package examination

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context, userId uuid.UUID) (ExaminationList, error)
	Save(ctx context.Context, exam Examination) error
}
