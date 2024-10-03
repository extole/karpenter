//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package state

import (
	"k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	"github.com/extole/karpenter/pkg/apis/v1beta1"
	"github.com/extole/karpenter/pkg/scheduling"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateNode) DeepCopyInto(out *StateNode) {
	*out = *in
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = new(v1.Node)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeClaim != nil {
		in, out := &in.NodeClaim, &out.NodeClaim
		*out = new(v1beta1.NodeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.daemonSetRequests != nil {
		in, out := &in.daemonSetRequests, &out.daemonSetRequests
		*out = make(map[types.NamespacedName]v1.ResourceList, len(*in))
		for key, val := range *in {
			var outVal map[v1.ResourceName]resource.Quantity
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(v1.ResourceList, len(*in))
				for key, val := range *in {
					(*out)[key] = val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.daemonSetLimits != nil {
		in, out := &in.daemonSetLimits, &out.daemonSetLimits
		*out = make(map[types.NamespacedName]v1.ResourceList, len(*in))
		for key, val := range *in {
			var outVal map[v1.ResourceName]resource.Quantity
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(v1.ResourceList, len(*in))
				for key, val := range *in {
					(*out)[key] = val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.podRequests != nil {
		in, out := &in.podRequests, &out.podRequests
		*out = make(map[types.NamespacedName]v1.ResourceList, len(*in))
		for key, val := range *in {
			var outVal map[v1.ResourceName]resource.Quantity
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(v1.ResourceList, len(*in))
				for key, val := range *in {
					(*out)[key] = val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.podLimits != nil {
		in, out := &in.podLimits, &out.podLimits
		*out = make(map[types.NamespacedName]v1.ResourceList, len(*in))
		for key, val := range *in {
			var outVal map[v1.ResourceName]resource.Quantity
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(v1.ResourceList, len(*in))
				for key, val := range *in {
					(*out)[key] = val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.hostPortUsage != nil {
		in, out := &in.hostPortUsage, &out.hostPortUsage
		*out = new(scheduling.HostPortUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.volumeUsage != nil {
		in, out := &in.volumeUsage, &out.volumeUsage
		*out = new(scheduling.VolumeUsage)
		(*in).DeepCopyInto(*out)
	}
	in.nominatedUntil.DeepCopyInto(&out.nominatedUntil)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateNode.
func (in *StateNode) DeepCopy() *StateNode {
	if in == nil {
		return nil
	}
	out := new(StateNode)
	in.DeepCopyInto(out)
	return out
}
