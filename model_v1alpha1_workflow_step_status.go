/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1alpha1WorkflowStepStatus struct for V1alpha1WorkflowStepStatus
type V1alpha1WorkflowStepStatus struct {
	FirstExecuteTime string `json:"firstExecuteTime,omitempty"`
	Id string `json:"id"`
	LastExecuteTime string `json:"lastExecuteTime,omitempty"`
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Phase string `json:"phase,omitempty"`
	Reason string `json:"reason,omitempty"`
	SubSteps []V1alpha1StepStatus `json:"subSteps,omitempty"`
	Type string `json:"type,omitempty"`
}