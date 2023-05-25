/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1WorkflowStep struct for V1WorkflowStep
type V1WorkflowStep struct {
	Alias string `json:"alias,omitempty"`
	DependsOn []string `json:"dependsOn,omitempty"`
	Description string `json:"description,omitempty"`
	If string `json:"if,omitempty"`
	Inputs []V1alpha1InputItem `json:"inputs,omitempty"`
	Meta V1alpha1WorkflowStepMeta `json:"meta,omitempty"`
	Name string `json:"name"`
	Outputs []V1alpha1OutputItem `json:"outputs,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	SubSteps []V1WorkflowStepBase `json:"subSteps,omitempty"`
	Timeout string `json:"timeout,omitempty"`
	Type string `json:"type"`
}
