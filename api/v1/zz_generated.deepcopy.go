//go:build !ignore_autogenerated

//
// Copyright 2022 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSVInjector) DeepCopyInto(out *CSVInjector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSVInjector.
func (in *CSVInjector) DeepCopy() *CSVInjector {
	if in == nil {
		return nil
	}
	out := new(CSVInjector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAcceptance) DeepCopyInto(out *LicenseAcceptance) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAcceptance.
func (in *LicenseAcceptance) DeepCopy() *LicenseAcceptance {
	if in == nil {
		return nil
	}
	out := new(LicenseAcceptance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceScope) DeepCopyInto(out *NamespaceScope) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceScope.
func (in *NamespaceScope) DeepCopy() *NamespaceScope {
	if in == nil {
		return nil
	}
	out := new(NamespaceScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceScope) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceScopeList) DeepCopyInto(out *NamespaceScopeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceScope, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceScopeList.
func (in *NamespaceScopeList) DeepCopy() *NamespaceScopeList {
	if in == nil {
		return nil
	}
	out := new(NamespaceScopeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceScopeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceScopeSpec) DeepCopyInto(out *NamespaceScopeSpec) {
	*out = *in
	if in.NamespaceMembers != nil {
		in, out := &in.NamespaceMembers, &out.NamespaceMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceAccountMembers != nil {
		in, out := &in.ServiceAccountMembers, &out.ServiceAccountMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RestartLabels != nil {
		in, out := &in.RestartLabels, &out.RestartLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.CSVInjector = in.CSVInjector
	out.License = in.License
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceScopeSpec.
func (in *NamespaceScopeSpec) DeepCopy() *NamespaceScopeSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceScopeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceScopeStatus) DeepCopyInto(out *NamespaceScopeStatus) {
	*out = *in
	if in.ValidatedMembers != nil {
		in, out := &in.ValidatedMembers, &out.ValidatedMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ManagedCSVList != nil {
		in, out := &in.ManagedCSVList, &out.ManagedCSVList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PatchedCSVList != nil {
		in, out := &in.PatchedCSVList, &out.PatchedCSVList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ManagedWebhookList != nil {
		in, out := &in.ManagedWebhookList, &out.ManagedWebhookList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PatchedWebhookList != nil {
		in, out := &in.PatchedWebhookList, &out.PatchedWebhookList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceScopeStatus.
func (in *NamespaceScopeStatus) DeepCopy() *NamespaceScopeStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceScopeStatus)
	in.DeepCopyInto(out)
	return out
}
