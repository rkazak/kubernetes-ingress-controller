/*
Copyright 2021 Kong, Inc.

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
	"context"

	v1beta1 "github.com/kong/kubernetes-ingress-controller/apis/configuration/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUDPIngresses implements UDPIngressInterface
type FakeUDPIngresses struct {
	Fake *FakeConfigurationV1beta1
	ns   string
}

var udpingressesResource = schema.GroupVersionResource{Group: "configuration", Version: "v1beta1", Resource: "udpingresses"}

var udpingressesKind = schema.GroupVersionKind{Group: "configuration", Version: "v1beta1", Kind: "UDPIngress"}

// Get takes name of the uDPIngress, and returns the corresponding uDPIngress object, and an error if there is any.
func (c *FakeUDPIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.UDPIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(udpingressesResource, c.ns, name), &v1beta1.UDPIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UDPIngress), err
}

// List takes label and field selectors, and returns the list of UDPIngresses that match those selectors.
func (c *FakeUDPIngresses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.UDPIngressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(udpingressesResource, udpingressesKind, c.ns, opts), &v1beta1.UDPIngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.UDPIngressList{ListMeta: obj.(*v1beta1.UDPIngressList).ListMeta}
	for _, item := range obj.(*v1beta1.UDPIngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested uDPIngresses.
func (c *FakeUDPIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(udpingressesResource, c.ns, opts))

}

// Create takes the representation of a uDPIngress and creates it.  Returns the server's representation of the uDPIngress, and an error, if there is any.
func (c *FakeUDPIngresses) Create(ctx context.Context, uDPIngress *v1beta1.UDPIngress, opts v1.CreateOptions) (result *v1beta1.UDPIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(udpingressesResource, c.ns, uDPIngress), &v1beta1.UDPIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UDPIngress), err
}

// Update takes the representation of a uDPIngress and updates it. Returns the server's representation of the uDPIngress, and an error, if there is any.
func (c *FakeUDPIngresses) Update(ctx context.Context, uDPIngress *v1beta1.UDPIngress, opts v1.UpdateOptions) (result *v1beta1.UDPIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(udpingressesResource, c.ns, uDPIngress), &v1beta1.UDPIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UDPIngress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUDPIngresses) UpdateStatus(ctx context.Context, uDPIngress *v1beta1.UDPIngress, opts v1.UpdateOptions) (*v1beta1.UDPIngress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(udpingressesResource, "status", c.ns, uDPIngress), &v1beta1.UDPIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UDPIngress), err
}

// Delete takes name of the uDPIngress and deletes it. Returns an error if one occurs.
func (c *FakeUDPIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(udpingressesResource, c.ns, name), &v1beta1.UDPIngress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUDPIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(udpingressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.UDPIngressList{})
	return err
}

// Patch applies the patch and returns the patched uDPIngress.
func (c *FakeUDPIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.UDPIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(udpingressesResource, c.ns, name, pt, data, subresources...), &v1beta1.UDPIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UDPIngress), err
}
