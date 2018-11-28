// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSProvisioner) DeepCopyInto(out *EFSProvisioner) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSProvisioner.
func (in *EFSProvisioner) DeepCopy() *EFSProvisioner {
	if in == nil {
		return nil
	}
	out := new(EFSProvisioner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EFSProvisioner) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSProvisionerList) DeepCopyInto(out *EFSProvisionerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EFSProvisioner, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSProvisionerList.
func (in *EFSProvisionerList) DeepCopy() *EFSProvisionerList {
	if in == nil {
		return nil
	}
	out := new(EFSProvisionerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EFSProvisionerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSProvisionerSpec) DeepCopyInto(out *EFSProvisionerSpec) {
	*out = *in
	out.OperatorSpec = in.OperatorSpec
	if in.ReclaimPolicy != nil {
		in, out := &in.ReclaimPolicy, &out.ReclaimPolicy
		*out = new(v1.PersistentVolumeReclaimPolicy)
		**out = **in
	}
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.BasePath != nil {
		in, out := &in.BasePath, &out.BasePath
		*out = new(string)
		**out = **in
	}
	if in.SupplementalGroup != nil {
		in, out := &in.SupplementalGroup, &out.SupplementalGroup
		*out = new(int64)
		**out = **in
	}
	if in.GidAllocate != nil {
		in, out := &in.GidAllocate, &out.GidAllocate
		*out = new(bool)
		**out = **in
	}
	if in.GidMin != nil {
		in, out := &in.GidMin, &out.GidMin
		*out = new(int)
		**out = **in
	}
	if in.GidMax != nil {
		in, out := &in.GidMax, &out.GidMax
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSProvisionerSpec.
func (in *EFSProvisionerSpec) DeepCopy() *EFSProvisionerSpec {
	if in == nil {
		return nil
	}
	out := new(EFSProvisionerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSProvisionerStatus) DeepCopyInto(out *EFSProvisionerStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSProvisionerStatus.
func (in *EFSProvisionerStatus) DeepCopy() *EFSProvisionerStatus {
	if in == nil {
		return nil
	}
	out := new(EFSProvisionerStatus)
	in.DeepCopyInto(out)
	return out
}
