/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1beta1Workflow struct for V1beta1Workflow
type V1beta1Workflow struct {
	Mode V1alpha1WorkflowExecuteMode `json:"mode,omitempty"`
	Ref string `json:"ref,omitempty"`
	Steps []V1alpha1WorkflowStep `json:"steps,omitempty"`
}
