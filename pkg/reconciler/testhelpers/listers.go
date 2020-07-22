package testhelpers

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	fakekubeclientset "k8s.io/client-go/kubernetes/fake"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/reconciler/testing"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned/fake"
	v1alpha1Listers "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha1"
	"github.com/pivotal/kpack/pkg/duckbuilder"
)

var clientSetSchemes = []func(*runtime.Scheme) error{
	fake.AddToScheme,
	fakekubeclientset.AddToScheme,
}

type Listers struct {
	sorter testing.ObjectSorter
}

func NewListers(objs []runtime.Object) Listers {
	scheme := runtime.NewScheme()

	for _, addTo := range clientSetSchemes {
		addTo(scheme)
	}

	ls := Listers{
		sorter: testing.NewObjectSorter(scheme),
	}

	ls.sorter.AddObjects(objs...)

	return ls
}

func (l *Listers) indexerFor(obj runtime.Object) cache.Indexer {
	return l.sorter.IndexerForObjectType(obj)
}

func (l *Listers) BuildServiceObjects() []runtime.Object {
	return l.sorter.ObjectsForSchemeFunc(fake.AddToScheme)
}

func (l *Listers) GetKubeObjects() []runtime.Object {
	return l.sorter.ObjectsForSchemeFunc(fakekubeclientset.AddToScheme)
}

func (l *Listers) GetImageLister() v1alpha1Listers.ImageLister {
	return v1alpha1Listers.NewImageLister(l.indexerFor(&v1alpha1.Image{}))
}

func (l *Listers) GetBuildLister() v1alpha1Listers.BuildLister {
	return v1alpha1Listers.NewBuildLister(l.indexerFor(&v1alpha1.Build{}))
}

func (l *Listers) GetCustomBuilderLister() v1alpha1Listers.CustomBuilderLister {
	return v1alpha1Listers.NewCustomBuilderLister(l.indexerFor(&v1alpha1.CustomBuilder{}))
}

func (l *Listers) GetCustomClusterBuilderLister() v1alpha1Listers.CustomClusterBuilderLister {
	return v1alpha1Listers.NewCustomClusterBuilderLister(l.indexerFor(&v1alpha1.CustomClusterBuilder{}))
}

func (l *Listers) GetClusterStoreLister() v1alpha1Listers.ClusterStoreLister {
	return v1alpha1Listers.NewClusterStoreLister(l.indexerFor(&v1alpha1.ClusterStore{}))
}

func (l *Listers) GetClusterStackLister() v1alpha1Listers.ClusterStackLister {
	return v1alpha1Listers.NewClusterStackLister(l.indexerFor(&v1alpha1.ClusterStack{}))
}

func (l *Listers) GetSourceResolverLister() v1alpha1Listers.SourceResolverLister {
	return v1alpha1Listers.NewSourceResolverLister(l.indexerFor(&v1alpha1.SourceResolver{}))
}

func (l *Listers) GetPersistentVolumeClaimLister() corev1listers.PersistentVolumeClaimLister {
	return corev1listers.NewPersistentVolumeClaimLister(l.indexerFor(&corev1.PersistentVolumeClaim{}))
}

func (l *Listers) GetPodLister() corev1listers.PodLister {
	return corev1listers.NewPodLister(l.indexerFor(&corev1.Pod{}))
}

func (l *Listers) GetDuckBuilderLister() *duckbuilder.DuckBuilderLister {
	return &duckbuilder.DuckBuilderLister{
		CustomBuilderLister:        l.GetCustomBuilderLister(),
		CustomClusterBuilderLister: l.GetCustomClusterBuilderLister(),
	}
}
