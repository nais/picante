package monitor

import (
	"context"
	"github.com/in-toto/in-toto-golang/in_toto"
	"github.com/nais/dependencytrack/pkg/client"
	"net/http"
	"picante/internal/attestation"
	"picante/internal/workload"
)

type Mocker struct {
	client.Client
	testData *TestData
}

type TestData struct {
	Statement *in_toto.CycloneDXStatement
	Projects  []*client.Project
}

func NewTestData(project ...*client.Project) (*TestData, error) {
	t := &TestData{
		Statement: &in_toto.CycloneDXStatement{
			StatementHeader: in_toto.StatementHeader{},
			Predicate:       map[string]any{},
		},
	}

	if len(project) == 0 {
		t.Projects = []*client.Project{}
	}
	return t, nil
}

func NewMocker(data *TestData) (*Mocker, error) {
	c := client.New("na", "na", "na")
	return &Mocker{c, data}, nil
}

func (m Mocker) WithTestData(data *TestData) *Mocker {
	m.testData = data
	return &m
}

var _ attestation.Verifier = &Mocker{}
var _ client.Client = &Mocker{}

func (m Mocker) Verify(_ context.Context, container workload.Container) (*attestation.ImageMetadata, error) {
	return &attestation.ImageMetadata{
		BundleVerified: false,
		Image:          container.Image,
		Statement:      m.testData.Statement,
		ContainerName:  container.Name,
	}, nil
}

func (m Mocker) UpdateProject(ctx context.Context, uuid, name, version, group string, tags []string) (*client.Project, error) {
	updated := make([]*client.Project, 0)
	var u *client.Project
	for _, p := range m.testData.Projects {
		if p.Uuid == uuid {
			u = p
			u.Name = name
			u.Version = version
			u.Group = group
			u.Tags = []client.Tag{}
			for _, tag := range tags {
				u.Tags = append(u.Tags, client.Tag{Name: tag})
			}
			updated = append(updated, u)
		} else {
			updated = append(updated, p)
		}
	}
	return u, nil
}

func (m Mocker) DeleteProject(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (m Mocker) DeleteProjects(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (m Mocker) GetProjects(ctx context.Context) ([]*client.Project, error) {
	return m.testData.Projects, nil
}

func (m Mocker) GetProject(ctx context.Context, name, version string) (*client.Project, error) {
	for _, p := range m.testData.Projects {
		if p.Name == name && p.Version == version {
			return p, nil
		}
	}
	return nil, nil
}

func (m Mocker) GetProjectsByTag(ctx context.Context, tag string) ([]*client.Project, error) {
	for _, p := range m.testData.Projects {
		for _, t := range p.Tags {
			if t.Name == tag {
				return []*client.Project{p}, nil
			}
		}
	}
	return []*client.Project{}, nil
}

func (m Mocker) GetTeam(ctx context.Context, team string) (*client.Team, error) {
	//TODO implement me
	panic("implement me")
}
func (m Mocker) PortfolioRefresh(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (m *Mocker) CreateProject(ctx context.Context, name, version, group string, tags []string) (*client.Project, error) {
	t := []client.Tag{}
	for _, tag := range tags {
		t = append(t, client.Tag{Name: tag})
	}

	p := &client.Project{
		Uuid:       name,
		Classifier: "APPLICATION",
		Group:      group,
		Name:       name,
		Version:    version,
		Publisher:  "Team",
		Tags:       t,
	}
	m.testData.Projects = append(m.testData.Projects, p)
	return p, nil
}

func (m Mocker) UpdateProjectInfo(ctx context.Context, uuid, version, group string, tags []string) error {
	//TODO implement me
	panic("implement me")
}

func (m Mocker) UploadProject(ctx context.Context, name, version, parentUuid string, autoCreate bool, bom []byte) error {
	return nil
}

func (m Mocker) Headers(ctx context.Context) (http.Header, error) {
	//TODO implement me
	panic("implement me")
}
