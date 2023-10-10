/*
Copyright 2022 The Katalyst Authors.

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

	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/tide/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTideNodePools implements TideNodePoolInterface
type FakeTideNodePools struct {
	Fake *FakeTideV1alpha1
}

var tidenodepoolsResource = schema.GroupVersionResource{Group: "tide.katalyst.kubewharf.io", Version: "v1alpha1", Resource: "tidenodepools"}

var tidenodepoolsKind = schema.GroupVersionKind{Group: "tide.katalyst.kubewharf.io", Version: "v1alpha1", Kind: "TideNodePool"}

// Get takes name of the tideNodePool, and returns the corresponding tideNodePool object, and an error if there is any.
func (c *FakeTideNodePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TideNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(tidenodepoolsResource, name), &v1alpha1.TideNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TideNodePool), err
}

// List takes label and field selectors, and returns the list of TideNodePools that match those selectors.
func (c *FakeTideNodePools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TideNodePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(tidenodepoolsResource, tidenodepoolsKind, opts), &v1alpha1.TideNodePoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TideNodePoolList{ListMeta: obj.(*v1alpha1.TideNodePoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.TideNodePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tideNodePools.
func (c *FakeTideNodePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(tidenodepoolsResource, opts))
}

// Create takes the representation of a tideNodePool and creates it.  Returns the server's representation of the tideNodePool, and an error, if there is any.
func (c *FakeTideNodePools) Create(ctx context.Context, tideNodePool *v1alpha1.TideNodePool, opts v1.CreateOptions) (result *v1alpha1.TideNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tidenodepoolsResource, tideNodePool), &v1alpha1.TideNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TideNodePool), err
}

// Update takes the representation of a tideNodePool and updates it. Returns the server's representation of the tideNodePool, and an error, if there is any.
func (c *FakeTideNodePools) Update(ctx context.Context, tideNodePool *v1alpha1.TideNodePool, opts v1.UpdateOptions) (result *v1alpha1.TideNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(tidenodepoolsResource, tideNodePool), &v1alpha1.TideNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TideNodePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTideNodePools) UpdateStatus(ctx context.Context, tideNodePool *v1alpha1.TideNodePool, opts v1.UpdateOptions) (*v1alpha1.TideNodePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(tidenodepoolsResource, "status", tideNodePool), &v1alpha1.TideNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TideNodePool), err
}

// Delete takes name of the tideNodePool and deletes it. Returns an error if one occurs.
func (c *FakeTideNodePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(tidenodepoolsResource, name, opts), &v1alpha1.TideNodePool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTideNodePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(tidenodepoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TideNodePoolList{})
	return err
}

// Patch applies the patch and returns the patched tideNodePool.
func (c *FakeTideNodePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TideNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(tidenodepoolsResource, name, pt, data, subresources...), &v1alpha1.TideNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TideNodePool), err
}
