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

// FakeClusterOperators implements ClusterOperatorInterface
type FakeClusterOperators struct {
	Fake *FakeConfigV1
}

var clusteroperatorsResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "clusteroperators"}

var clusteroperatorsKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "ClusterOperator"}

// Get takes name of the clusterOperator, and returns the corresponding clusterOperator object, and an error if there is any.
func (c *FakeClusterOperators) Get(name string, options v1.GetOptions) (result *config_v1.ClusterOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteroperatorsResource, name), &config_v1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.ClusterOperator), err
}

// List takes label and field selectors, and returns the list of ClusterOperators that match those selectors.
func (c *FakeClusterOperators) List(opts v1.ListOptions) (result *config_v1.ClusterOperatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteroperatorsResource, clusteroperatorsKind, opts), &config_v1.ClusterOperatorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &config_v1.ClusterOperatorList{ListMeta: obj.(*config_v1.ClusterOperatorList).ListMeta}
	for _, item := range obj.(*config_v1.ClusterOperatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterOperators.
func (c *FakeClusterOperators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteroperatorsResource, opts))
}

// Create takes the representation of a clusterOperator and creates it.  Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *FakeClusterOperators) Create(clusterOperator *config_v1.ClusterOperator) (result *config_v1.ClusterOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteroperatorsResource, clusterOperator), &config_v1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.ClusterOperator), err
}

// Update takes the representation of a clusterOperator and updates it. Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *FakeClusterOperators) Update(clusterOperator *config_v1.ClusterOperator) (result *config_v1.ClusterOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteroperatorsResource, clusterOperator), &config_v1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.ClusterOperator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterOperators) UpdateStatus(clusterOperator *config_v1.ClusterOperator) (*config_v1.ClusterOperator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusteroperatorsResource, "status", clusterOperator), &config_v1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.ClusterOperator), err
}

// Delete takes name of the clusterOperator and deletes it. Returns an error if one occurs.
func (c *FakeClusterOperators) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusteroperatorsResource, name), &config_v1.ClusterOperator{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterOperators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteroperatorsResource, listOptions)

	_, err := c.Fake.Invokes(action, &config_v1.ClusterOperatorList{})
	return err
}

// Patch applies the patch and returns the patched clusterOperator.
func (c *FakeClusterOperators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *config_v1.ClusterOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteroperatorsResource, name, data, subresources...), &config_v1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.ClusterOperator), err
}
