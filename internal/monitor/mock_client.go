// Code generated by mockery v2.36.1. DO NOT EDIT.

package monitor

import (
	context "context"

	client "github.com/nais/dependencytrack/pkg/client"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// AddToTeam provides a mock function with given fields: ctx, username, uuid
func (_m *MockClient) AddToTeam(ctx context.Context, username string, uuid string) error {
	ret := _m.Called(ctx, username, uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, username, uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeAdminPassword provides a mock function with given fields: ctx, oldPassword, newPassword
func (_m *MockClient) ChangeAdminPassword(ctx context.Context, oldPassword string, newPassword string) error {
	ret := _m.Called(ctx, oldPassword, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, oldPassword, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfigPropertyAggregate provides a mock function with given fields: ctx, properties
func (_m *MockClient) ConfigPropertyAggregate(ctx context.Context, properties []client.ConfigProperty) ([]client.ConfigProperty, error) {
	ret := _m.Called(ctx, properties)

	var r0 []client.ConfigProperty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []client.ConfigProperty) ([]client.ConfigProperty, error)); ok {
		return rf(ctx, properties)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []client.ConfigProperty) []client.ConfigProperty); ok {
		r0 = rf(ctx, properties)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.ConfigProperty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []client.ConfigProperty) error); ok {
		r1 = rf(ctx, properties)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAdminUsers provides a mock function with given fields: ctx, users, teamUuid
func (_m *MockClient) CreateAdminUsers(ctx context.Context, users *client.AdminUsers, teamUuid string) error {
	ret := _m.Called(ctx, users, teamUuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.AdminUsers, string) error); ok {
		r0 = rf(ctx, users, teamUuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateChildProject provides a mock function with given fields: ctx, project, name, version, group, classifier, tags
func (_m *MockClient) CreateChildProject(ctx context.Context, project *client.Project, name string, version string, group string, classifier string, tags []string) (*client.Project, error) {
	ret := _m.Called(ctx, project, name, version, group, classifier, tags)

	var r0 *client.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.Project, string, string, string, string, []string) (*client.Project, error)); ok {
		return rf(ctx, project, name, version, group, classifier, tags)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *client.Project, string, string, string, string, []string) *client.Project); ok {
		r0 = rf(ctx, project, name, version, group, classifier, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *client.Project, string, string, string, string, []string) error); ok {
		r1 = rf(ctx, project, name, version, group, classifier, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateManagedUser provides a mock function with given fields: ctx, username, password
func (_m *MockClient) CreateManagedUser(ctx context.Context, username string, password string) error {
	ret := _m.Called(ctx, username, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, username, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOidcUser provides a mock function with given fields: ctx, email
func (_m *MockClient) CreateOidcUser(ctx context.Context, email string) error {
	ret := _m.Called(ctx, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProject provides a mock function with given fields: ctx, name, version, group, tags
func (_m *MockClient) CreateProject(ctx context.Context, name string, version string, group string, tags []string) (*client.Project, error) {
	ret := _m.Called(ctx, name, version, group, tags)

	var r0 *client.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []string) (*client.Project, error)); ok {
		return rf(ctx, name, version, group, tags)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []string) *client.Project); ok {
		r0 = rf(ctx, name, version, group, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, []string) error); ok {
		r1 = rf(ctx, name, version, group, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTeam provides a mock function with given fields: ctx, teamName, permissions
func (_m *MockClient) CreateTeam(ctx context.Context, teamName string, permissions []client.Permission) (*client.Team, error) {
	ret := _m.Called(ctx, teamName, permissions)

	var r0 *client.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []client.Permission) (*client.Team, error)); ok {
		return rf(ctx, teamName, permissions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []client.Permission) *client.Team); ok {
		r0 = rf(ctx, teamName, permissions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []client.Permission) error); ok {
		r1 = rf(ctx, teamName, permissions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteManagedUser provides a mock function with given fields: ctx, username
func (_m *MockClient) DeleteManagedUser(ctx context.Context, username string) error {
	ret := _m.Called(ctx, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOidcUser provides a mock function with given fields: ctx, username
func (_m *MockClient) DeleteOidcUser(ctx context.Context, username string) error {
	ret := _m.Called(ctx, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProject provides a mock function with given fields: ctx, uuid
func (_m *MockClient) DeleteProject(ctx context.Context, uuid string) error {
	ret := _m.Called(ctx, uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProjects provides a mock function with given fields: ctx, name
func (_m *MockClient) DeleteProjects(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTeam provides a mock function with given fields: ctx, uuid
func (_m *MockClient) DeleteTeam(ctx context.Context, uuid string) error {
	ret := _m.Called(ctx, uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserMembership provides a mock function with given fields: ctx, uuid, username
func (_m *MockClient) DeleteUserMembership(ctx context.Context, uuid string, username string) error {
	ret := _m.Called(ctx, uuid, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, uuid, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateApiKey provides a mock function with given fields: ctx, uuid
func (_m *MockClient) GenerateApiKey(ctx context.Context, uuid string) (string, error) {
	ret := _m.Called(ctx, uuid)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigProperties provides a mock function with given fields: ctx
func (_m *MockClient) GetConfigProperties(ctx context.Context) ([]client.ConfigProperty, error) {
	ret := _m.Called(ctx)

	var r0 []client.ConfigProperty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]client.ConfigProperty, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []client.ConfigProperty); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.ConfigProperty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEcosystems provides a mock function with given fields: ctx
func (_m *MockClient) GetEcosystems(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFindings provides a mock function with given fields: ctx, projectUuid
func (_m *MockClient) GetFindings(ctx context.Context, projectUuid string) ([]*client.Finding, error) {
	ret := _m.Called(ctx, projectUuid)

	var r0 []*client.Finding
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*client.Finding, error)); ok {
		return rf(ctx, projectUuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*client.Finding); ok {
		r0 = rf(ctx, projectUuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*client.Finding)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectUuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOidcUsers provides a mock function with given fields: ctx
func (_m *MockClient) GetOidcUsers(ctx context.Context) ([]client.User, error) {
	ret := _m.Called(ctx)

	var r0 []client.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]client.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []client.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProject provides a mock function with given fields: ctx, name, version
func (_m *MockClient) GetProject(ctx context.Context, name string, version string) (*client.Project, error) {
	ret := _m.Called(ctx, name, version)

	var r0 *client.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*client.Project, error)); ok {
		return rf(ctx, name, version)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *client.Project); ok {
		r0 = rf(ctx, name, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjects provides a mock function with given fields: ctx
func (_m *MockClient) GetProjects(ctx context.Context) ([]*client.Project, error) {
	ret := _m.Called(ctx)

	var r0 []*client.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*client.Project, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*client.Project); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*client.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectsByTag provides a mock function with given fields: ctx, tag
func (_m *MockClient) GetProjectsByTag(ctx context.Context, tag string) ([]*client.Project, error) {
	ret := _m.Called(ctx, tag)

	var r0 []*client.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*client.Project, error)); ok {
		return rf(ctx, tag)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*client.Project); ok {
		r0 = rf(ctx, tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*client.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeam provides a mock function with given fields: ctx, team
func (_m *MockClient) GetTeam(ctx context.Context, team string) (*client.Team, error) {
	ret := _m.Called(ctx, team)

	var r0 *client.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*client.Team, error)); ok {
		return rf(ctx, team)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *client.Team); ok {
		r0 = rf(ctx, team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, team)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeams provides a mock function with given fields: ctx
func (_m *MockClient) GetTeams(ctx context.Context) ([]client.Team, error) {
	ret := _m.Called(ctx)

	var r0 []client.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]client.Team, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []client.Team); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Headers provides a mock function with given fields: ctx
func (_m *MockClient) Headers(ctx context.Context) (http.Header, error) {
	ret := _m.Called(ctx)

	var r0 http.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (http.Header, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) http.Header); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PortfolioRefresh provides a mock function with given fields: ctx
func (_m *MockClient) PortfolioRefresh(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAdminUsers provides a mock function with given fields: ctx, users
func (_m *MockClient) RemoveAdminUsers(ctx context.Context, users *client.AdminUsers) error {
	ret := _m.Called(ctx, users)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.AdminUsers) error); ok {
		r0 = rf(ctx, users)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProject provides a mock function with given fields: ctx, uuid, name, version, group, tags
func (_m *MockClient) UpdateProject(ctx context.Context, uuid string, name string, version string, group string, tags []string) (*client.Project, error) {
	ret := _m.Called(ctx, uuid, name, version, group, tags)

	var r0 *client.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, []string) (*client.Project, error)); ok {
		return rf(ctx, uuid, name, version, group, tags)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, []string) *client.Project); ok {
		r0 = rf(ctx, uuid, name, version, group, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, []string) error); ok {
		r1 = rf(ctx, uuid, name, version, group, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProjectInfo provides a mock function with given fields: ctx, uuid, version, group, tags
func (_m *MockClient) UpdateProjectInfo(ctx context.Context, uuid string, version string, group string, tags []string) error {
	ret := _m.Called(ctx, uuid, version, group, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []string) error); ok {
		r0 = rf(ctx, uuid, version, group, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadProject provides a mock function with given fields: ctx, name, version, bom
func (_m *MockClient) UploadProject(ctx context.Context, name string, version string, bom []byte) error {
	ret := _m.Called(ctx, name, version, bom)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte) error); ok {
		r0 = rf(ctx, name, version, bom)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Version provides a mock function with given fields: ctx
func (_m *MockClient) Version(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
