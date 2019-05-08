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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/aws-iam-authenticator/pkg/apis/iamauthenticator/v1alpha1"
	scheme "sigs.k8s.io/aws-iam-authenticator/pkg/generated/clientset/versioned/scheme"
)

// IAMIdentityMappingsGetter has a method to return a IAMIdentityMappingInterface.
// A group's client should implement this interface.
type IAMIdentityMappingsGetter interface {
	IAMIdentityMappings() IAMIdentityMappingInterface
}

// IAMIdentityMappingInterface has methods to work with IAMIdentityMapping resources.
type IAMIdentityMappingInterface interface {
	Create(*v1alpha1.IAMIdentityMapping) (*v1alpha1.IAMIdentityMapping, error)
	Update(*v1alpha1.IAMIdentityMapping) (*v1alpha1.IAMIdentityMapping, error)
	UpdateStatus(*v1alpha1.IAMIdentityMapping) (*v1alpha1.IAMIdentityMapping, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IAMIdentityMapping, error)
	List(opts v1.ListOptions) (*v1alpha1.IAMIdentityMappingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IAMIdentityMapping, err error)
	IAMIdentityMappingExpansion
}

// iAMIdentityMappings implements IAMIdentityMappingInterface
type iAMIdentityMappings struct {
	client rest.Interface
}

// newIAMIdentityMappings returns a IAMIdentityMappings
func newIAMIdentityMappings(c *IamauthenticatorV1alpha1Client) *iAMIdentityMappings {
	return &iAMIdentityMappings{
		client: c.RESTClient(),
	}
}

// Get takes name of the iAMIdentityMapping, and returns the corresponding iAMIdentityMapping object, and an error if there is any.
func (c *iAMIdentityMappings) Get(name string, options v1.GetOptions) (result *v1alpha1.IAMIdentityMapping, err error) {
	result = &v1alpha1.IAMIdentityMapping{}
	err = c.client.Get().
		Resource("iamidentitymappings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IAMIdentityMappings that match those selectors.
func (c *iAMIdentityMappings) List(opts v1.ListOptions) (result *v1alpha1.IAMIdentityMappingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IAMIdentityMappingList{}
	err = c.client.Get().
		Resource("iamidentitymappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iAMIdentityMappings.
func (c *iAMIdentityMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("iamidentitymappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iAMIdentityMapping and creates it.  Returns the server's representation of the iAMIdentityMapping, and an error, if there is any.
func (c *iAMIdentityMappings) Create(iAMIdentityMapping *v1alpha1.IAMIdentityMapping) (result *v1alpha1.IAMIdentityMapping, err error) {
	result = &v1alpha1.IAMIdentityMapping{}
	err = c.client.Post().
		Resource("iamidentitymappings").
		Body(iAMIdentityMapping).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iAMIdentityMapping and updates it. Returns the server's representation of the iAMIdentityMapping, and an error, if there is any.
func (c *iAMIdentityMappings) Update(iAMIdentityMapping *v1alpha1.IAMIdentityMapping) (result *v1alpha1.IAMIdentityMapping, err error) {
	result = &v1alpha1.IAMIdentityMapping{}
	err = c.client.Put().
		Resource("iamidentitymappings").
		Name(iAMIdentityMapping.Name).
		Body(iAMIdentityMapping).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iAMIdentityMappings) UpdateStatus(iAMIdentityMapping *v1alpha1.IAMIdentityMapping) (result *v1alpha1.IAMIdentityMapping, err error) {
	result = &v1alpha1.IAMIdentityMapping{}
	err = c.client.Put().
		Resource("iamidentitymappings").
		Name(iAMIdentityMapping.Name).
		SubResource("status").
		Body(iAMIdentityMapping).
		Do().
		Into(result)
	return
}

// Delete takes name of the iAMIdentityMapping and deletes it. Returns an error if one occurs.
func (c *iAMIdentityMappings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("iamidentitymappings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iAMIdentityMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("iamidentitymappings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iAMIdentityMapping.
func (c *iAMIdentityMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IAMIdentityMapping, err error) {
	result = &v1alpha1.IAMIdentityMapping{}
	err = c.client.Patch(pt).
		Resource("iamidentitymappings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
