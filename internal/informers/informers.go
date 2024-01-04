package informers

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"picante/internal/config"
)

type Config struct {
	infs         []cache.SharedIndexInformer
	eventHandler cache.ResourceEventHandler
	log          *log.Entry
	ls           []config.Label
}

func New(k8sClient *kubernetes.Clientset, log *log.Entry, ls []config.Label, eventHandler cache.ResourceEventHandler) *Config {
	tweakListOpts := informers.WithTweakListOptions(
		func(options *v1.ListOptions) {
			options.LabelSelector = toLabelSelectors(ls)
			options.FieldSelector = "metadata.namespace!=kube-system," +
				"metadata.namespace!=kube-public," +
				"metadata.namespace!=cnrm-system"
		})

	factory := informers.NewSharedInformerFactoryWithOptions(k8sClient, 0, tweakListOpts)
	infs := []cache.SharedIndexInformer{
		factory.Apps().V1().ReplicaSets().Informer(),
		// TODO Exclude jobs as they are not needed for now
		// factory.Batch().V1().Jobs().Informer(),
		factory.Apps().V1().StatefulSets().Informer(),
		factory.Apps().V1().DaemonSets().Informer(),
	}
	return &Config{
		infs:         infs,
		eventHandler: eventHandler,
		log:          log,
		ls:           ls,
	}
}

func (c *Config) RunInformers(ctx context.Context) error {
	for _, informer := range c.infs {
		log.Infof("setting up informer")
		err := informer.SetWatchErrorHandler(cache.DefaultWatchErrorHandler)
		if err != nil {
			return fmt.Errorf("set watch error handler: %w", err)
		}

		log.Info("setting up eventHandler, event handler")
		event, err := informer.AddEventHandler(
			c.eventHandler,
			/*cache.ResourceEventHandlerFuncs{
				AddFunc:    c.eventHandler.OnAdd,
				UpdateFunc: c.eventHandler.OnUpdate,
				DeleteFunc: c.eventHandler.OnDelete,
			}*/)
		if err != nil {
			return fmt.Errorf("add event handler: %w", err)
		}

		go informer.Run(ctx.Done())
		if !cache.WaitForCacheSync(ctx.Done(), informer.HasSynced) {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return fmt.Errorf("timed out waiting for caches to sync")
		}

		log.Infof("informer cache synced: %v", event.HasSynced())
	}
	return nil
}

func toLabelSelectors(labels []config.Label) string {
	var labelSelector string
	if len(labels) == 0 {
		return labelSelector
	}

	for _, label := range labels {
		if len(labelSelector) == 0 {
			labelSelector = fmt.Sprintf("%s=%s", label.Name, label.Value)
			continue
		}
		labelSelector = fmt.Sprintf("%s,%s=%s", labelSelector, label.Name, label.Value)
	}
	return labelSelector
}
