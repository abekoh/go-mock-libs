//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../../gomock/domain/model/$GOPACKAGE/$GOFILE
//go:generate moq -out=../../../moq/domain/model/$GOPACKAGE/$GOFILE -pkg=$GOPACKAGE -stub . Repository
package examination

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context, userId uuid.UUID) (ExaminationList, error)
	Save(ctx context.Context, exam Examination) error
}
