package monitor

import (
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"picante/internal/attestation"
	"picante/internal/config"
	"picante/internal/informers"
	"picante/internal/workload"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"testing"
)

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var k8sCfg *rest.Config
var k8sClient *kubernetes.Clientset
var testEnv *envtest.Environment

var ctx context.Context
var cancel context.CancelFunc

var mocker *Mocker

func TestInformers(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Informer Suite")
}

var _ = BeforeSuite(func() {

	log.SetLevel(log.DebugLevel)
	testLogger := log.WithFields(log.Fields{
		"component": "testSuite",
	})
	ctx, cancel = context.WithCancel(context.TODO())
	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		// The BinaryAssetsDirectory is only required if you want to run the tests directly
		// without call the makefile target test. If not informed it will look for the
		// default path defined in controller-runtime which is /usr/local/kubebuilder/.
		// Note that you must have the required binaries setup under the bin directory to perform
		// the tests directly. When we run make test it will be setup and used automatically.
		/*BinaryAssetsDirectory: filepath.Join("..", "..", "..", "bin", "k8s",
		fmt.Sprintf("1.28.0-%s-%s", runtime.GOOS, runtime.GOARCH)),

		*/
	}

	var err error
	// cfg is defined in this file globally.
	k8sCfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sCfg).NotTo(BeNil())

	k8sClient, err = kubernetes.NewForConfig(k8sCfg)
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	_, err = k8sClient.CoreV1().Namespaces().Create(ctx, &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ns1",
		},
	}, metav1.CreateOptions{})
	Expect(err).NotTo(HaveOccurred())

	cfg, err := setupConfig()
	Expect(err).NotTo(HaveOccurred())

	testData, err := NewTestData()
	Expect(err).NotTo(HaveOccurred())
	mocker, err = NewMocker(testData)
	Expect(err).NotTo(HaveOccurred())
	/*
		verifier, err := attestation.NewVerifier(
			cfg.Cosign.RekorURL,
			cfg.Cosign.LocalImage,
			cfg.Cosign.IgnoreTLog,
			cfg.GitHub.Organizations,
			cfg.GetPreConfiguredIdentities(),
			cfg.Cosign.KeyRef,
		)
		Expect(err).NotTo(HaveOccurred())
	*/
	/*
		c := dtrack.New(cfg.Storage.Api, cfg.Storage.Username, cfg.Storage.Password, dtrack.WithApiKeySource(cfg.Storage.Team))
	*/
	eh := NewEventHandler(ctx, mocker, mocker, cfg.Cluster)
	inf := informers.New(k8sClient, testLogger, cfg.Features.LabelSelectors, eh)

	Expect(inf.RunInformers(ctx)).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	cancel()
	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})

type MockVerifier struct {
}

func (m MockVerifier) Verify(ctx context.Context, container workload.Container) (*attestation.ImageMetadata, error) {
	return &attestation.ImageMetadata{}, nil
}

var _ attestation.Verifier = &MockVerifier{}

func setupConfig() (*config.Config, error) {
	log.Info("-------- setting up configuration -----------")
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	if err := config.Validate([]string{
		config.MetricsAddress,
		config.StorageApi,
		config.StorageUsername,
		config.StoragePassword,
		config.CosignLocalImage,
		config.Identities,
	}); err != nil {
		return cfg, err
	}

	config.Print([]string{
		config.StoragePassword,
		config.StorageUsername,
	})

	log.Info("-------- configuration loaded ----------")
	return cfg, nil
}
