package application

import (
	"context"

	"github.com/abekoh/go-mock-libs/domain/model/examination"
	"github.com/abekoh/go-mock-libs/domain/model/user"
)

type UserExamService struct {
	userRepository user.Repository
	examRepository examination.Repository
}

func NewUserExamService(userRepository user.Repository, examRepository examination.Repository) *UserExamService {
	return &UserExamService{userRepository: userRepository, examRepository: examRepository}
}

func (s UserExamService) Get(ctx context.Context, req UserExamGetRequest) (UserExamResponse, error) {
	id, err := req.UUID()
	if err != nil {
		return UserExamResponse{}, err
	}
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return UserExamResponse{}, err
	}
	exams, err := s.examRepository.GetAll(ctx, user.ID())
	if err != nil {
		return UserExamResponse{}, err
	}
	return NewUserExamResponse(user, exams), nil
}

func (s UserExamService) AddUser(ctx context.Context, req UserAddRequest) error {
	newUser, err := req.NewUser()
	if err != nil {
		return err
	}
	return s.userRepository.Save(ctx, newUser)
}

func (s UserExamService) AddExam(ctx context.Context, req ExamAddRequest) error {
	newExam, err := req.NewExam()
	if err != nil {
		return err
	}
	return s.examRepository.Save(ctx, newExam)
}
