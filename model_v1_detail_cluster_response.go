/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
import (
	"time"
)
// V1DetailClusterResponse struct for V1DetailClusterResponse
type V1DetailClusterResponse struct {
	Alias string `json:"alias"`
	ApiServerURL string `json:"apiServerURL"`
	CreateTime time.Time `json:"createTime"`
	DashboardURL string `json:"dashboardURL"`
	Description string `json:"description"`
	Icon string `json:"icon"`
	KubeConfig string `json:"kubeConfig"`
	KubeConfigSecret string `json:"kubeConfigSecret"`
	Labels map[string]string `json:"labels"`
	Name string `json:"name"`
	Provider ModelProviderInfo `json:"provider"`
	Reason string `json:"reason"`
	ResourceInfo V1ClusterResourceInfo `json:"resourceInfo"`
	Status string `json:"status"`
	UpdateTime time.Time `json:"updateTime"`
}