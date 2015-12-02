/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-conversions.sh

package v1alpha1

import (
	reflect "reflect"

	api "k8s.io/kubernetes/pkg/api"
	componentconfig "k8s.io/kubernetes/pkg/apis/componentconfig"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func autoconvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*componentconfig.KubeProxyConfiguration))(in)
	}
	if err := s.Convert(&in.TypeMeta, &out.TypeMeta, 0); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.CleanupIPTables = in.CleanupIPTables
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = int32(in.HealthzPort)
	out.HostnameOverride = in.HostnameOverride
	out.IPTablesSyncePeriodSeconds = int32(in.IPTablesSyncePeriodSeconds)
	out.KubeAPIBurst = int32(in.KubeAPIBurst)
	out.KubeAPIQPS = int32(in.KubeAPIQPS)
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		out.OOMScoreAdj = new(int32)
		*out.OOMScoreAdj = int32(*in.OOMScoreAdj)
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPTimeoutMilliseconds = int32(in.UDPTimeoutMilliseconds)
	return nil
}

func convert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	return autoconvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in, out, s)
}

func autoconvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*KubeProxyConfiguration))(in)
	}
	if err := s.Convert(&in.TypeMeta, &out.TypeMeta, 0); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.CleanupIPTables = in.CleanupIPTables
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = int(in.HealthzPort)
	out.HostnameOverride = in.HostnameOverride
	out.IPTablesSyncePeriodSeconds = int(in.IPTablesSyncePeriodSeconds)
	out.KubeAPIBurst = int(in.KubeAPIBurst)
	out.KubeAPIQPS = int(in.KubeAPIQPS)
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		out.OOMScoreAdj = new(int)
		*out.OOMScoreAdj = int(*in.OOMScoreAdj)
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = componentconfig.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPTimeoutMilliseconds = int(in.UDPTimeoutMilliseconds)
	return nil
}

func convert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	return autoconvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in, out, s)
}

func init() {
	err := api.Scheme.AddGeneratedConversionFuncs(
		autoconvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration,
		autoconvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration,
	)
	if err != nil {
		// If one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}
