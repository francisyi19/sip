// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraConfig) DeepCopyInto(out *InfraConfig) {
	*out = *in
	if in.OptionalData != nil {
		in, out := &in.OptionalData, &out.OptionalData
		*out = new(OptsConfig)
		**out = **in
	}
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraConfig.
func (in *InfraConfig) DeepCopy() *InfraConfig {
	if in == nil {
		return nil
	}
	out := new(InfraConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSet) DeepCopyInto(out *NodeSet) {
	*out = *in
	if in.Scheduling != nil {
		in, out := &in.Scheduling, &out.Scheduling
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(VmCount)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSet.
func (in *NodeSet) DeepCopy() *NodeSet {
	if in == nil {
		return nil
	}
	out := new(NodeSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptsConfig) DeepCopyInto(out *OptsConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptsConfig.
func (in *OptsConfig) DeepCopy() *OptsConfig {
	if in == nil {
		return nil
	}
	out := new(OptsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPCluster) DeepCopyInto(out *SIPCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPCluster.
func (in *SIPCluster) DeepCopy() *SIPCluster {
	if in == nil {
		return nil
	}
	out := new(SIPCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SIPCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPClusterList) DeepCopyInto(out *SIPClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SIPCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPClusterList.
func (in *SIPClusterList) DeepCopy() *SIPClusterList {
	if in == nil {
		return nil
	}
	out := new(SIPClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SIPClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPClusterSpec) DeepCopyInto(out *SIPClusterSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[VmRoles]NodeSet, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.InfraServices != nil {
		in, out := &in.InfraServices, &out.InfraServices
		*out = make(map[InfraService]InfraConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPClusterSpec.
func (in *SIPClusterSpec) DeepCopy() *SIPClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SIPClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPClusterStatus) DeepCopyInto(out *SIPClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPClusterStatus.
func (in *SIPClusterStatus) DeepCopy() *SIPClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SIPClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmCount) DeepCopyInto(out *VmCount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmCount.
func (in *VmCount) DeepCopy() *VmCount {
	if in == nil {
		return nil
	}
	out := new(VmCount)
	in.DeepCopyInto(out)
	return out
}
