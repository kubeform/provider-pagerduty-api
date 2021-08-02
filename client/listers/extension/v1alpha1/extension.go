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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-pagerduty-api/apis/extension/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ExtensionLister helps list Extensions.
// All objects returned here must be treated as read-only.
type ExtensionLister interface {
	// List lists all Extensions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Extension, err error)
	// Extensions returns an object that can list and get Extensions.
	Extensions(namespace string) ExtensionNamespaceLister
	ExtensionListerExpansion
}

// extensionLister implements the ExtensionLister interface.
type extensionLister struct {
	indexer cache.Indexer
}

// NewExtensionLister returns a new ExtensionLister.
func NewExtensionLister(indexer cache.Indexer) ExtensionLister {
	return &extensionLister{indexer: indexer}
}

// List lists all Extensions in the indexer.
func (s *extensionLister) List(selector labels.Selector) (ret []*v1alpha1.Extension, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Extension))
	})
	return ret, err
}

// Extensions returns an object that can list and get Extensions.
func (s *extensionLister) Extensions(namespace string) ExtensionNamespaceLister {
	return extensionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExtensionNamespaceLister helps list and get Extensions.
// All objects returned here must be treated as read-only.
type ExtensionNamespaceLister interface {
	// List lists all Extensions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Extension, err error)
	// Get retrieves the Extension from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Extension, error)
	ExtensionNamespaceListerExpansion
}

// extensionNamespaceLister implements the ExtensionNamespaceLister
// interface.
type extensionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Extensions in the indexer for a given namespace.
func (s extensionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Extension, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Extension))
	})
	return ret, err
}

// Get retrieves the Extension from the indexer for a given namespace and name.
func (s extensionNamespaceLister) Get(name string) (*v1alpha1.Extension, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("extension"), name)
	}
	return obj.(*v1alpha1.Extension), nil
}