/*
Copyright The Flagger Authors.

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
	gloov1 "github.com/weaveworks/flagger/pkg/apis/gloo/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUpstreamGroups implements UpstreamGroupInterface
type FakeUpstreamGroups struct {
	Fake *FakeGlooV1
	ns   string
}

var upstreamgroupsResource = schema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "upstreamgroups"}

var upstreamgroupsKind = schema.GroupVersionKind{Group: "gloo.solo.io", Version: "v1", Kind: "UpstreamGroup"}

// Get takes name of the upstreamGroup, and returns the corresponding upstreamGroup object, and an error if there is any.
func (c *FakeUpstreamGroups) Get(name string, options v1.GetOptions) (result *gloov1.UpstreamGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(upstreamgroupsResource, c.ns, name), &gloov1.UpstreamGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloov1.UpstreamGroup), err
}

// List takes label and field selectors, and returns the list of UpstreamGroups that match those selectors.
func (c *FakeUpstreamGroups) List(opts v1.ListOptions) (result *gloov1.UpstreamGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(upstreamgroupsResource, upstreamgroupsKind, c.ns, opts), &gloov1.UpstreamGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gloov1.UpstreamGroupList{ListMeta: obj.(*gloov1.UpstreamGroupList).ListMeta}
	for _, item := range obj.(*gloov1.UpstreamGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested upstreamGroups.
func (c *FakeUpstreamGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(upstreamgroupsResource, c.ns, opts))

}

// Create takes the representation of a upstreamGroup and creates it.  Returns the server's representation of the upstreamGroup, and an error, if there is any.
func (c *FakeUpstreamGroups) Create(upstreamGroup *gloov1.UpstreamGroup) (result *gloov1.UpstreamGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(upstreamgroupsResource, c.ns, upstreamGroup), &gloov1.UpstreamGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloov1.UpstreamGroup), err
}

// Update takes the representation of a upstreamGroup and updates it. Returns the server's representation of the upstreamGroup, and an error, if there is any.
func (c *FakeUpstreamGroups) Update(upstreamGroup *gloov1.UpstreamGroup) (result *gloov1.UpstreamGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(upstreamgroupsResource, c.ns, upstreamGroup), &gloov1.UpstreamGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloov1.UpstreamGroup), err
}

// Delete takes name of the upstreamGroup and deletes it. Returns an error if one occurs.
func (c *FakeUpstreamGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(upstreamgroupsResource, c.ns, name), &gloov1.UpstreamGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUpstreamGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(upstreamgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &gloov1.UpstreamGroupList{})
	return err
}

// Patch applies the patch and returns the patched upstreamGroup.
func (c *FakeUpstreamGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *gloov1.UpstreamGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(upstreamgroupsResource, c.ns, name, pt, data, subresources...), &gloov1.UpstreamGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloov1.UpstreamGroup), err
}
