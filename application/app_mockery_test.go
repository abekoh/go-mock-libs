package application

import (
	"context"
	"testing"

	examMock "github.com/abekoh/go-mock-libs/mocks/domain/model/examination"
	userMock "github.com/abekoh/go-mock-libs/mocks/domain/model/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserExamAppService_Get_Mockery(t *testing.T) {
	usr := testUser()
	usrID := usr.ID()
	exams := testExams(usrID)

	t.Run("指定したIDでユーザ取得が実行される", func(t *testing.T) {
		userRepo := userMock.NewRepository(t)
		userRepo.On("Get", mock.Anything, mock.MatchedBy(func(inp uuid.UUID) bool {
			return inp == usrID
		})).Return(usr, nil).Once()
		examRepo := examMock.NewRepository(t)
		examRepo.On("GetAll", mock.Anything, mock.Anything).Return(exams, nil)

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: usrID.String()})
	})

	t.Run("指定したIDで試験取得が実行される", func(t *testing.T) {
		userRepo := userMock.NewRepository(t)
		userRepo.On("Get", mock.Anything, mock.Anything).Return(usr, nil)
		examRepo := examMock.NewRepository(t)
		examRepo.On("GetAll", mock.Anything, mock.MatchedBy(func(inp uuid.UUID) bool {
			return inp == usrID
		})).Return(exams, nil).Once()

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: usrID.String()})
	})

	t.Run("レスポンスが正しくマッピングされている", func(t *testing.T) {
		userRepo := userMock.NewRepository(t)
		userRepo.On("Get", mock.Anything, mock.Anything).Return(usr, nil)
		examRepo := examMock.NewRepository(t)
		examRepo.On("GetAll", mock.Anything, mock.Anything).Return(exams, nil)

		target := NewUserExamService(userRepo, examRepo)
		res, err := target.Get(context.Background(), UserExamGetRequest{ID: usrID.String()})

		assert.Equal(t, UserExamResponse{
			ID:       usrID.String(),
			FullName: "Taro Yamada",
			Birthday: "1990/12/31",
			Exams: ExamResponseList{
				ExamResponse{
					ID:    exams[0].ID().String(),
					Type:  "English",
					Score: 85,
					Date:  "2022/04/25",
				},
				ExamResponse{
					ID:    exams[1].ID().String(),
					Type:  "Math",
					Score: 53,
					Date:  "2022/05/01",
				},
			},
		}, res)
		assert.Nil(t, err)
	})

}
