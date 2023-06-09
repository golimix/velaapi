/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1DetailDefinitionResponse struct for V1DetailDefinitionResponse
type V1DetailDefinitionResponse struct {
	Alias string `json:"alias"`
	Component V1beta1ComponentDefinitionSpec `json:"component,omitempty"`
	Description string `json:"description"`
	Icon string `json:"icon"`
	Labels map[string]string `json:"labels"`
	Name string `json:"name"`
	OwnerAddon string `json:"ownerAddon"`
	Policy V1beta1PolicyDefinitionSpec `json:"policy,omitempty"`
	Schema string `json:"schema"`
	Status string `json:"status"`
	Trait V1beta1TraitDefinitionSpec `json:"trait,omitempty"`
	UiSchema []UtilsUiParameter `json:"uiSchema"`
	WorkflowStep V1beta1WorkflowStepDefinitionSpec `json:"workflowStep,omitempty"`
	WorkloadType string `json:"workloadType,omitempty"`
}
