/*
Copyright 2022 The Kubernetes Authors.

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
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/kueue/apis/visibility/v1alpha1"
	visibilityv1alpha1 "sigs.k8s.io/kueue/client-go/applyconfiguration/visibility/v1alpha1"
)

// FakeLocalQueues implements LocalQueueInterface
type FakeLocalQueues struct {
	Fake *FakeVisibilityV1alpha1
	ns   string
}

var localqueuesResource = v1alpha1.SchemeGroupVersion.WithResource("localqueues")

var localqueuesKind = v1alpha1.SchemeGroupVersion.WithKind("LocalQueue")

// Get takes name of the localQueue, and returns the corresponding localQueue object, and an error if there is any.
func (c *FakeLocalQueues) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocalQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(localqueuesResource, c.ns, name), &v1alpha1.LocalQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalQueue), err
}

// List takes label and field selectors, and returns the list of LocalQueues that match those selectors.
func (c *FakeLocalQueues) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocalQueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(localqueuesResource, localqueuesKind, c.ns, opts), &v1alpha1.LocalQueueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LocalQueueList{ListMeta: obj.(*v1alpha1.LocalQueueList).ListMeta}
	for _, item := range obj.(*v1alpha1.LocalQueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested localQueues.
func (c *FakeLocalQueues) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(localqueuesResource, c.ns, opts))

}

// Create takes the representation of a localQueue and creates it.  Returns the server's representation of the localQueue, and an error, if there is any.
func (c *FakeLocalQueues) Create(ctx context.Context, localQueue *v1alpha1.LocalQueue, opts v1.CreateOptions) (result *v1alpha1.LocalQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(localqueuesResource, c.ns, localQueue), &v1alpha1.LocalQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalQueue), err
}

// Update takes the representation of a localQueue and updates it. Returns the server's representation of the localQueue, and an error, if there is any.
func (c *FakeLocalQueues) Update(ctx context.Context, localQueue *v1alpha1.LocalQueue, opts v1.UpdateOptions) (result *v1alpha1.LocalQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(localqueuesResource, c.ns, localQueue), &v1alpha1.LocalQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalQueue), err
}

// Delete takes name of the localQueue and deletes it. Returns an error if one occurs.
func (c *FakeLocalQueues) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(localqueuesResource, c.ns, name, opts), &v1alpha1.LocalQueue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocalQueues) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(localqueuesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LocalQueueList{})
	return err
}

// Patch applies the patch and returns the patched localQueue.
func (c *FakeLocalQueues) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(localqueuesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LocalQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalQueue), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied localQueue.
func (c *FakeLocalQueues) Apply(ctx context.Context, localQueue *visibilityv1alpha1.LocalQueueApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LocalQueue, err error) {
	if localQueue == nil {
		return nil, fmt.Errorf("localQueue provided to Apply must not be nil")
	}
	data, err := json.Marshal(localQueue)
	if err != nil {
		return nil, err
	}
	name := localQueue.Name
	if name == nil {
		return nil, fmt.Errorf("localQueue.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(localqueuesResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.LocalQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalQueue), err
}

// GetPendingWorkloadsSummary takes name of the localQueue, and returns the corresponding pendingWorkloadsSummary object, and an error if there is any.
func (c *FakeLocalQueues) GetPendingWorkloadsSummary(ctx context.Context, localQueueName string, options v1.GetOptions) (result *v1alpha1.PendingWorkloadsSummary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(localqueuesResource, c.ns, "pendingworkloads", localQueueName), &v1alpha1.PendingWorkloadsSummary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PendingWorkloadsSummary), err
}
