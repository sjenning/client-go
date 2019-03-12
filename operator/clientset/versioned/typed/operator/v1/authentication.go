// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/sjenning/api/operator/v1"
	scheme "github.com/sjenning/client-go/operator/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AuthenticationsGetter has a method to return a AuthenticationInterface.
// A group's client should implement this interface.
type AuthenticationsGetter interface {
	Authentications() AuthenticationInterface
}

// AuthenticationInterface has methods to work with Authentication resources.
type AuthenticationInterface interface {
	Create(*v1.Authentication) (*v1.Authentication, error)
	Update(*v1.Authentication) (*v1.Authentication, error)
	UpdateStatus(*v1.Authentication) (*v1.Authentication, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Authentication, error)
	List(opts meta_v1.ListOptions) (*v1.AuthenticationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Authentication, err error)
	AuthenticationExpansion
}

// authentications implements AuthenticationInterface
type authentications struct {
	client rest.Interface
}

// newAuthentications returns a Authentications
func newAuthentications(c *OperatorV1Client) *authentications {
	return &authentications{
		client: c.RESTClient(),
	}
}

// Get takes name of the authentication, and returns the corresponding authentication object, and an error if there is any.
func (c *authentications) Get(name string, options meta_v1.GetOptions) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Get().
		Resource("authentications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Authentications that match those selectors.
func (c *authentications) List(opts meta_v1.ListOptions) (result *v1.AuthenticationList, err error) {
	result = &v1.AuthenticationList{}
	err = c.client.Get().
		Resource("authentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authentications.
func (c *authentications) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("authentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a authentication and creates it.  Returns the server's representation of the authentication, and an error, if there is any.
func (c *authentications) Create(authentication *v1.Authentication) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Post().
		Resource("authentications").
		Body(authentication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a authentication and updates it. Returns the server's representation of the authentication, and an error, if there is any.
func (c *authentications) Update(authentication *v1.Authentication) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Put().
		Resource("authentications").
		Name(authentication.Name).
		Body(authentication).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *authentications) UpdateStatus(authentication *v1.Authentication) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Put().
		Resource("authentications").
		Name(authentication.Name).
		SubResource("status").
		Body(authentication).
		Do().
		Into(result)
	return
}

// Delete takes name of the authentication and deletes it. Returns an error if one occurs.
func (c *authentications) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("authentications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authentications) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("authentications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched authentication.
func (c *authentications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Authentication, err error) {
	result = &v1.Authentication{}
	err = c.client.Patch(pt).
		Resource("authentications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
