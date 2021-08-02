// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Play) DeepCopyInto(out *Play) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Play.
func (in *Play) DeepCopy() *Play {
	if in == nil {
		return nil
	}
	out := new(Play)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Play) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayList) DeepCopyInto(out *PlayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Play, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayList.
func (in *PlayList) DeepCopy() *PlayList {
	if in == nil {
		return nil
	}
	out := new(PlayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpec) DeepCopyInto(out *PlaySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PlaySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpec.
func (in *PlaySpec) DeepCopy() *PlaySpec {
	if in == nil {
		return nil
	}
	out := new(PlaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpecResource) DeepCopyInto(out *PlaySpecResource) {
	*out = *in
	if in.ConferenceNumber != nil {
		in, out := &in.ConferenceNumber, &out.ConferenceNumber
		*out = new(string)
		**out = **in
	}
	if in.ConferenceURL != nil {
		in, out := &in.ConferenceURL, &out.ConferenceURL
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Responder != nil {
		in, out := &in.Responder, &out.Responder
		*out = make([]PlaySpecResponder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RespondersMessage != nil {
		in, out := &in.RespondersMessage, &out.RespondersMessage
		*out = new(string)
		**out = **in
	}
	if in.Runnability != nil {
		in, out := &in.Runnability, &out.Runnability
		*out = new(string)
		**out = **in
	}
	if in.Subscriber != nil {
		in, out := &in.Subscriber, &out.Subscriber
		*out = make([]PlaySpecSubscriber, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubscribersMessage != nil {
		in, out := &in.SubscribersMessage, &out.SubscribersMessage
		*out = new(string)
		**out = **in
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpecResource.
func (in *PlaySpecResource) DeepCopy() *PlaySpecResource {
	if in == nil {
		return nil
	}
	out := new(PlaySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpecResponder) DeepCopyInto(out *PlaySpecResponder) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EscalationRule != nil {
		in, out := &in.EscalationRule, &out.EscalationRule
		*out = make([]PlaySpecResponderEscalationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumLoops != nil {
		in, out := &in.NumLoops, &out.NumLoops
		*out = new(int64)
		**out = **in
	}
	if in.OnCallHandoffNotifications != nil {
		in, out := &in.OnCallHandoffNotifications, &out.OnCallHandoffNotifications
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make([]PlaySpecResponderService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = make([]PlaySpecResponderTeam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpecResponder.
func (in *PlaySpecResponder) DeepCopy() *PlaySpecResponder {
	if in == nil {
		return nil
	}
	out := new(PlaySpecResponder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpecResponderEscalationRule) DeepCopyInto(out *PlaySpecResponderEscalationRule) {
	*out = *in
	if in.EscalationDelayInMinutes != nil {
		in, out := &in.EscalationDelayInMinutes, &out.EscalationDelayInMinutes
		*out = new(int64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]PlaySpecResponderEscalationRuleTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpecResponderEscalationRule.
func (in *PlaySpecResponderEscalationRule) DeepCopy() *PlaySpecResponderEscalationRule {
	if in == nil {
		return nil
	}
	out := new(PlaySpecResponderEscalationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpecResponderEscalationRuleTarget) DeepCopyInto(out *PlaySpecResponderEscalationRuleTarget) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpecResponderEscalationRuleTarget.
func (in *PlaySpecResponderEscalationRuleTarget) DeepCopy() *PlaySpecResponderEscalationRuleTarget {
	if in == nil {
		return nil
	}
	out := new(PlaySpecResponderEscalationRuleTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpecResponderService) DeepCopyInto(out *PlaySpecResponderService) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpecResponderService.
func (in *PlaySpecResponderService) DeepCopy() *PlaySpecResponderService {
	if in == nil {
		return nil
	}
	out := new(PlaySpecResponderService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpecResponderTeam) DeepCopyInto(out *PlaySpecResponderTeam) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpecResponderTeam.
func (in *PlaySpecResponderTeam) DeepCopy() *PlaySpecResponderTeam {
	if in == nil {
		return nil
	}
	out := new(PlaySpecResponderTeam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaySpecSubscriber) DeepCopyInto(out *PlaySpecSubscriber) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaySpecSubscriber.
func (in *PlaySpecSubscriber) DeepCopy() *PlaySpecSubscriber {
	if in == nil {
		return nil
	}
	out := new(PlaySpecSubscriber)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlayStatus) DeepCopyInto(out *PlayStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlayStatus.
func (in *PlayStatus) DeepCopy() *PlayStatus {
	if in == nil {
		return nil
	}
	out := new(PlayStatus)
	in.DeepCopyInto(out)
	return out
}
