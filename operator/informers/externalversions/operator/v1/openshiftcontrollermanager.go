// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	operator_v1 "github.com/sjenning/api/operator/v1"
	versioned "github.com/sjenning/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/sjenning/client-go/operator/informers/externalversions/internalinterfaces"
	v1 "github.com/sjenning/client-go/operator/listers/operator/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OpenShiftControllerManagerInformer provides access to a shared informer and lister for
// OpenShiftControllerManagers.
type OpenShiftControllerManagerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.OpenShiftControllerManagerLister
}

type openShiftControllerManagerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOpenShiftControllerManagerInformer constructs a new informer for OpenShiftControllerManager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOpenShiftControllerManagerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOpenShiftControllerManagerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOpenShiftControllerManagerInformer constructs a new informer for OpenShiftControllerManager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOpenShiftControllerManagerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().OpenShiftControllerManagers().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().OpenShiftControllerManagers().Watch(options)
			},
		},
		&operator_v1.OpenShiftControllerManager{},
		resyncPeriod,
		indexers,
	)
}

func (f *openShiftControllerManagerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOpenShiftControllerManagerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *openShiftControllerManagerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operator_v1.OpenShiftControllerManager{}, f.defaultInformer)
}

func (f *openShiftControllerManagerInformer) Lister() v1.OpenShiftControllerManagerLister {
	return v1.NewOpenShiftControllerManagerLister(f.Informer().GetIndexer())
}
