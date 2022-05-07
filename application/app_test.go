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

func testUser() user.User {
	userID := uuid.New()
	name, _ := user.NewName("Kotaro", "Abe")
	birthday, _ := types.NewDate(1990, 12, 31)
	return user.NewUserWithID(userID, name, birthday)
}

func testExams(userID uuid.UUID) examination.ExaminationList {
	examID1 := uuid.New()
	examDate1, _ := types.NewDate(2022, 4, 25)
	exam1 := examination.NewExaminationWithID(examID1, userID, examination.ExaminationTypeEnglish, examDate1, examination.Score(85))
	examID2 := uuid.New()
	examDate2, _ := types.NewDate(2022, 5, 1)
	exam2 := examination.NewExaminationWithID(examID2, userID, examination.ExaminationTypeMath, examDate2, examination.Score(53))
	return examination.ExaminationList{exam1, exam2}
}

func TestUserExamAppService_Get(t *testing.T) {
	user := testUser()
	userID := user.ID()

	exams := testExams(userID)

	t.Run("指定したIDでユーザ取得が実行される", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := userMock.NewMockRepository(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), userID).Return(user, nil)
		examRepo := examMock.NewMockRepository(ctrl)
		examRepo.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(exams, nil).AnyTimes()

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: userID.String()})
	})

	t.Run("指定したIDで試験取得が実行される", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := userMock.NewMockRepository(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()
		examRepo := examMock.NewMockRepository(ctrl)
		examRepo.EXPECT().GetAll(gomock.Any(), userID).Return(exams, nil)

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: userID.String()})
	})

	t.Run("レスポンスが正しくマッピングされている", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := userMock.NewMockRepository(ctrl)
		userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()
		examRepo := examMock.NewMockRepository(ctrl)
		examRepo.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(exams, nil).AnyTimes()

		target := NewUserExamService(userRepo, examRepo)
		res, err := target.Get(context.Background(), UserExamGetRequest{ID: userID.String()})

		assert.Equal(t, UserExamResponse{
			ID:       userID.String(),
			FullName: "Kotaro Abe",
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
