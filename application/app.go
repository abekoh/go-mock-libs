package application

import (
	"context"

	"github.com/abekoh/go-mock-libs/domain/model"
)

type UserAppService struct {
	userRepository model.UserRepository
}

func NewUserAppService(userRepository model.UserRepository) *UserAppService {
	return &UserAppService{userRepository: userRepository}
}

func (s UserAppService) Get(ctx context.Context, req UserGetRequest) (UserResponse, error) {
	id, err := req.UUID()
	if err != nil {
		return UserResponse{}, err
	}
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return UserResponse{}, err
	}
	return NewUserResponse(user), nil
}

func (s UserAppService) GetAll(ctx context.Context) (UserListResponse, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return UserListResponse{}, err
	}
	return NewUserListResponse(users), nil
}

func (s UserAppService) Add(ctx context.Context, req UserAddRequest) (UserResponse, error) {
	newUser, err := req.NewUser()
	if err != nil {
		return UserResponse{}, err
	}
	if err := s.userRepository.Save(ctx, newUser); err != nil {
		return UserResponse{}, err
	}
	return NewUserResponse(newUser), nil
}
