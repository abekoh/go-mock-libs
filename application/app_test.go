package application

import (
	"context"
	"testing"

	"github.com/abekoh/go-mock-libs/domain/model"
	mock "github.com/abekoh/go-mock-libs/gomock/model"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserAppService_Get(t *testing.T) {
	t.Run("指定したIDで取得が実行される", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := uuid.New()

		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().Get(gomock.Any(), id).Return(model.User{}, nil)

		target := NewUserAppService(repo)
		target.Get(context.Background(), UserGetRequest{ID: id.String()})
	})

	t.Run("レスポンスが正しくマッピングされている", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := uuid.New()

		name, _ := model.NewName("Kotaro", "Abe")
		birthday, _ := model.NewBirthday(1990, 12, 31)
		user := model.NewUserWithID(id, name, birthday)

		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(user, nil)

		target := NewUserAppService(repo)
		res, err := target.Get(context.Background(), UserGetRequest{ID: id.String()})

		assert.Equal(t, UserResponse{
			ID:       id.String(),
			FullName: "Kotaro Abe",
			Birthday: "1990/12/31",
		}, res)
		assert.Nil(t, err)
	})
}
