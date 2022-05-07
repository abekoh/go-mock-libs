//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../../gomock/domain/model/$GOPACKAGE/$GOFILE
package examination

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context, userId uuid.UUID) (ExaminationList, error)
	Save(ctx context.Context, exam Examination) error
}
