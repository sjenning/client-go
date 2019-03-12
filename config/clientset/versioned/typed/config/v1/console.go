// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/sjenning/api/config/v1"
	scheme "github.com/sjenning/client-go/config/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConsolesGetter has a method to return a ConsoleInterface.
// A group's client should implement this interface.
type ConsolesGetter interface {
	Consoles() ConsoleInterface
}

// ConsoleInterface has methods to work with Console resources.
type ConsoleInterface interface {
	Create(*v1.Console) (*v1.Console, error)
	Update(*v1.Console) (*v1.Console, error)
	UpdateStatus(*v1.Console) (*v1.Console, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Console, error)
	List(opts meta_v1.ListOptions) (*v1.ConsoleList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Console, err error)
	ConsoleExpansion
}

// consoles implements ConsoleInterface
type consoles struct {
	client rest.Interface
}

// newConsoles returns a Consoles
func newConsoles(c *ConfigV1Client) *consoles {
	return &consoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the console, and returns the corresponding console object, and an error if there is any.
func (c *consoles) Get(name string, options meta_v1.GetOptions) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Get().
		Resource("consoles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Consoles that match those selectors.
func (c *consoles) List(opts meta_v1.ListOptions) (result *v1.ConsoleList, err error) {
	result = &v1.ConsoleList{}
	err = c.client.Get().
		Resource("consoles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested consoles.
func (c *consoles) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("consoles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a console and creates it.  Returns the server's representation of the console, and an error, if there is any.
func (c *consoles) Create(console *v1.Console) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Post().
		Resource("consoles").
		Body(console).
		Do().
		Into(result)
	return
}

// Update takes the representation of a console and updates it. Returns the server's representation of the console, and an error, if there is any.
func (c *consoles) Update(console *v1.Console) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Put().
		Resource("consoles").
		Name(console.Name).
		Body(console).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *consoles) UpdateStatus(console *v1.Console) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Put().
		Resource("consoles").
		Name(console.Name).
		SubResource("status").
		Body(console).
		Do().
		Into(result)
	return
}

// Delete takes name of the console and deletes it. Returns an error if one occurs.
func (c *consoles) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("consoles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *consoles) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("consoles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched console.
func (c *consoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Patch(pt).
		Resource("consoles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
