/*
Copyright AppsCode Inc. and Contributors

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
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-pagerduty-api/apis/user/v1alpha1"
	scheme "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UsersGetter has a method to return a UserInterface.
// A group's client should implement this interface.
type UsersGetter interface {
	Users(namespace string) UserInterface
}

// UserInterface has methods to work with User resources.
type UserInterface interface {
	Create(ctx context.Context, user *v1alpha1.User, opts v1.CreateOptions) (*v1alpha1.User, error)
	Update(ctx context.Context, user *v1alpha1.User, opts v1.UpdateOptions) (*v1alpha1.User, error)
	UpdateStatus(ctx context.Context, user *v1alpha1.User, opts v1.UpdateOptions) (*v1alpha1.User, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.User, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.UserList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.User, err error)
	UserExpansion
}

// users implements UserInterface
type users struct {
	client rest.Interface
	ns     string
}

// newUsers returns a Users
func newUsers(c *UserV1alpha1Client, namespace string) *users {
	return &users{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the user, and returns the corresponding user object, and an error if there is any.
func (c *users) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.User, err error) {
	result = &v1alpha1.User{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("users").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Users that match those selectors.
func (c *users) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("users").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested users.
func (c *users) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("users").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a user and creates it.  Returns the server's representation of the user, and an error, if there is any.
func (c *users) Create(ctx context.Context, user *v1alpha1.User, opts v1.CreateOptions) (result *v1alpha1.User, err error) {
	result = &v1alpha1.User{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("users").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(user).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a user and updates it. Returns the server's representation of the user, and an error, if there is any.
func (c *users) Update(ctx context.Context, user *v1alpha1.User, opts v1.UpdateOptions) (result *v1alpha1.User, err error) {
	result = &v1alpha1.User{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("users").
		Name(user.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(user).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *users) UpdateStatus(ctx context.Context, user *v1alpha1.User, opts v1.UpdateOptions) (result *v1alpha1.User, err error) {
	result = &v1alpha1.User{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("users").
		Name(user.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(user).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the user and deletes it. Returns an error if one occurs.
func (c *users) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("users").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *users) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("users").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched user.
func (c *users) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.User, err error) {
	result = &v1alpha1.User{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("users").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
