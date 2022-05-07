package application

import (
	"context"
	"testing"

	"github.com/abekoh/go-mock-libs/domain/model/examination"
	"github.com/abekoh/go-mock-libs/domain/model/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserExamAppService_Get(t *testing.T) {
	testUser := testUser()
	testUserID := testUser.ID()
	exams := testExams(testUserID)

	t.Run("指定したIDでユーザ取得が実行される", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := user.NewMockRepository(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), testUserID).Return(testUser, nil)
		examRepo := examination.NewMockRepository(ctrl)
		examRepo.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(exams, nil).AnyTimes()

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: testUserID.String()})
	})

	t.Run("指定したIDで試験取得が実行される", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := user.NewMockRepository(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(testUser, nil).AnyTimes()
		examRepo := examination.NewMockRepository(ctrl)
		examRepo.EXPECT().GetAll(gomock.Any(), testUserID).Return(exams, nil)

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: testUserID.String()})
	})

	t.Run("レスポンスが正しくマッピングされている", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := user.NewMockRepository(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(testUser, nil).AnyTimes()
		examRepo := examination.NewMockRepository(ctrl)
		examRepo.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(exams, nil).AnyTimes()

		target := NewUserExamService(userRepo, examRepo)
		res, err := target.Get(context.Background(), UserExamGetRequest{ID: testUserID.String()})

		assert.Equal(t, UserExamResponse{
			ID:       testUserID.String(),
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
