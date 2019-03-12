// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	operator_v1 "github.com/sjenning/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeAPIServers implements KubeAPIServerInterface
type FakeKubeAPIServers struct {
	Fake *FakeOperatorV1
}

var kubeapiserversResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "kubeapiservers"}

var kubeapiserversKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "KubeAPIServer"}

// Get takes name of the kubeAPIServer, and returns the corresponding kubeAPIServer object, and an error if there is any.
func (c *FakeKubeAPIServers) Get(name string, options v1.GetOptions) (result *operator_v1.KubeAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubeapiserversResource, name), &operator_v1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.KubeAPIServer), err
}

// List takes label and field selectors, and returns the list of KubeAPIServers that match those selectors.
func (c *FakeKubeAPIServers) List(opts v1.ListOptions) (result *operator_v1.KubeAPIServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubeapiserversResource, kubeapiserversKind, opts), &operator_v1.KubeAPIServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operator_v1.KubeAPIServerList{ListMeta: obj.(*operator_v1.KubeAPIServerList).ListMeta}
	for _, item := range obj.(*operator_v1.KubeAPIServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeAPIServers.
func (c *FakeKubeAPIServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubeapiserversResource, opts))
}

// Create takes the representation of a kubeAPIServer and creates it.  Returns the server's representation of the kubeAPIServer, and an error, if there is any.
func (c *FakeKubeAPIServers) Create(kubeAPIServer *operator_v1.KubeAPIServer) (result *operator_v1.KubeAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubeapiserversResource, kubeAPIServer), &operator_v1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.KubeAPIServer), err
}

// Update takes the representation of a kubeAPIServer and updates it. Returns the server's representation of the kubeAPIServer, and an error, if there is any.
func (c *FakeKubeAPIServers) Update(kubeAPIServer *operator_v1.KubeAPIServer) (result *operator_v1.KubeAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubeapiserversResource, kubeAPIServer), &operator_v1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.KubeAPIServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeAPIServers) UpdateStatus(kubeAPIServer *operator_v1.KubeAPIServer) (*operator_v1.KubeAPIServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubeapiserversResource, "status", kubeAPIServer), &operator_v1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.KubeAPIServer), err
}

// Delete takes name of the kubeAPIServer and deletes it. Returns an error if one occurs.
func (c *FakeKubeAPIServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(kubeapiserversResource, name), &operator_v1.KubeAPIServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeAPIServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubeapiserversResource, listOptions)

	_, err := c.Fake.Invokes(action, &operator_v1.KubeAPIServerList{})
	return err
}

// Patch applies the patch and returns the patched kubeAPIServer.
func (c *FakeKubeAPIServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *operator_v1.KubeAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeapiserversResource, name, data, subresources...), &operator_v1.KubeAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operator_v1.KubeAPIServer), err
}
