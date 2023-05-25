/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1beta1ComponentDefinitionSpec struct for V1beta1ComponentDefinitionSpec
type V1beta1ComponentDefinitionSpec struct {
	ChildResourceKinds []CommonChildResourceKind `json:"childResourceKinds,omitempty"`
	Extension string `json:"extension,omitempty"`
	PodSpecPath string `json:"podSpecPath,omitempty"`
	RevisionLabel string `json:"revisionLabel,omitempty"`
	Schematic CommonSchematic `json:"schematic,omitempty"`
	Status CommonStatus `json:"status,omitempty"`
	Workload CommonWorkloadTypeDescriptor `json:"workload"`
}