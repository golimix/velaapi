/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1CreateConfigDistributionRequest struct for V1CreateConfigDistributionRequest
type V1CreateConfigDistributionRequest struct {
	Configs []V1NamespacedName `json:"configs"`
	Name string `json:"name"`
	Targets []V1ClusterTarget `json:"targets"`
}
