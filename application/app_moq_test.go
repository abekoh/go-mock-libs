package application

import (
	"context"
	"testing"

	"github.com/abekoh/go-mock-libs/domain/model/examination"
	"github.com/abekoh/go-mock-libs/domain/model/user"
	examMock "github.com/abekoh/go-mock-libs/moq/domain/model/examination"
	userMock "github.com/abekoh/go-mock-libs/moq/domain/model/user"
	"github.com/google/uuid"
)

func TestUserExamAppService_Get_Moq(t *testing.T) {
	usr := testUser()
	usrID := usr.ID()
	exams := testExams(usrID)

	t.Run("指定したIDでユーザ取得が実行される", func(t *testing.T) {
		userRepo := &userMock.RepositoryMock{
			GetFunc: func(context.Context, uuid.UUID) (user.User, error) {
				return usr, nil
			},
		}
		examRepo := &examMock.RepositoryMock{
			GetAllFunc: func(context.Context, uuid.UUID) (examination.ExaminationList, error) {
				return exams, nil
			},
		}

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: usrID.String()})
	})
}
