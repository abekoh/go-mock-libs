// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package examination

import (
	context "context"
	"github.com/abekoh/go-mock-libs/domain/model/examination"
	uuid "github.com/google/uuid"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			GetAllFunc: func(ctx context.Context, userId uuid.UUID) (examination.ExaminationList, error) {
// 				panic("mock out the GetAll method")
// 			},
// 			SaveFunc: func(ctx context.Context, exam examination.Examination) error {
// 				panic("mock out the Save method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context, userId uuid.UUID) (examination.ExaminationList, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, exam examination.Examination) error

	// calls tracks calls to the methods.
	calls struct {
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserId is the userId argument value.
			UserId uuid.UUID
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Exam is the exam argument value.
			Exam examination.Examination
		}
	}
	lockGetAll sync.RWMutex
	lockSave   sync.RWMutex
}

// GetAll calls GetAllFunc.
func (mock *RepositoryMock) GetAll(ctx context.Context, userId uuid.UUID) (examination.ExaminationList, error) {
	callInfo := struct {
		Ctx    context.Context
		UserId uuid.UUID
	}{
		Ctx:    ctx,
		UserId: userId,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	if mock.GetAllFunc == nil {
		var (
			examinationListOut examination.ExaminationList
			errOut             error
		)
		return examinationListOut, errOut
	}
	return mock.GetAllFunc(ctx, userId)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//     len(mockedRepository.GetAllCalls())
func (mock *RepositoryMock) GetAllCalls() []struct {
	Ctx    context.Context
	UserId uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserId uuid.UUID
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *RepositoryMock) Save(ctx context.Context, exam examination.Examination) error {
	callInfo := struct {
		Ctx  context.Context
		Exam examination.Examination
	}{
		Ctx:  ctx,
		Exam: exam,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	if mock.SaveFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.SaveFunc(ctx, exam)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedRepository.SaveCalls())
func (mock *RepositoryMock) SaveCalls() []struct {
	Ctx  context.Context
	Exam examination.Examination
} {
	var calls []struct {
		Ctx  context.Context
		Exam examination.Examination
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}