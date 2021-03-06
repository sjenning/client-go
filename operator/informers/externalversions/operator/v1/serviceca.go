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

// ServiceCAInformer provides access to a shared informer and lister for
// ServiceCAs.
type ServiceCAInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServiceCALister
}

type serviceCAInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewServiceCAInformer constructs a new informer for ServiceCA type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceCAInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceCAInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredServiceCAInformer constructs a new informer for ServiceCA type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceCAInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().ServiceCAs().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().ServiceCAs().Watch(options)
			},
		},
		&operator_v1.ServiceCA{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceCAInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceCAInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceCAInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operator_v1.ServiceCA{}, f.defaultInformer)
}

func (f *serviceCAInformer) Lister() v1.ServiceCALister {
	return v1.NewServiceCALister(f.Informer().GetIndexer())
}
