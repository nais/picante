// Code generated by mockery v2.41.0. DO NOT EDIT.

package attestation

import (
	context "context"
	workload "picante/internal/workload"

	mock "github.com/stretchr/testify/mock"
)

// MockVerifier is an autogenerated mock type for the Verifier type
type MockVerifier struct {
	mock.Mock
}

// Verify provides a mock function with given fields: ctx, container
func (_m *MockVerifier) Verify(ctx context.Context, container workload.Container) (*ImageMetadata, error) {
	ret := _m.Called(ctx, container)

	if len(ret) == 0 {
		panic("no return value specified for Verify")
	}

	var r0 *ImageMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workload.Container) (*ImageMetadata, error)); ok {
		return rf(ctx, container)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workload.Container) *ImageMetadata); ok {
		r0 = rf(ctx, container)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ImageMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workload.Container) error); ok {
		r1 = rf(ctx, container)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockVerifier creates a new instance of MockVerifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVerifier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVerifier {
	mock := &MockVerifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
