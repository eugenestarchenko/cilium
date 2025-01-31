//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package l2respondermap

import (
	bpf "github.com/cilium/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L2ResponderKey) DeepCopyInto(out *L2ResponderKey) {
	*out = *in
	in.IP.DeepCopyInto(&out.IP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L2ResponderKey.
func (in *L2ResponderKey) DeepCopy() *L2ResponderKey {
	if in == nil {
		return nil
	}
	out := new(L2ResponderKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *L2ResponderKey) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L2ResponderStats) DeepCopyInto(out *L2ResponderStats) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L2ResponderStats.
func (in *L2ResponderStats) DeepCopy() *L2ResponderStats {
	if in == nil {
		return nil
	}
	out := new(L2ResponderStats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *L2ResponderStats) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
