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

	configurationv1 "github.com/kong/kubernetes-ingress-controller/apis/configuration/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKongClusterPlugins implements KongClusterPluginInterface
type FakeKongClusterPlugins struct {
	Fake *FakeConfigurationV1
}

var kongclusterpluginsResource = schema.GroupVersionResource{Group: "configuration", Version: "v1", Resource: "kongclusterplugins"}

var kongclusterpluginsKind = schema.GroupVersionKind{Group: "configuration", Version: "v1", Kind: "KongClusterPlugin"}

// Get takes name of the kongClusterPlugin, and returns the corresponding kongClusterPlugin object, and an error if there is any.
func (c *FakeKongClusterPlugins) Get(ctx context.Context, name string, options v1.GetOptions) (result *configurationv1.KongClusterPlugin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kongclusterpluginsResource, name), &configurationv1.KongClusterPlugin{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configurationv1.KongClusterPlugin), err
}

// List takes label and field selectors, and returns the list of KongClusterPlugins that match those selectors.
func (c *FakeKongClusterPlugins) List(ctx context.Context, opts v1.ListOptions) (result *configurationv1.KongClusterPluginList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kongclusterpluginsResource, kongclusterpluginsKind, opts), &configurationv1.KongClusterPluginList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configurationv1.KongClusterPluginList{ListMeta: obj.(*configurationv1.KongClusterPluginList).ListMeta}
	for _, item := range obj.(*configurationv1.KongClusterPluginList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kongClusterPlugins.
func (c *FakeKongClusterPlugins) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kongclusterpluginsResource, opts))
}

// Create takes the representation of a kongClusterPlugin and creates it.  Returns the server's representation of the kongClusterPlugin, and an error, if there is any.
func (c *FakeKongClusterPlugins) Create(ctx context.Context, kongClusterPlugin *configurationv1.KongClusterPlugin, opts v1.CreateOptions) (result *configurationv1.KongClusterPlugin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kongclusterpluginsResource, kongClusterPlugin), &configurationv1.KongClusterPlugin{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configurationv1.KongClusterPlugin), err
}

// Update takes the representation of a kongClusterPlugin and updates it. Returns the server's representation of the kongClusterPlugin, and an error, if there is any.
func (c *FakeKongClusterPlugins) Update(ctx context.Context, kongClusterPlugin *configurationv1.KongClusterPlugin, opts v1.UpdateOptions) (result *configurationv1.KongClusterPlugin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kongclusterpluginsResource, kongClusterPlugin), &configurationv1.KongClusterPlugin{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configurationv1.KongClusterPlugin), err
}

// Delete takes name of the kongClusterPlugin and deletes it. Returns an error if one occurs.
func (c *FakeKongClusterPlugins) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(kongclusterpluginsResource, name), &configurationv1.KongClusterPlugin{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKongClusterPlugins) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kongclusterpluginsResource, listOpts)

	_, err := c.Fake.Invokes(action, &configurationv1.KongClusterPluginList{})
	return err
}

// Patch applies the patch and returns the patched kongClusterPlugin.
func (c *FakeKongClusterPlugins) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configurationv1.KongClusterPlugin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kongclusterpluginsResource, name, pt, data, subresources...), &configurationv1.KongClusterPlugin{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configurationv1.KongClusterPlugin), err
}
