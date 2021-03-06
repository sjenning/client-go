// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	config_v1 "github.com/sjenning/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAPIServers implements APIServerInterface
type FakeAPIServers struct {
	Fake *FakeConfigV1
}

var apiserversResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "apiservers"}

var apiserversKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "APIServer"}

// Get takes name of the aPIServer, and returns the corresponding aPIServer object, and an error if there is any.
func (c *FakeAPIServers) Get(name string, options v1.GetOptions) (result *config_v1.APIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apiserversResource, name), &config_v1.APIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.APIServer), err
}

// List takes label and field selectors, and returns the list of APIServers that match those selectors.
func (c *FakeAPIServers) List(opts v1.ListOptions) (result *config_v1.APIServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apiserversResource, apiserversKind, opts), &config_v1.APIServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &config_v1.APIServerList{ListMeta: obj.(*config_v1.APIServerList).ListMeta}
	for _, item := range obj.(*config_v1.APIServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIServers.
func (c *FakeAPIServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apiserversResource, opts))
}

// Create takes the representation of a aPIServer and creates it.  Returns the server's representation of the aPIServer, and an error, if there is any.
func (c *FakeAPIServers) Create(aPIServer *config_v1.APIServer) (result *config_v1.APIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apiserversResource, aPIServer), &config_v1.APIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.APIServer), err
}

// Update takes the representation of a aPIServer and updates it. Returns the server's representation of the aPIServer, and an error, if there is any.
func (c *FakeAPIServers) Update(aPIServer *config_v1.APIServer) (result *config_v1.APIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apiserversResource, aPIServer), &config_v1.APIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.APIServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIServers) UpdateStatus(aPIServer *config_v1.APIServer) (*config_v1.APIServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apiserversResource, "status", aPIServer), &config_v1.APIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.APIServer), err
}

// Delete takes name of the aPIServer and deletes it. Returns an error if one occurs.
func (c *FakeAPIServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(apiserversResource, name), &config_v1.APIServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apiserversResource, listOptions)

	_, err := c.Fake.Invokes(action, &config_v1.APIServerList{})
	return err
}

// Patch applies the patch and returns the patched aPIServer.
func (c *FakeAPIServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *config_v1.APIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apiserversResource, name, data, subresources...), &config_v1.APIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.APIServer), err
}
