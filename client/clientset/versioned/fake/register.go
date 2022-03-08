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

package fake

import (
	addonv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/addon/v1alpha1"
	businessv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/business/v1alpha1"
	escalationv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/escalation/v1alpha1"
	eventv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/event/v1alpha1"
	extensionv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/extension/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-pagerduty-api/apis/maintenance/v1alpha1"
	responsev1alpha1 "kubeform.dev/provider-pagerduty-api/apis/response/v1alpha1"
	rulesetv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/ruleset/v1alpha1"
	schedulev1alpha1 "kubeform.dev/provider-pagerduty-api/apis/schedule/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-pagerduty-api/apis/service/v1alpha1"
	slackv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/slack/v1alpha1"
	tagv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/tag/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/team/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/user/v1alpha1"
	webhookv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/webhook/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	addonv1alpha1.AddToScheme,
	businessv1alpha1.AddToScheme,
	escalationv1alpha1.AddToScheme,
	eventv1alpha1.AddToScheme,
	extensionv1alpha1.AddToScheme,
	maintenancev1alpha1.AddToScheme,
	responsev1alpha1.AddToScheme,
	rulesetv1alpha1.AddToScheme,
	schedulev1alpha1.AddToScheme,
	servicev1alpha1.AddToScheme,
	slackv1alpha1.AddToScheme,
	tagv1alpha1.AddToScheme,
	teamv1alpha1.AddToScheme,
	userv1alpha1.AddToScheme,
	webhookv1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
