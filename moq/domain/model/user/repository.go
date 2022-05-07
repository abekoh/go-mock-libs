// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package user

import (
	context "context"
	"github.com/abekoh/go-mock-libs/domain/model/user"
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
// 			GetFunc: func(ctx context.Context, id uuid.UUID) (user.User, error) {
// 				panic("mock out the Get method")
// 			},
// 			SaveFunc: func(ctx context.Context, userMoqParam user.User) error {
// 				panic("mock out the Save method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id uuid.UUID) (user.User, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, userMoqParam user.User) error

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserMoqParam is the userMoqParam argument value.
			UserMoqParam user.User
		}
	}
	lockGet  sync.RWMutex
	lockSave sync.RWMutex
}

// Get calls GetFunc.
func (mock *RepositoryMock) Get(ctx context.Context, id uuid.UUID) (user.User, error) {
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	if mock.GetFunc == nil {
		var (
			userOut user.User
			errOut  error
		)
		return userOut, errOut
	}
	return mock.GetFunc(ctx, id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedRepository.GetCalls())
func (mock *RepositoryMock) GetCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *RepositoryMock) Save(ctx context.Context, userMoqParam user.User) error {
	callInfo := struct {
		Ctx          context.Context
		UserMoqParam user.User
	}{
		Ctx:          ctx,
		UserMoqParam: userMoqParam,
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
	return mock.SaveFunc(ctx, userMoqParam)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedRepository.SaveCalls())
func (mock *RepositoryMock) SaveCalls() []struct {
	Ctx          context.Context
	UserMoqParam user.User
} {
	var calls []struct {
		Ctx          context.Context
		UserMoqParam user.User
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}
