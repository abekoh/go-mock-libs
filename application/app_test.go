package application

import (
	"context"
	"testing"

	"github.com/abekoh/go-mock-libs/domain/model/examination"
	"github.com/abekoh/go-mock-libs/domain/model/user"
	"github.com/abekoh/go-mock-libs/domain/types"
	examMock "github.com/abekoh/go-mock-libs/gomock/domain/model/examination"
	userMock "github.com/abekoh/go-mock-libs/gomock/domain/model/user"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserAppService_Get(t *testing.T) {
	t.Run("指定したIDで取得が実行される", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := uuid.New()

		userRepo := userMock.NewMockRepository(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), id).Return(user.User{}, nil)
		examRepo := examMock.NewMockRepository(ctrl)
		examRepo.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(examination.ExaminationList{}, nil)

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: id.String()})
	})

	t.Run("レスポンスが正しくマッピングされている", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := uuid.New()

		name, _ := user.NewName("Kotaro", "Abe")
		birthday, _ := types.NewDate(1990, 12, 31)
		user := user.NewUserWithID(id, name, birthday)

		userRepo := userMock.NewMockRepository(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(user, nil)
		examRepo := examMock.NewMockRepository(ctrl)
		examRepo.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(examination.ExaminationList{}, nil)

		target := NewUserExamService(userRepo, examRepo)
		res, err := target.Get(context.Background(), UserExamGetRequest{ID: id.String()})

		assert.Equal(t, UserExamResponse{
			ID:       id.String(),
			FullName: "Kotaro Abe",
			Birthday: "1990/12/31",
			Exams:    ExamResponseList{},
		}, res)
		assert.Nil(t, err)
	})
}
