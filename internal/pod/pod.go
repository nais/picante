package pod

import (
	v1 "k8s.io/api/core/v1"
)

const (
	DefaultPredicateType          = "cyclonedx"
	AppK8sIoNameLabelKey          = "app.kubernetes.io/name"
	SalsaKeyRefLabelKey           = "nais.io/salsa-key-ref"
	SalsaKeylessProviderLabelKey  = "nais.io/salsa-keyRef-provider"
	SalsaPredicateLabelKey        = "nais.io/salsa-predicateType"
	TeamLabelKey                  = "team"
	IgnoreTransparencyLogLabelKey = "nais.io/salsa-ignore-transparency-log"
)

type Info struct {
	ContainerImages []string
	Name            string
	Namespace       string
	Verifier        *Verifier
	PodName         string
	Team            string
}

type Verifier struct {
	KeyRef          string
	KeylessProvider string
	PredicateType   string
	IgnoreTLog      string
}

func GetInfo(obj any) (*Info, error) {
	pod := obj.(*v1.Pod)
	labels := pod.GetLabels()

	var c []string
	for _, container := range pod.Spec.Containers {
		c = append(c, container.Image)
	}

	for _, container := range pod.Spec.InitContainers {
		c = append(c, container.Image)
	}

	return &Info{
		ContainerImages: c,
		Name:            labels[AppK8sIoNameLabelKey],
		Namespace:       pod.GetNamespace(),
		PodName:         pod.GetName(),
		Team:            labels[TeamLabelKey],
		Verifier: &Verifier{
			PredicateType:   labels[SalsaPredicateLabelKey],
			KeyRef:          labels[SalsaKeyRefLabelKey],
			KeylessProvider: labels[SalsaKeylessProviderLabelKey],
			IgnoreTLog:      labels[IgnoreTransparencyLogLabelKey],
		},
	}, nil
}
func (p *Info) IgnoreTLog() bool {
	if p.Verifier == nil {
		return false
	}

	if p.Verifier.IgnoreTLog == "true" {
		return true
	}
	return false
}

func (p *Info) GetPredicateType() string {
	if p.Verifier == nil {
		return DefaultPredicateType
	}
	if p.Verifier.PredicateType == "" {
		return DefaultPredicateType
	}
	return p.Verifier.PredicateType
}

func (p *Info) KeylessVerification() bool {
	if p.Verifier == nil {
		return false
	}

	if p.Verifier.KeyRef == "true" {
		return false
	}
	return true
}
