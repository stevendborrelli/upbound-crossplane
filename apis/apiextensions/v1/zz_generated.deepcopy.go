// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	commonv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComposedTemplate) DeepCopyInto(out *ComposedTemplate) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	in.Base.DeepCopyInto(&out.Base)
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConnectionDetails != nil {
		in, out := &in.ConnectionDetails, &out.ConnectionDetails
		*out = make([]ConnectionDetail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReadinessChecks != nil {
		in, out := &in.ReadinessChecks, &out.ReadinessChecks
		*out = make([]ReadinessCheck, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComposedTemplate.
func (in *ComposedTemplate) DeepCopy() *ComposedTemplate {
	if in == nil {
		return nil
	}
	out := new(ComposedTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeResourceDefinition) DeepCopyInto(out *CompositeResourceDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeResourceDefinition.
func (in *CompositeResourceDefinition) DeepCopy() *CompositeResourceDefinition {
	if in == nil {
		return nil
	}
	out := new(CompositeResourceDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositeResourceDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeResourceDefinitionControllerStatus) DeepCopyInto(out *CompositeResourceDefinitionControllerStatus) {
	*out = *in
	out.CompositeResourceTypeRef = in.CompositeResourceTypeRef
	out.CompositeResourceClaimTypeRef = in.CompositeResourceClaimTypeRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeResourceDefinitionControllerStatus.
func (in *CompositeResourceDefinitionControllerStatus) DeepCopy() *CompositeResourceDefinitionControllerStatus {
	if in == nil {
		return nil
	}
	out := new(CompositeResourceDefinitionControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeResourceDefinitionList) DeepCopyInto(out *CompositeResourceDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CompositeResourceDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeResourceDefinitionList.
func (in *CompositeResourceDefinitionList) DeepCopy() *CompositeResourceDefinitionList {
	if in == nil {
		return nil
	}
	out := new(CompositeResourceDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositeResourceDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeResourceDefinitionSpec) DeepCopyInto(out *CompositeResourceDefinitionSpec) {
	*out = *in
	in.Names.DeepCopyInto(&out.Names)
	if in.ClaimNames != nil {
		in, out := &in.ClaimNames, &out.ClaimNames
		*out = new(apiextensionsv1.CustomResourceDefinitionNames)
		(*in).DeepCopyInto(*out)
	}
	if in.ConnectionSecretKeys != nil {
		in, out := &in.ConnectionSecretKeys, &out.ConnectionSecretKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultCompositionRef != nil {
		in, out := &in.DefaultCompositionRef, &out.DefaultCompositionRef
		*out = new(commonv1.Reference)
		**out = **in
	}
	if in.EnforcedCompositionRef != nil {
		in, out := &in.EnforcedCompositionRef, &out.EnforcedCompositionRef
		*out = new(commonv1.Reference)
		**out = **in
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]CompositeResourceDefinitionVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeResourceDefinitionSpec.
func (in *CompositeResourceDefinitionSpec) DeepCopy() *CompositeResourceDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(CompositeResourceDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeResourceDefinitionStatus) DeepCopyInto(out *CompositeResourceDefinitionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	out.Controllers = in.Controllers
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeResourceDefinitionStatus.
func (in *CompositeResourceDefinitionStatus) DeepCopy() *CompositeResourceDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(CompositeResourceDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeResourceDefinitionVersion) DeepCopyInto(out *CompositeResourceDefinitionVersion) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(CompositeResourceValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalPrinterColumns != nil {
		in, out := &in.AdditionalPrinterColumns, &out.AdditionalPrinterColumns
		*out = make([]apiextensionsv1.CustomResourceColumnDefinition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeResourceDefinitionVersion.
func (in *CompositeResourceDefinitionVersion) DeepCopy() *CompositeResourceDefinitionVersion {
	if in == nil {
		return nil
	}
	out := new(CompositeResourceDefinitionVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeResourceValidation) DeepCopyInto(out *CompositeResourceValidation) {
	*out = *in
	in.OpenAPIV3Schema.DeepCopyInto(&out.OpenAPIV3Schema)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeResourceValidation.
func (in *CompositeResourceValidation) DeepCopy() *CompositeResourceValidation {
	if in == nil {
		return nil
	}
	out := new(CompositeResourceValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Composition) DeepCopyInto(out *Composition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Composition.
func (in *Composition) DeepCopy() *Composition {
	if in == nil {
		return nil
	}
	out := new(Composition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Composition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositionList) DeepCopyInto(out *CompositionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Composition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositionList.
func (in *CompositionList) DeepCopy() *CompositionList {
	if in == nil {
		return nil
	}
	out := new(CompositionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositionSpec) DeepCopyInto(out *CompositionSpec) {
	*out = *in
	out.CompositeTypeRef = in.CompositeTypeRef
	if in.PatchSets != nil {
		in, out := &in.PatchSets, &out.PatchSets
		*out = make([]PatchSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ComposedTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WriteConnectionSecretsToNamespace != nil {
		in, out := &in.WriteConnectionSecretsToNamespace, &out.WriteConnectionSecretsToNamespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositionSpec.
func (in *CompositionSpec) DeepCopy() *CompositionSpec {
	if in == nil {
		return nil
	}
	out := new(CompositionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositionStatus) DeepCopyInto(out *CompositionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositionStatus.
func (in *CompositionStatus) DeepCopy() *CompositionStatus {
	if in == nil {
		return nil
	}
	out := new(CompositionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionDetail) DeepCopyInto(out *ConnectionDetail) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ConnectionDetailType)
		**out = **in
	}
	if in.FromConnectionSecretKey != nil {
		in, out := &in.FromConnectionSecretKey, &out.FromConnectionSecretKey
		*out = new(string)
		**out = **in
	}
	if in.FromFieldPath != nil {
		in, out := &in.FromFieldPath, &out.FromFieldPath
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionDetail.
func (in *ConnectionDetail) DeepCopy() *ConnectionDetail {
	if in == nil {
		return nil
	}
	out := new(ConnectionDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConvertTransform) DeepCopyInto(out *ConvertTransform) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConvertTransform.
func (in *ConvertTransform) DeepCopy() *ConvertTransform {
	if in == nil {
		return nil
	}
	out := new(ConvertTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MapTransform) DeepCopyInto(out *MapTransform) {
	*out = *in
	if in.Pairs != nil {
		in, out := &in.Pairs, &out.Pairs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MapTransform.
func (in *MapTransform) DeepCopy() *MapTransform {
	if in == nil {
		return nil
	}
	out := new(MapTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MathTransform) DeepCopyInto(out *MathTransform) {
	*out = *in
	if in.Multiply != nil {
		in, out := &in.Multiply, &out.Multiply
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MathTransform.
func (in *MathTransform) DeepCopy() *MathTransform {
	if in == nil {
		return nil
	}
	out := new(MathTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Patch) DeepCopyInto(out *Patch) {
	*out = *in
	if in.FromFieldPath != nil {
		in, out := &in.FromFieldPath, &out.FromFieldPath
		*out = new(string)
		**out = **in
	}
	if in.ToFieldPath != nil {
		in, out := &in.ToFieldPath, &out.ToFieldPath
		*out = new(string)
		**out = **in
	}
	if in.PatchSetName != nil {
		in, out := &in.PatchSetName, &out.PatchSetName
		*out = new(string)
		**out = **in
	}
	if in.Transforms != nil {
		in, out := &in.Transforms, &out.Transforms
		*out = make([]Transform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Patch.
func (in *Patch) DeepCopy() *Patch {
	if in == nil {
		return nil
	}
	out := new(Patch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchSet) DeepCopyInto(out *PatchSet) {
	*out = *in
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchSet.
func (in *PatchSet) DeepCopy() *PatchSet {
	if in == nil {
		return nil
	}
	out := new(PatchSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheck) DeepCopyInto(out *ReadinessCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheck.
func (in *ReadinessCheck) DeepCopy() *ReadinessCheck {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringTransform) DeepCopyInto(out *StringTransform) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringTransform.
func (in *StringTransform) DeepCopy() *StringTransform {
	if in == nil {
		return nil
	}
	out := new(StringTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transform) DeepCopyInto(out *Transform) {
	*out = *in
	if in.Math != nil {
		in, out := &in.Math, &out.Math
		*out = new(MathTransform)
		(*in).DeepCopyInto(*out)
	}
	if in.Map != nil {
		in, out := &in.Map, &out.Map
		*out = new(MapTransform)
		(*in).DeepCopyInto(*out)
	}
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = new(StringTransform)
		**out = **in
	}
	if in.Convert != nil {
		in, out := &in.Convert, &out.Convert
		*out = new(ConvertTransform)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transform.
func (in *Transform) DeepCopy() *Transform {
	if in == nil {
		return nil
	}
	out := new(Transform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TypeReference) DeepCopyInto(out *TypeReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeReference.
func (in *TypeReference) DeepCopy() *TypeReference {
	if in == nil {
		return nil
	}
	out := new(TypeReference)
	in.DeepCopyInto(out)
	return out
}
