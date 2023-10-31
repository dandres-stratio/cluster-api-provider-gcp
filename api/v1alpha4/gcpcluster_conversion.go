/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	infrav1beta1 "sigs.k8s.io/cluster-api-provider-gcp/api/v1beta1"
	v1beta1 "sigs.k8s.io/cluster-api-provider-gcp/api/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this GCPCluster to the Hub version (v1beta1).
func (src *GCPCluster) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1beta1.GCPCluster)

	if err := Convert_v1alpha4_GCPCluster_To_v1beta1_GCPCluster(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1beta1.GCPCluster{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	for _, restoredSubnet := range restored.Spec.Network.Subnets {
		for i, dstSubnet := range dst.Spec.Network.Subnets {
			if dstSubnet.Name != restoredSubnet.Name {
				continue
			}
			dst.Spec.Network.Subnets[i].Purpose = restoredSubnet.Purpose

			break
		}
	}
	if restored.Spec.CredentialsRef != nil {
		dst.Spec.CredentialsRef = restored.Spec.CredentialsRef.DeepCopy()
	}

	for _, restoredTag := range restored.Spec.ResourceManagerTags {
		dst.Spec.ResourceManagerTags = append(dst.Spec.ResourceManagerTags, *restoredTag.DeepCopy())
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *GCPCluster) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1beta1.GCPCluster)

	if err := Convert_v1beta1_GCPCluster_To_v1alpha4_GCPCluster(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

// ConvertTo converts this GCPClusterList to the Hub version (v1beta1).
func (src *GCPClusterList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1beta1.GCPClusterList)
	return Convert_v1alpha4_GCPClusterList_To_v1beta1_GCPClusterList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *GCPClusterList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1beta1.GCPClusterList)
	return Convert_v1beta1_GCPClusterList_To_v1alpha4_GCPClusterList(src, dst, nil)
}

// Convert_v1beta1_SubnetSpec_To_v1alpha4_SubnetSpec is an autogenerated conversion function.
func Convert_v1beta1_SubnetSpec_To_v1alpha4_SubnetSpec(in *infrav1beta1.SubnetSpec, out *SubnetSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_SubnetSpec_To_v1alpha4_SubnetSpec(in, out, s)
}

// Convert_v1beta1_GCPClusterSpec_To_v1alpha4_GCPClusterSpec is an autogenerated conversion function.
func Convert_v1beta1_GCPClusterSpec_To_v1alpha4_GCPClusterSpec(in *v1beta1.GCPClusterSpec, out *GCPClusterSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_GCPClusterSpec_To_v1alpha4_GCPClusterSpec(in, out, s)
}
