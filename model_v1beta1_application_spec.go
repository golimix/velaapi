/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1beta1ApplicationSpec struct for V1beta1ApplicationSpec
type V1beta1ApplicationSpec struct {
	Components []CommonApplicationComponent `json:"components"`
	Policies []V1beta1AppPolicy `json:"policies,omitempty"`
	Workflow V1beta1Workflow `json:"workflow,omitempty"`
}