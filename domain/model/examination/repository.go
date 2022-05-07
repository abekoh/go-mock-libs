package examination

import "context"

type Repository interface {
	GetAll(ctx context.Context) (ExaminationList, error)
	Save(ctx context.Context, exam Examination) error
}
