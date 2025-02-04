//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummary) DeepCopyInto(out *ClusterSummary) {
	*out = *in
	in.NumStatistic.DeepCopyInto(&out.NumStatistic)
	if in.ClusterConditions != nil {
		in, out := &in.ClusterConditions, &out.ClusterConditions
		*out = make(map[string][]v1.Condition, len(*in))
		for key, val := range *in {
			var outVal []v1.Condition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]v1.Condition, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummary.
func (in *ClusterSummary) DeepCopy() *ClusterSummary {
	if in == nil {
		return nil
	}
	out := new(ClusterSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneCfg) DeepCopyInto(out *ControlPlaneCfg) {
	*out = *in
	if in.EndPointCfg != nil {
		in, out := &in.EndPointCfg, &out.EndPointCfg
		*out = new(EndPointCfg)
		(*in).DeepCopyInto(*out)
	}
	if in.ETCD != nil {
		in, out := &in.ETCD, &out.ETCD
		*out = new(ETCD)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]Component, len(*in))
		copy(*out, *in)
	}
	if in.InstallCRD != nil {
		in, out := &in.InstallCRD, &out.InstallCRD
		*out = new(bool)
		**out = **in
	}
	if in.ChartExtraValues != nil {
		in, out := &in.ChartExtraValues, &out.ChartExtraValues
		*out = new(LocalConfigmapReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneCfg.
func (in *ControlPlaneCfg) DeepCopy() *ControlPlaneCfg {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCD) DeepCopyInto(out *ETCD) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCD.
func (in *ETCD) DeepCopy() *ETCD {
	if in == nil {
		return nil
	}
	out := new(ETCD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndPointCfg) DeepCopyInto(out *EndPointCfg) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalSecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndPointCfg.
func (in *EndPointCfg) DeepCopy() *EndPointCfg {
	if in == nil {
		return nil
	}
	out := new(EndPointCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Images) DeepCopyInto(out *Images) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Images.
func (in *Images) DeepCopy() *Images {
	if in == nil {
		return nil
	}
	out := new(Images)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaDeployment) DeepCopyInto(out *KarmadaDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaDeployment.
func (in *KarmadaDeployment) DeepCopy() *KarmadaDeployment {
	if in == nil {
		return nil
	}
	out := new(KarmadaDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KarmadaDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaDeploymentList) DeepCopyInto(out *KarmadaDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KarmadaDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaDeploymentList.
func (in *KarmadaDeploymentList) DeepCopy() *KarmadaDeploymentList {
	if in == nil {
		return nil
	}
	out := new(KarmadaDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KarmadaDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaDeploymentSpec) DeepCopyInto(out *KarmadaDeploymentSpec) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(InstallMode)
		**out = **in
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = new(Images)
		**out = **in
	}
	in.ControlPlane.DeepCopyInto(&out.ControlPlane)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaDeploymentSpec.
func (in *KarmadaDeploymentSpec) DeepCopy() *KarmadaDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(KarmadaDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaDeploymentStatus) DeepCopyInto(out *KarmadaDeploymentStatus) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalSecretReference)
		**out = **in
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(KarmadaResourceSummary)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaDeploymentStatus.
func (in *KarmadaDeploymentStatus) DeepCopy() *KarmadaDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(KarmadaDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KarmadaResourceSummary) DeepCopyInto(out *KarmadaResourceSummary) {
	*out = *in
	if in.ResourceTotalNum != nil {
		in, out := &in.ResourceTotalNum, &out.ResourceTotalNum
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PolicyTotalNum != nil {
		in, out := &in.PolicyTotalNum, &out.PolicyTotalNum
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ClusterSummary != nil {
		in, out := &in.ClusterSummary, &out.ClusterSummary
		*out = new(ClusterSummary)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSummary != nil {
		in, out := &in.NodeSummary, &out.NodeSummary
		*out = new(NodeSummary)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceSummary != nil {
		in, out := &in.ResourceSummary, &out.ResourceSummary
		*out = new(ResourceSummary)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkLoadSummary != nil {
		in, out := &in.WorkLoadSummary, &out.WorkLoadSummary
		*out = make(map[string]NumStatistic, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KarmadaResourceSummary.
func (in *KarmadaResourceSummary) DeepCopy() *KarmadaResourceSummary {
	if in == nil {
		return nil
	}
	out := new(KarmadaResourceSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalConfigmapReference) DeepCopyInto(out *LocalConfigmapReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalConfigmapReference.
func (in *LocalConfigmapReference) DeepCopy() *LocalConfigmapReference {
	if in == nil {
		return nil
	}
	out := new(LocalConfigmapReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSecretReference) DeepCopyInto(out *LocalSecretReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSecretReference.
func (in *LocalSecretReference) DeepCopy() *LocalSecretReference {
	if in == nil {
		return nil
	}
	out := new(LocalSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSummary) DeepCopyInto(out *NodeSummary) {
	*out = *in
	in.NumStatistic.DeepCopyInto(&out.NumStatistic)
	if in.ClusterNodeSummary != nil {
		in, out := &in.ClusterNodeSummary, &out.ClusterNodeSummary
		*out = make(map[string]NumStatistic, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSummary.
func (in *NodeSummary) DeepCopy() *NodeSummary {
	if in == nil {
		return nil
	}
	out := new(NodeSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumStatistic) DeepCopyInto(out *NumStatistic) {
	*out = *in
	if in.TotalNum != nil {
		in, out := &in.TotalNum, &out.TotalNum
		*out = new(int32)
		**out = **in
	}
	if in.ReadyNum != nil {
		in, out := &in.ReadyNum, &out.ReadyNum
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumStatistic.
func (in *NumStatistic) DeepCopy() *NumStatistic {
	if in == nil {
		return nil
	}
	out := new(NumStatistic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatistic) DeepCopyInto(out *ResourceStatistic) {
	*out = *in
	if in.Allocatable != nil {
		in, out := &in.Allocatable, &out.Allocatable
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Allocating != nil {
		in, out := &in.Allocating, &out.Allocating
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Allocated != nil {
		in, out := &in.Allocated, &out.Allocated
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatistic.
func (in *ResourceStatistic) DeepCopy() *ResourceStatistic {
	if in == nil {
		return nil
	}
	out := new(ResourceStatistic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSummary) DeepCopyInto(out *ResourceSummary) {
	*out = *in
	in.ResourceStatistic.DeepCopyInto(&out.ResourceStatistic)
	if in.ClusterResource != nil {
		in, out := &in.ClusterResource, &out.ClusterResource
		*out = make(map[string]ResourceStatistic, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSummary.
func (in *ResourceSummary) DeepCopy() *ResourceSummary {
	if in == nil {
		return nil
	}
	out := new(ResourceSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSummary) DeepCopyInto(out *WorkloadSummary) {
	*out = *in
	in.NumStatistic.DeepCopyInto(&out.NumStatistic)
	if in.ClusterWorkload != nil {
		in, out := &in.ClusterWorkload, &out.ClusterWorkload
		*out = make(map[string]NumStatistic, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSummary.
func (in *WorkloadSummary) DeepCopy() *WorkloadSummary {
	if in == nil {
		return nil
	}
	out := new(WorkloadSummary)
	in.DeepCopyInto(out)
	return out
}
