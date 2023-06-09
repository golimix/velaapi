/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1AddonRegistry struct for V1AddonRegistry
type V1AddonRegistry struct {
	Git AddonGitAddonSource `json:"git,omitempty"`
	Gitee AddonGiteeAddonSource `json:"gitee,omitempty"`
	Gitlab AddonGitlabAddonSource `json:"gitlab,omitempty"`
	Helm AddonHelmSource `json:"helm,omitempty"`
	Name string `json:"name"`
	Oss AddonOssAddonSource `json:"oss,omitempty"`
}
