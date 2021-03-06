// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/sjenning/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BuildLister helps list Builds.
type BuildLister interface {
	// List lists all Builds in the indexer.
	List(selector labels.Selector) (ret []*v1.Build, err error)
	// Get retrieves the Build from the index for a given name.
	Get(name string) (*v1.Build, error)
	BuildListerExpansion
}

// buildLister implements the BuildLister interface.
type buildLister struct {
	indexer cache.Indexer
}

// NewBuildLister returns a new BuildLister.
func NewBuildLister(indexer cache.Indexer) BuildLister {
	return &buildLister{indexer: indexer}
}

// List lists all Builds in the indexer.
func (s *buildLister) List(selector labels.Selector) (ret []*v1.Build, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Build))
	})
	return ret, err
}

// Get retrieves the Build from the index for a given name.
func (s *buildLister) Get(name string) (*v1.Build, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("build"), name)
	}
	return obj.(*v1.Build), nil
}
