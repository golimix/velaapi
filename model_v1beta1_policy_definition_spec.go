/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1beta1PolicyDefinitionSpec struct for V1beta1PolicyDefinitionSpec
type V1beta1PolicyDefinitionSpec struct {
	DefinitionRef CommonDefinitionReference `json:"definitionRef,omitempty"`
	ManageHealthCheck bool `json:"manageHealthCheck,omitempty"`
	Schematic CommonSchematic `json:"schematic,omitempty"`
}
