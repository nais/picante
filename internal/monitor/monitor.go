package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/nais/dependencytrack/pkg/client"
	log "github.com/sirupsen/logrus"
	"picante/internal/workload"

	"picante/internal/attestation"
)

type Config struct {
	Client   client.Client
	Cluster  string
	verifier attestation.Verifier
	logger   *log.Entry
	ctx      context.Context
}

func NewMonitor(ctx context.Context, client client.Client, verifier attestation.Verifier, cluster string) *Config {
	return &Config{
		Client:   client,
		Cluster:  cluster,
		verifier: verifier,
		logger:   log.WithFields(log.Fields{"package": "monitor"}),
		ctx:      ctx,
	}
}

func (c *Config) OnDelete(obj any) {
	c.logger.WithFields(log.Fields{"event": "OnDelete"})

	w := workload.GetMetadata(obj, c.logger)

	if w == nil {
		return
	}
	if w.GetName() == "" {
		c.logger.Warnf("%s:no app name found: %s ", "delete", w.GetKind())
		return
	}
	if w.Active() {
		c.logger.Debugf("%s:%s:%s:%s is active, skipping", "delete", w.GetKind(), w.GetName(), w.GetIdentifier())
		return
	}

	for _, container := range w.GetContainers() {
		if err := c.deleteProject(w, container); err != nil {
			c.logger.Errorf("delete: %v", err)
			continue
		}
	}
}

func (c *Config) OnUpdate(old any, new any) {
	c.logger.WithFields(log.Fields{"event": "update"})

	wNew := workload.GetMetadata(new, c.logger)
	if wNew == nil {
		return
	}

	if !c.validWorkload("update", wNew) {
		return
	}

	if err := c.verifyContainers(c.ctx, wNew); err != nil {
		c.logger.Warnf("update: verify attestation: %v", err)
	}
}

func (c *Config) OnAdd(obj any) {
	c.logger.WithFields(log.Fields{"event": "add"})

	w := workload.GetMetadata(obj, c.logger)
	if w == nil {
		return
	}

	if !c.validWorkload("add", w) {
		return
	}

	if err := c.verifyContainers(c.ctx, w); err != nil {
		c.logger.Warnf("add: verify attestation: %v", err)
	}
}

func (c *Config) verifyContainers(ctx context.Context, w workload.Workload) error {
	for _, container := range w.GetContainers() {
		appName := w.GetName()
		project := workload.ProjectName(w, c.Cluster, container.Name)
		projectVersion := version(container.Image)
		pp, err := c.Client.GetProject(ctx, project, projectVersion)
		if err != nil {
			return err
		}

		// This if, checks if the project exists and if the project has a sbom.
		// it can produce a 409 conflict if the project exists with the version, but does not have a sbom
		// This is ok, because we want to update the project with if the sbom does not exist
		if pp != nil && pp.LastBomImportFormat != "" {
			c.logger.WithFields(log.Fields{
				"project":        project,
				"projectVersion": projectVersion,
				"workload":       w.GetName(),
				"container":      container.Name,
			}).Debug("project exist and has bom, skipping")
			continue
		} else {

			metadata, err := c.verifier.Verify(c.ctx, container)
			if err != nil {
				c.logger.Warnf("verify attestation, skipping: %v", err)
				continue
			}

			p, err := c.retrieveProject(ctx, project, c.Cluster, w.GetNamespace(), appName)
			if err != nil {
				c.logger.Warnf("retrieve project, skipping %v", err)
				continue
			}

			tags := []string{
				project,
				w.GetNamespace(),
				appName,
				metadata.ContainerName,
				metadata.Image,
				c.Cluster,
				projectVersion,
				"digest:" + metadata.Digest,
			}

			if p != nil {

				//if !c.digestHasChanged(metadata, p) {
				//	c.logger.WithFields(log.Fields{
				//		"projectVersion": projectVersion,
				//		"workload":       w.GetName(),
				//		"container":      metadata.ContainerName,
				//		"digest":         metadata.Digest,
				//	}).Info("project exist and has same digest, skipping")
				//	continue
				//}

				c.logger.WithFields(log.Fields{
					"projectVersion": projectVersion,
					"workload":       w.GetName(),
					"container":      metadata.ContainerName,
					"digest":         metadata.Digest,
				}).Info("project exist update version:", p.Version, " to: ", projectVersion)

				_, err := c.Client.UpdateProject(ctx, p.Uuid, project, projectVersion, w.GetNamespace(), tags)
				if err != nil {
					return err
				}

				if err = c.uploadSBOMToProject(ctx, metadata, project, p.Uuid, projectVersion); err != nil {
					return err
				}

			} else {
				c.logger.WithFields(log.Fields{
					"projectVersion": projectVersion,
					"workload":       w.GetName(),
					"container":      metadata.ContainerName,
					"digest":         metadata.Digest,
				}).Info("project does not exist, creating:", project, ":version::", projectVersion)

				createdP, err := c.Client.CreateProject(ctx, project, projectVersion, w.GetNamespace(), tags)
				if err != nil {
					return err
				}

				if err = c.uploadSBOMToProject(ctx, metadata, project, createdP.Uuid, projectVersion); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// TODO
//func (c *Config) digestHasChanged(metadata *attestation.ImageMetadata, p *client.Project) bool {
//	for _, tag := range p.Tags {
//		if strings.Contains(tag.Name, "digest:") {
//			d := strings.Split(tag.Name, ":")[1]
//			if d == metadata.Digest {
//				return false
//			}
//		}
//	}
//	return true
//}

func (c *Config) uploadSBOMToProject(ctx context.Context, metadata *attestation.ImageMetadata, project, parentUuid, projectVersion string) error {
	b, err := json.Marshal(metadata.Statement.Predicate)
	if err != nil {
		return err
	}

	if err = c.Client.UploadProject(ctx, project, projectVersion, parentUuid, false, b); err != nil {
		return err
	}
	return nil
}

func (c *Config) deleteProject(w workload.Workload, container workload.Container) error {
	project := workload.ProjectName(w, c.Cluster, container.Name)
	projectVersion := version(container.Image)
	pr, err := c.Client.GetProject(c.ctx, project, projectVersion)
	if err != nil {
		return fmt.Errorf("delete: get project: %v", err)
	}

	if pr == nil {
		c.logger.Debugf("%s:trying to delete project:%s:%s, project not found", w.GetKind(), project, projectVersion)
		return nil
	}

	if err = c.Client.DeleteProject(c.ctx, pr.Uuid); err != nil {
		return fmt.Errorf("delete project:%s: %v", project, err)
	}

	c.logger.Infof("%s:deleted project:%s:%s", w.GetKind(), project, projectVersion)
	return nil
}

func (c *Config) retrieveProject(ctx context.Context, projectName, env, team, app string) (*client.Project, error) {
	tag := url.QueryEscape(projectName)
	projects, err := c.Client.GetProjectsByTag(ctx, tag)
	if err != nil {
		return nil, fmt.Errorf("getting projects from DependencyTrack: %w", err)
	}

	if len(projects) == 0 {
		return nil, nil
	}
	var p *client.Project
	for _, project := range projects {
		if containsAllTags(project.Tags, env, team, app) && project.Classifier == "APPLICATION" {
			p = project
			break
		}
	}
	return p, nil
}

func (c *Config) validWorkload(event string, w workload.Workload) bool {
	if w == nil {
		return false
	}
	if w.GetName() == "" {
		c.logger.Warnf("%s:no app name found: %s ", event, w.GetKind())
		return false
	}
	if !w.Active() {
		c.logger.Debugf("%s:%s:%s:%s is not active, skipping", event, w.GetKind(), w.GetName(), w.GetIdentifier())
		return false
	}
	return true
}

func version(image string) string {
	if !strings.Contains(image, "@") {
		i := strings.LastIndex(image, ":")
		return image[i+1:]
	}

	return handleImageDigest(image)
}

func handleImageDigest(image string) string {
	// format: <image>@<digest>
	imageArray := strings.Split(image, "@")
	i := strings.LastIndex(imageArray[0], ":")
	// format: <image>:<tag>@<digest>
	if i != -1 {
		return imageArray[0][i+1:] + "@" + imageArray[1]
	}
	return imageArray[1]
}

func containsAllTags(tags []client.Tag, s ...string) bool {
	found := 0
	for _, t := range s {
		for _, tag := range tags {
			if tag.Name == t {
				found += 1
				break
			}
		}
	}
	return found == len(s)
}
