/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi

// ModelWorkflowStep struct for ModelWorkflowStep
type ModelWorkflowStep struct {
	Alias       string                   `json:"alias"`
	DependsOn   []string                 `json:"dependsOn"`
	Description string                   `json:"description"`
	If          string                   `json:"if,omitempty"`
	Inputs      []V1alpha1InputItem      `json:"inputs,omitempty"`
	Meta        V1alpha1WorkflowStepMeta `json:"meta,omitempty"`
	Name        string                   `json:"name"`
	OrderIndex  int32                    `json:"orderIndex"`
	Outputs     []V1alpha1OutputItem     `json:"outputs,omitempty"`
	Properties  map[string]interface{}   `json:"properties,omitempty"`
	SubSteps    []ModelWorkflowStepBase  `json:"subSteps,omitempty"`
	Timeout     string                   `json:"timeout,omitempty"`
	Type        string                   `json:"type"`
}
