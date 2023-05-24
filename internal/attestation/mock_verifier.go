// Code generated by mockery v2.27.1. DO NOT EDIT.

package attestation

import (
	context "context"
	pod "picante/internal/pod"

	mock "github.com/stretchr/testify/mock"
)

// MockVerifier is an autogenerated mock type for the Verifier type
type MockVerifier struct {
	mock.Mock
}

// Verify provides a mock function with given fields: ctx, _a1
func (_m *MockVerifier) Verify(ctx context.Context, _a1 *pod.Info) ([]*ImageMetadata, error) {
	ret := _m.Called(ctx, _a1)

	var r0 []*ImageMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pod.Info) ([]*ImageMetadata, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pod.Info) []*ImageMetadata); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ImageMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pod.Info) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockVerifier interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockVerifier creates a new instance of MockVerifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockVerifier(t mockConstructorTestingTNewMockVerifier) *MockVerifier {
	mock := &MockVerifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
