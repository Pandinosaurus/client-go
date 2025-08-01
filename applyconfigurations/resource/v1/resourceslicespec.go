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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// ResourceSliceSpecApplyConfiguration represents a declarative configuration of the ResourceSliceSpec type for use
// with apply.
type ResourceSliceSpecApplyConfiguration struct {
	Driver                 *string                                `json:"driver,omitempty"`
	Pool                   *ResourcePoolApplyConfiguration        `json:"pool,omitempty"`
	NodeName               *string                                `json:"nodeName,omitempty"`
	NodeSelector           *corev1.NodeSelectorApplyConfiguration `json:"nodeSelector,omitempty"`
	AllNodes               *bool                                  `json:"allNodes,omitempty"`
	Devices                []DeviceApplyConfiguration             `json:"devices,omitempty"`
	PerDeviceNodeSelection *bool                                  `json:"perDeviceNodeSelection,omitempty"`
	SharedCounters         []CounterSetApplyConfiguration         `json:"sharedCounters,omitempty"`
}

// ResourceSliceSpecApplyConfiguration constructs a declarative configuration of the ResourceSliceSpec type for use with
// apply.
func ResourceSliceSpec() *ResourceSliceSpecApplyConfiguration {
	return &ResourceSliceSpecApplyConfiguration{}
}

// WithDriver sets the Driver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Driver field is set to the value of the last call.
func (b *ResourceSliceSpecApplyConfiguration) WithDriver(value string) *ResourceSliceSpecApplyConfiguration {
	b.Driver = &value
	return b
}

// WithPool sets the Pool field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pool field is set to the value of the last call.
func (b *ResourceSliceSpecApplyConfiguration) WithPool(value *ResourcePoolApplyConfiguration) *ResourceSliceSpecApplyConfiguration {
	b.Pool = value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *ResourceSliceSpecApplyConfiguration) WithNodeName(value string) *ResourceSliceSpecApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithNodeSelector sets the NodeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeSelector field is set to the value of the last call.
func (b *ResourceSliceSpecApplyConfiguration) WithNodeSelector(value *corev1.NodeSelectorApplyConfiguration) *ResourceSliceSpecApplyConfiguration {
	b.NodeSelector = value
	return b
}

// WithAllNodes sets the AllNodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllNodes field is set to the value of the last call.
func (b *ResourceSliceSpecApplyConfiguration) WithAllNodes(value bool) *ResourceSliceSpecApplyConfiguration {
	b.AllNodes = &value
	return b
}

// WithDevices adds the given value to the Devices field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Devices field.
func (b *ResourceSliceSpecApplyConfiguration) WithDevices(values ...*DeviceApplyConfiguration) *ResourceSliceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithDevices")
		}
		b.Devices = append(b.Devices, *values[i])
	}
	return b
}

// WithPerDeviceNodeSelection sets the PerDeviceNodeSelection field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PerDeviceNodeSelection field is set to the value of the last call.
func (b *ResourceSliceSpecApplyConfiguration) WithPerDeviceNodeSelection(value bool) *ResourceSliceSpecApplyConfiguration {
	b.PerDeviceNodeSelection = &value
	return b
}

// WithSharedCounters adds the given value to the SharedCounters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SharedCounters field.
func (b *ResourceSliceSpecApplyConfiguration) WithSharedCounters(values ...*CounterSetApplyConfiguration) *ResourceSliceSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSharedCounters")
		}
		b.SharedCounters = append(b.SharedCounters, *values[i])
	}
	return b
}
