/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1CreateWorkflowRequest struct for V1CreateWorkflowRequest
type V1CreateWorkflowRequest struct {
	Alias string `json:"alias,omitempty"`
	Default bool `json:"default"`
	Description string `json:"description,omitempty"`
	EnvName string `json:"envName"`
	Mode string `json:"mode"`
	Name string `json:"name"`
	Steps []V1WorkflowStep `json:"steps,omitempty"`
	SubMode string `json:"subMode"`
}
