/*
Copyright 2018 The Federation v2 Authors.

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
package internalversion

import (
	federation "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation"
	scheme "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedJobOverridesGetter has a method to return a FederatedJobOverrideInterface.
// A group's client should implement this interface.
type FederatedJobOverridesGetter interface {
	FederatedJobOverrides(namespace string) FederatedJobOverrideInterface
}

// FederatedJobOverrideInterface has methods to work with FederatedJobOverride resources.
type FederatedJobOverrideInterface interface {
	Create(*federation.FederatedJobOverride) (*federation.FederatedJobOverride, error)
	Update(*federation.FederatedJobOverride) (*federation.FederatedJobOverride, error)
	UpdateStatus(*federation.FederatedJobOverride) (*federation.FederatedJobOverride, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*federation.FederatedJobOverride, error)
	List(opts v1.ListOptions) (*federation.FederatedJobOverrideList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *federation.FederatedJobOverride, err error)
	FederatedJobOverrideExpansion
}

// federatedJobOverrides implements FederatedJobOverrideInterface
type federatedJobOverrides struct {
	client rest.Interface
	ns     string
}

// newFederatedJobOverrides returns a FederatedJobOverrides
func newFederatedJobOverrides(c *FederationClient, namespace string) *federatedJobOverrides {
	return &federatedJobOverrides{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedJobOverride, and returns the corresponding federatedJobOverride object, and an error if there is any.
func (c *federatedJobOverrides) Get(name string, options v1.GetOptions) (result *federation.FederatedJobOverride, err error) {
	result = &federation.FederatedJobOverride{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedJobOverrides that match those selectors.
func (c *federatedJobOverrides) List(opts v1.ListOptions) (result *federation.FederatedJobOverrideList, err error) {
	result = &federation.FederatedJobOverrideList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedJobOverrides.
func (c *federatedJobOverrides) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedJobOverride and creates it.  Returns the server's representation of the federatedJobOverride, and an error, if there is any.
func (c *federatedJobOverrides) Create(federatedJobOverride *federation.FederatedJobOverride) (result *federation.FederatedJobOverride, err error) {
	result = &federation.FederatedJobOverride{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		Body(federatedJobOverride).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedJobOverride and updates it. Returns the server's representation of the federatedJobOverride, and an error, if there is any.
func (c *federatedJobOverrides) Update(federatedJobOverride *federation.FederatedJobOverride) (result *federation.FederatedJobOverride, err error) {
	result = &federation.FederatedJobOverride{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		Name(federatedJobOverride.Name).
		Body(federatedJobOverride).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedJobOverrides) UpdateStatus(federatedJobOverride *federation.FederatedJobOverride) (result *federation.FederatedJobOverride, err error) {
	result = &federation.FederatedJobOverride{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		Name(federatedJobOverride.Name).
		SubResource("status").
		Body(federatedJobOverride).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedJobOverride and deletes it. Returns an error if one occurs.
func (c *federatedJobOverrides) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedJobOverrides) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedJobOverride.
func (c *federatedJobOverrides) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *federation.FederatedJobOverride, err error) {
	result = &federation.FederatedJobOverride{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedjoboverrides").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
