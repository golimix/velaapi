/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// ModelWorkflowSpec struct for ModelWorkflowSpec
type ModelWorkflowSpec struct {
	Mode V1alpha1WorkflowExecuteMode `json:"mode,omitempty"`
	Steps []ModelWorkflowStep `json:"steps,omitempty"`
}