/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/alexellis/inlets-operator/pkg/apis/inletsoperator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTunnels implements TunnelInterface
type FakeTunnels struct {
	Fake *FakeInletsoperatorV1alpha1
	ns   string
}

var tunnelsResource = schema.GroupVersionResource{Group: "inletsoperator.k8s.io", Version: "v1alpha1", Resource: "tunnels"}

var tunnelsKind = schema.GroupVersionKind{Group: "inletsoperator.k8s.io", Version: "v1alpha1", Kind: "Tunnel"}

// Get takes name of the tunnel, and returns the corresponding tunnel object, and an error if there is any.
func (c *FakeTunnels) Get(name string, options v1.GetOptions) (result *v1alpha1.Tunnel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tunnelsResource, c.ns, name), &v1alpha1.Tunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tunnel), err
}

// List takes label and field selectors, and returns the list of Tunnels that match those selectors.
func (c *FakeTunnels) List(opts v1.ListOptions) (result *v1alpha1.TunnelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tunnelsResource, tunnelsKind, c.ns, opts), &v1alpha1.TunnelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TunnelList{ListMeta: obj.(*v1alpha1.TunnelList).ListMeta}
	for _, item := range obj.(*v1alpha1.TunnelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tunnels.
func (c *FakeTunnels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tunnelsResource, c.ns, opts))

}

// Create takes the representation of a tunnel and creates it.  Returns the server's representation of the tunnel, and an error, if there is any.
func (c *FakeTunnels) Create(tunnel *v1alpha1.Tunnel) (result *v1alpha1.Tunnel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tunnelsResource, c.ns, tunnel), &v1alpha1.Tunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tunnel), err
}

// Update takes the representation of a tunnel and updates it. Returns the server's representation of the tunnel, and an error, if there is any.
func (c *FakeTunnels) Update(tunnel *v1alpha1.Tunnel) (result *v1alpha1.Tunnel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tunnelsResource, c.ns, tunnel), &v1alpha1.Tunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tunnel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTunnels) UpdateStatus(tunnel *v1alpha1.Tunnel) (*v1alpha1.Tunnel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tunnelsResource, "status", c.ns, tunnel), &v1alpha1.Tunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tunnel), err
}

// Delete takes name of the tunnel and deletes it. Returns an error if one occurs.
func (c *FakeTunnels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tunnelsResource, c.ns, name), &v1alpha1.Tunnel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTunnels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tunnelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TunnelList{})
	return err
}

// Patch applies the patch and returns the patched tunnel.
func (c *FakeTunnels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Tunnel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tunnelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Tunnel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tunnel), err
}
