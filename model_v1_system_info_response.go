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
// V1SystemInfoResponse struct for V1SystemInfoResponse
type V1SystemInfoResponse struct {
	DexUserDefaultPlatformRoles []string `json:"dexUserDefaultPlatformRoles,omitempty"`
	DexUserDefaultProjects []ModelProjectRef `json:"dexUserDefaultProjects,omitempty"`
	EnableCollection bool `json:"enableCollection"`
	InstallTime time.Time `json:"installTime,omitempty"`
	LoginType string `json:"loginType"`
	PlatformID string `json:"platformID"`
	StatisticInfo V1StatisticInfo `json:"statisticInfo,omitempty"`
	SystemVersion V1SystemVersion `json:"systemVersion"`
}
