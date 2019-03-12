// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/sjenning/api/operator/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubeControllerManagerLister helps list KubeControllerManagers.
type KubeControllerManagerLister interface {
	// List lists all KubeControllerManagers in the indexer.
	List(selector labels.Selector) (ret []*v1.KubeControllerManager, err error)
	// Get retrieves the KubeControllerManager from the index for a given name.
	Get(name string) (*v1.KubeControllerManager, error)
	KubeControllerManagerListerExpansion
}

// kubeControllerManagerLister implements the KubeControllerManagerLister interface.
type kubeControllerManagerLister struct {
	indexer cache.Indexer
}

// NewKubeControllerManagerLister returns a new KubeControllerManagerLister.
func NewKubeControllerManagerLister(indexer cache.Indexer) KubeControllerManagerLister {
	return &kubeControllerManagerLister{indexer: indexer}
}

// List lists all KubeControllerManagers in the indexer.
func (s *kubeControllerManagerLister) List(selector labels.Selector) (ret []*v1.KubeControllerManager, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KubeControllerManager))
	})
	return ret, err
}

// Get retrieves the KubeControllerManager from the index for a given name.
func (s *kubeControllerManagerLister) Get(name string) (*v1.KubeControllerManager, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kubecontrollermanager"), name)
	}
	return obj.(*v1.KubeControllerManager), nil
}
