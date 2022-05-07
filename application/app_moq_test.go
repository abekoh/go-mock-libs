package application

import (
	"context"
	"testing"

	"github.com/abekoh/go-mock-libs/domain/model/examination"
	"github.com/abekoh/go-mock-libs/domain/model/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserExamAppService_Get_Moq(t *testing.T) {
	usr := testUser()
	usrID := usr.ID()
	exams := testExams(usrID)

	t.Run("指定したIDでユーザ取得が実行される", func(t *testing.T) {
		userRepo := &user.RepositoryMock{
			GetFunc: func(ctx context.Context, id uuid.UUID) (user.User, error) {
				if id != usrID {
					t.Errorf("want = %v, got = %v", usrID, id)
				}
				return usr, nil
			},
		}
		examRepo := &examination.RepositoryMock{}

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: usrID.String()})
	})

	t.Run("指定したIDで試験取得が実行される", func(t *testing.T) {
		userRepo := &user.RepositoryMock{
			GetFunc: func(ctx context.Context, id uuid.UUID) (user.User, error) {
				return usr, nil
			},
		}
		examRepo := &examination.RepositoryMock{
			GetAllFunc: func(ctx context.Context, id uuid.UUID) (examination.ExaminationList, error) {
				if id != usrID {
					t.Errorf("want = %v, got = %v", usrID, id)
				}
				return exams, nil
			},
		}

		target := NewUserExamService(userRepo, examRepo)
		target.Get(context.Background(), UserExamGetRequest{ID: usrID.String()})
	})

	t.Run("レスポンスが正しくマッピングされている", func(t *testing.T) {
		userRepo := &user.RepositoryMock{
			GetFunc: func(ctx context.Context, id uuid.UUID) (user.User, error) {
				return usr, nil
			},
		}
		examRepo := &examination.RepositoryMock{
			GetAllFunc: func(ctx context.Context, id uuid.UUID) (examination.ExaminationList, error) {
				if id != usrID {
					t.Errorf("want = %v, got = %v", usrID, id)
				}
				return exams, nil
			},
		}

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
