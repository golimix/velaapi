/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// ModelProviderInfo struct for ModelProviderInfo
type ModelProviderInfo struct {
	ClusterID string `json:"clusterID"`
	ClusterName string `json:"clusterName,omitempty"`
	Labels map[string]string `json:"labels"`
	Provider string `json:"provider"`
	RegionID string `json:"regionID,omitempty"`
	VpcID string `json:"vpcID,omitempty"`
	Zone string `json:"zone,omitempty"`
	ZoneID string `json:"zoneID,omitempty"`
}
