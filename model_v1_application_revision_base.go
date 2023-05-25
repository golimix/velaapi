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
// V1ApplicationRevisionBase struct for V1ApplicationRevisionBase
type V1ApplicationRevisionBase struct {
	CodeInfo ModelCodeInfo `json:"codeInfo,omitempty"`
	CreateTime time.Time `json:"createTime"`
	DeployUser V1NameAlias `json:"deployUser,omitempty"`
	EnvName string `json:"envName"`
	ImageInfo ModelImageInfo `json:"imageInfo,omitempty"`
	Note string `json:"note"`
	Reason string `json:"reason,omitempty"`
	Status string `json:"status"`
	TriggerType string `json:"triggerType"`
	Version string `json:"version"`
}