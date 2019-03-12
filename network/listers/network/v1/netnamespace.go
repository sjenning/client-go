// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/sjenning/api/network/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetNamespaceLister helps list NetNamespaces.
type NetNamespaceLister interface {
	// List lists all NetNamespaces in the indexer.
	List(selector labels.Selector) (ret []*v1.NetNamespace, err error)
	// Get retrieves the NetNamespace from the index for a given name.
	Get(name string) (*v1.NetNamespace, error)
	NetNamespaceListerExpansion
}

// netNamespaceLister implements the NetNamespaceLister interface.
type netNamespaceLister struct {
	indexer cache.Indexer
}

// NewNetNamespaceLister returns a new NetNamespaceLister.
func NewNetNamespaceLister(indexer cache.Indexer) NetNamespaceLister {
	return &netNamespaceLister{indexer: indexer}
}

// List lists all NetNamespaces in the indexer.
func (s *netNamespaceLister) List(selector labels.Selector) (ret []*v1.NetNamespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetNamespace))
	})
	return ret, err
}

// Get retrieves the NetNamespace from the index for a given name.
func (s *netNamespaceLister) Get(name string) (*v1.NetNamespace, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("netnamespace"), name)
	}
	return obj.(*v1.NetNamespace), nil
}
