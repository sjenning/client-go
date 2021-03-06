// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/sjenning/api/security/v1"
	scheme "github.com/sjenning/client-go/security/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RangeAllocationsGetter has a method to return a RangeAllocationInterface.
// A group's client should implement this interface.
type RangeAllocationsGetter interface {
	RangeAllocations() RangeAllocationInterface
}

// RangeAllocationInterface has methods to work with RangeAllocation resources.
type RangeAllocationInterface interface {
	Create(*v1.RangeAllocation) (*v1.RangeAllocation, error)
	Update(*v1.RangeAllocation) (*v1.RangeAllocation, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.RangeAllocation, error)
	List(opts meta_v1.ListOptions) (*v1.RangeAllocationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.RangeAllocation, err error)
	RangeAllocationExpansion
}

// rangeAllocations implements RangeAllocationInterface
type rangeAllocations struct {
	client rest.Interface
}

// newRangeAllocations returns a RangeAllocations
func newRangeAllocations(c *SecurityV1Client) *rangeAllocations {
	return &rangeAllocations{
		client: c.RESTClient(),
	}
}

// Get takes name of the rangeAllocation, and returns the corresponding rangeAllocation object, and an error if there is any.
func (c *rangeAllocations) Get(name string, options meta_v1.GetOptions) (result *v1.RangeAllocation, err error) {
	result = &v1.RangeAllocation{}
	err = c.client.Get().
		Resource("rangeallocations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RangeAllocations that match those selectors.
func (c *rangeAllocations) List(opts meta_v1.ListOptions) (result *v1.RangeAllocationList, err error) {
	result = &v1.RangeAllocationList{}
	err = c.client.Get().
		Resource("rangeallocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rangeAllocations.
func (c *rangeAllocations) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("rangeallocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a rangeAllocation and creates it.  Returns the server's representation of the rangeAllocation, and an error, if there is any.
func (c *rangeAllocations) Create(rangeAllocation *v1.RangeAllocation) (result *v1.RangeAllocation, err error) {
	result = &v1.RangeAllocation{}
	err = c.client.Post().
		Resource("rangeallocations").
		Body(rangeAllocation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rangeAllocation and updates it. Returns the server's representation of the rangeAllocation, and an error, if there is any.
func (c *rangeAllocations) Update(rangeAllocation *v1.RangeAllocation) (result *v1.RangeAllocation, err error) {
	result = &v1.RangeAllocation{}
	err = c.client.Put().
		Resource("rangeallocations").
		Name(rangeAllocation.Name).
		Body(rangeAllocation).
		Do().
		Into(result)
	return
}

// Delete takes name of the rangeAllocation and deletes it. Returns an error if one occurs.
func (c *rangeAllocations) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("rangeallocations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rangeAllocations) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("rangeallocations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rangeAllocation.
func (c *rangeAllocations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.RangeAllocation, err error) {
	result = &v1.RangeAllocation{}
	err = c.client.Patch(pt).
		Resource("rangeallocations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
