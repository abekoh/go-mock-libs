package application

import (
	"context"
	"testing"

	"github.com/abekoh/go-mock-libs/domain/model/user"
	"github.com/abekoh/go-mock-libs/domain/types"
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

		repo := userMock.NewMockUserRepository(ctrl)
		repo.EXPECT().Get(gomock.Any(), id).Return(user.User{}, nil)

		target := NewUserAppService(repo)
		target.Get(context.Background(), UserGetRequest{ID: id.String()})
	})

	t.Run("レスポンスが正しくマッピングされている", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := uuid.New()

		name, _ := user.NewName("Kotaro", "Abe")
		birthday, _ := types.NewDate(1990, 12, 31)
		user := user.NewUserWithID(id, name, birthday)

		repo := userMock.NewMockUserRepository(ctrl)
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
