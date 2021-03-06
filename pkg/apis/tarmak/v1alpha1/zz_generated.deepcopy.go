/*
Sniperkit-Bot
- Status: analyzed
*/

// +build !ignore_autogenerated

// Copyright Jetstack Ltd. See LICENSE for details.

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"

	cluster_v1alpha1 "github.com/sniperkit/snk.fork.tarmak/pkg/apis/cluster/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterApplyFlags) DeepCopyInto(out *ClusterApplyFlags) {
	*out = *in
	if in.InfrastructureStacks != nil {
		in, out := &in.InfrastructureStacks, &out.InfrastructureStacks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterApplyFlags.
func (in *ClusterApplyFlags) DeepCopy() *ClusterApplyFlags {
	if in == nil {
		return nil
	}
	out := new(ClusterApplyFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDestroyFlags) DeepCopyInto(out *ClusterDestroyFlags) {
	*out = *in
	if in.InfrastructureStacks != nil {
		in, out := &in.InfrastructureStacks, &out.InfrastructureStacks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDestroyFlags.
func (in *ClusterDestroyFlags) DeepCopy() *ClusterDestroyFlags {
	if in == nil {
		return nil
	}
	out := new(ClusterDestroyFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFlags) DeepCopyInto(out *ClusterFlags) {
	*out = *in
	in.Apply.DeepCopyInto(&out.Apply)
	in.Destroy.DeepCopyInto(&out.Destroy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFlags.
func (in *ClusterFlags) DeepCopy() *ClusterFlags {
	if in == nil {
		return nil
	}
	out := new(ClusterFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]cluster_v1alpha1.Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]Provider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Environments != nil {
		in, out := &in.Environments, &out.Environments
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Config) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigList) DeepCopyInto(out *ConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Config, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigList.
func (in *ConfigList) DeepCopy() *ConfigList {
	if in == nil {
		return nil
	}
	out := new(ConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.SSH != nil {
		in, out := &in.SSH, &out.SSH
		if *in == nil {
			*out = nil
		} else {
			*out = new(cluster_v1alpha1.SSH)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flags) DeepCopyInto(out *Flags) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flags.
func (in *Flags) DeepCopy() *Flags {
	if in == nil {
		return nil
	}
	out := new(Flags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Amazon != nil {
		in, out := &in.Amazon, &out.Amazon
		if *in == nil {
			*out = nil
		} else {
			*out = new(ProviderAmazon)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.GCP != nil {
		in, out := &in.GCP, &out.GCP
		if *in == nil {
			*out = nil
		} else {
			*out = new(ProviderGCP)
			**out = **in
		}
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		if *in == nil {
			*out = nil
		} else {
			*out = new(ProviderAzure)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provider.
func (in *Provider) DeepCopy() *Provider {
	if in == nil {
		return nil
	}
	out := new(Provider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Provider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAmazon) DeepCopyInto(out *ProviderAmazon) {
	*out = *in
	if in.AllowedAccountIDs != nil {
		in, out := &in.AllowedAccountIDs, &out.AllowedAccountIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAmazon.
func (in *ProviderAmazon) DeepCopy() *ProviderAmazon {
	if in == nil {
		return nil
	}
	out := new(ProviderAmazon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderAzure) DeepCopyInto(out *ProviderAzure) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderAzure.
func (in *ProviderAzure) DeepCopy() *ProviderAzure {
	if in == nil {
		return nil
	}
	out := new(ProviderAzure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderGCP) DeepCopyInto(out *ProviderGCP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderGCP.
func (in *ProviderGCP) DeepCopy() *ProviderGCP {
	if in == nil {
		return nil
	}
	out := new(ProviderGCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderList) DeepCopyInto(out *ProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Provider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderList.
func (in *ProviderList) DeepCopy() *ProviderList {
	if in == nil {
		return nil
	}
	out := new(ProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
