package application

import (
	"context"
	"testing"

	"github.com/abekoh/go-mock-libs/domain/model/examination"
	"github.com/abekoh/go-mock-libs/domain/model/user"
)

func TestUserExamAppService_Get_Moq(t *testing.T) {
	usr := testUser()
	usrID := usr.ID()
	exams := testExams(usrID)

	t.Run("指定したIDでユーザ取得が実行される", func(t *testing.T) {
		userRepo := &user.RepositoryMock{}
		examRepo := &examination.RepositoryMock{}

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: usrID.String()})
	})
}
