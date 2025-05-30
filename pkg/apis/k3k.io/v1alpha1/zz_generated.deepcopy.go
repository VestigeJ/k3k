//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLimit) DeepCopyInto(out *ClusterLimit) {
	*out = *in
	if in.ServerLimit != nil {
		in, out := &in.ServerLimit, &out.ServerLimit
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.WorkerLimit != nil {
		in, out := &in.WorkerLimit, &out.WorkerLimit
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLimit.
func (in *ClusterLimit) DeepCopy() *ClusterLimit {
	if in == nil {
		return nil
	}
	out := new(ClusterLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	out.TypeMeta = in.TypeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSet) DeepCopyInto(out *ClusterSet) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSet.
func (in *ClusterSet) DeepCopy() *ClusterSet {
	if in == nil {
		return nil
	}
	out := new(ClusterSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetList) DeepCopyInto(out *ClusterSetList) {
	*out = *in
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	out.TypeMeta = in.TypeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetList.
func (in *ClusterSetList) DeepCopy() *ClusterSetList {
	if in == nil {
		return nil
	}
	out := new(ClusterSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetSpec) DeepCopyInto(out *ClusterSetSpec) {
	*out = *in
	if in.MaxLimits != nil {
		in, out := &in.MaxLimits, &out.MaxLimits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.DefaultLimits != nil {
		in, out := &in.DefaultLimits, &out.DefaultLimits
		*out = new(ClusterLimit)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultNodeSelector != nil {
		in, out := &in.DefaultNodeSelector, &out.DefaultNodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AllowedModeTypes != nil {
		in, out := &in.AllowedModeTypes, &out.AllowedModeTypes
		*out = make([]ClusterMode, len(*in))
		copy(*out, *in)
	}
	if in.PodSecurityAdmissionLevel != nil {
		in, out := &in.PodSecurityAdmissionLevel, &out.PodSecurityAdmissionLevel
		*out = new(PodSecurityAdmissionLevel)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetSpec.
func (in *ClusterSetSpec) DeepCopy() *ClusterSetSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetStatus) DeepCopyInto(out *ClusterSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetStatus.
func (in *ClusterSetStatus) DeepCopy() *ClusterSetStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = new(int32)
		**out = **in
	}
	if in.Agents != nil {
		in, out := &in.Agents, &out.Agents
		*out = new(int32)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = new(ClusterLimit)
		(*in).DeepCopyInto(*out)
	}
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.ServerArgs != nil {
		in, out := &in.ServerArgs, &out.ServerArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AgentArgs != nil {
		in, out := &in.AgentArgs, &out.AgentArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSSANs != nil {
		in, out := &in.TLSSANs, &out.TLSSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]Addon, len(*in))
		copy(*out, *in)
	}
	in.Persistence.DeepCopyInto(&out.Persistence)
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(ExposeConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.TLSSANs != nil {
		in, out := &in.TLSSANs, &out.TLSSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Persistence.DeepCopyInto(&out.Persistence)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposeConfig) DeepCopyInto(out *ExposeConfig) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(IngressConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(LoadBalancerConfig)
		**out = **in
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(NodePortConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposeConfig.
func (in *ExposeConfig) DeepCopy() *ExposeConfig {
	if in == nil {
		return nil
	}
	out := new(ExposeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressConfig) DeepCopyInto(out *IngressConfig) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressConfig.
func (in *IngressConfig) DeepCopy() *IngressConfig {
	if in == nil {
		return nil
	}
	out := new(IngressConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerConfig) DeepCopyInto(out *LoadBalancerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerConfig.
func (in *LoadBalancerConfig) DeepCopy() *LoadBalancerConfig {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePortConfig) DeepCopyInto(out *NodePortConfig) {
	*out = *in
	if in.ServerPort != nil {
		in, out := &in.ServerPort, &out.ServerPort
		*out = new(int32)
		**out = **in
	}
	if in.ServicePort != nil {
		in, out := &in.ServicePort, &out.ServicePort
		*out = new(int32)
		**out = **in
	}
	if in.ETCDPort != nil {
		in, out := &in.ETCDPort, &out.ETCDPort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePortConfig.
func (in *NodePortConfig) DeepCopy() *NodePortConfig {
	if in == nil {
		return nil
	}
	out := new(NodePortConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceConfig) DeepCopyInto(out *PersistenceConfig) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceConfig.
func (in *PersistenceConfig) DeepCopy() *PersistenceConfig {
	if in == nil {
		return nil
	}
	out := new(PersistenceConfig)
	in.DeepCopyInto(out)
	return out
}
