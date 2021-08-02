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
	v1alpha1 "kubeform.dev/provider-pagerduty-api/apis/ruleset/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RuleLister helps list Rules.
// All objects returned here must be treated as read-only.
type RuleLister interface {
	// List lists all Rules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Rule, err error)
	// Rules returns an object that can list and get Rules.
	Rules(namespace string) RuleNamespaceLister
	RuleListerExpansion
}

// ruleLister implements the RuleLister interface.
type ruleLister struct {
	indexer cache.Indexer
}

// NewRuleLister returns a new RuleLister.
func NewRuleLister(indexer cache.Indexer) RuleLister {
	return &ruleLister{indexer: indexer}
}

// List lists all Rules in the indexer.
func (s *ruleLister) List(selector labels.Selector) (ret []*v1alpha1.Rule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Rule))
	})
	return ret, err
}

// Rules returns an object that can list and get Rules.
func (s *ruleLister) Rules(namespace string) RuleNamespaceLister {
	return ruleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RuleNamespaceLister helps list and get Rules.
// All objects returned here must be treated as read-only.
type RuleNamespaceLister interface {
	// List lists all Rules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Rule, err error)
	// Get retrieves the Rule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Rule, error)
	RuleNamespaceListerExpansion
}

// ruleNamespaceLister implements the RuleNamespaceLister
// interface.
type ruleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Rules in the indexer for a given namespace.
func (s ruleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Rule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Rule))
	})
	return ret, err
}

// Get retrieves the Rule from the indexer for a given namespace and name.
func (s ruleNamespaceLister) Get(name string) (*v1alpha1.Rule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rule"), name)
	}
	return obj.(*v1alpha1.Rule), nil
}
