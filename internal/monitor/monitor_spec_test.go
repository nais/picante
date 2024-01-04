package monitor

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"picante/internal/test"
)

var _ = Describe("Informer replicasets", func() {
	const (
		timeout  = time.Second * 20
		interval = time.Millisecond * 300
	)
	ctx := context.Background()
	var testCounter int32 = 0
	BeforeEach(func() {
		err := deleteCollection("default", timeout, interval)
		Expect(err).Should(Succeed())
	})

	AfterEach(func() {
		atomic.AddInt32(&testCounter, 1)
		fmt.Println("Test finished: ", testCounter)
	})

	Context("When replicaset does not hava a dependencytrack project", func() {
		It("Should create dependencytrack project with uploaded SBOM", func() {
			Expect(createReplicaSet("default", "pod1", "nginx:latest", timeout, interval)).ShouldNot(BeNil())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "latest")
				return err == nil && p != nil
			}, timeout, interval).Should(BeTrue())
			ps, err := mocker.GetProjects(ctx)
			Expect(err).Should(Succeed())
			Expect(ps).Should(HaveLen(1))
			Expect(ps[0].Version).Should(Equal("latest"))
		})
	})
	Context("When replicaset does hava a dependencytrack project", func() {
		It("Should update dependencytrack project with uploaded SBOM", func() {
			Expect(createReplicaSet("default", "pod1", "nginx:latest", timeout, interval)).ShouldNot(BeNil())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "latest")
				return err == nil && p != nil
			}, timeout, interval).Should(BeTrue())
			ps, err := mocker.GetProjects(ctx)
			Expect(err).Should(Succeed())
			Expect(ps).Should(HaveLen(1))
			Expect(ps[0].Version).Should(Equal("latest"))
			r := createReplicaSet("default", "pod1", "nginx:evenmorelatest", timeout, interval)
			Expect(r).ShouldNot(BeNil())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "evenmorelatest")
				return err == nil && p != nil && p.Version == "evenmorelatest"
			}, timeout, interval).Should(BeTrue())
			ps, err = mocker.GetProjects(ctx)
			Expect(err).Should(Succeed())
			Expect(ps).Should(HaveLen(1))
		})
	})
	Context("When replicaset is deleted", func() {
		It("Should delete dependencytrack project if the project with version exists", func() {
			r := createReplicaSet("default", "pod1", "nginx:latest", timeout, interval)
			Expect(r).ShouldNot(BeNil())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "latest")
				return err == nil && p != nil
			}, timeout, interval).Should(BeTrue())
			ps, err := mocker.GetProjects(ctx)
			Expect(err).Should(Succeed())
			Expect(ps).Should(HaveLen(1))
			Expect(ps[0].Version).Should(Equal("latest"))

			deleteReplicaset("default", r.Name, timeout, interval)
			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "latest")
				return err == nil && p == nil
			}, timeout, interval).Should(BeTrue())
			ps, err = mocker.GetProjects(ctx)
			Expect(err).Should(Succeed())
			Expect(ps).Should(HaveLen(0))
		})
	})
	Context("When replicaset is deleted", func() {
		It("Should not delete dependencytrack project if replicaset have another version", func() {
			r1 := createReplicaSet("default", "pod1", "nginx:latest", timeout, interval)
			Expect(r1).ShouldNot(BeNil())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "latest")
				return err == nil && p != nil
			}, timeout, interval).Should(BeTrue())
			ps, err := mocker.GetProjects(ctx)
			Expect(err).Should(Succeed())
			Expect(ps).Should(HaveLen(1))
			Expect(ps[0].Version).Should(Equal("latest"))

			r2 := createReplicaSet("default", "pod1", "nginx:evenmorelatest", timeout, interval)
			Expect(r2).ShouldNot(BeNil())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "evenmorelatest")
				return err == nil && p != nil
			}, timeout, interval).Should(BeTrue())
			ps, err = mocker.GetProjects(ctx)
			Expect(err).Should(Succeed())
			Expect(ps).Should(HaveLen(1))

			deleteReplicaset("default", r1.Name, timeout, interval)
			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "evenmorelatest")
				return err == nil && p != nil
			}, timeout, interval).Should(BeTrue())

			Eventually(func() bool {
				p, err := mocker.GetProject(ctx, "ci:default:pod1", "latest")
				return err == nil && p == nil
			}, timeout, interval).Should(BeTrue())

			ps, err = mocker.GetProjects(ctx)
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
	return updateReplicaset(ns, w, timeout, interval)
}

func updateReplicaset(ns string, workload *v1.ReplicaSet, timeout, interval time.Duration) *v1.ReplicaSet {
	Eventually(func() bool {
		w, err := k8sClient.AppsV1().ReplicaSets(ns).UpdateStatus(ctx, workload, metav1.UpdateOptions{})
		return err == nil && w != nil
	}, timeout, interval).Should(BeTrue())
	return workload
}

func deleteCollection(ns string, timeout, interval time.Duration) error {
	Eventually(func() bool {
		err := k8sClient.AppsV1().ReplicaSets(ns).DeleteCollection(context.Background(), metav1.DeleteOptions{}, metav1.ListOptions{})
		return err == nil
	}, timeout, interval).Should(BeTrue())
	return nil
}

func deleteReplicaset(ns, name string, timeout, interval time.Duration) {
	Eventually(func() bool {
		err := k8sClient.AppsV1().ReplicaSets(ns).Delete(context.Background(), name, metav1.DeleteOptions{})
		fmt.Println(err)
		return err == nil
	}, timeout, interval).Should(BeTrue())
}
