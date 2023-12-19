package monitor

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"picante/internal/test"
	"time"
)

var _ = Describe("Informer replicasets", func() {
	const (
		timeout  = time.Second * 10
		interval = time.Millisecond * 250
	)

	ctx := context.Background()
	BeforeEach(func() {
		//data, err := NewTestData(&client.Project{})
		//Expect(err).Should(Succeed())

		//mocker = mocker.WithTestData(data)
	})

	AfterEach(func() {
	})

	Context("When container does not hava a dependencytrack project", func() {
		It("Should create dependencytrack project with uploaded SBOM", func() {
			Expect(createReplicaSet("default", "pod1", "nginx:latest", timeout, interval)).ShouldNot(BeNil())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "latest")
				return err == nil && p != nil
			}, timeout, interval).Should(BeTrue())
		})
	})
	Context("When container does hava a dependencytrack project", func() {
		It("Should update dependencytrack project with uploaded SBOM", func() {
			Expect(createReplicaSet("default", "pod1", "nginx:evenmorelatest", timeout, interval)).ShouldNot(BeNil())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "evenmorelatest")
				return err == nil && p != nil
			}, timeout, interval).Should(BeTrue())
			ps, err := mocker.GetProjects(ctx)
			Expect(err).Should(Succeed())
			Expect(ps).Should(HaveLen(1))
		})
	})
})

func createReplicaSet(ns, name, image string, timeout, interval time.Duration) *v1.ReplicaSet {
	w := test.Workload(ns, name, nil, nil, image)
	Eventually(func() bool {
		w, err := k8sClient.AppsV1().ReplicaSets(ns).Create(context.Background(), w, metav1.CreateOptions{})
		return err == nil && w != nil
	}, timeout, interval).Should(BeTrue())
	Eventually(func() bool {

		w, err := k8sClient.AppsV1().ReplicaSets(ns).UpdateStatus(ctx, w, metav1.UpdateOptions{})
		return err == nil && w != nil
	}, timeout, interval).Should(BeTrue())
	return w
}
