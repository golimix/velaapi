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
// V1DetailRevisionResponse struct for V1DetailRevisionResponse
type V1DetailRevisionResponse struct {
	AppPrimaryKey string `json:"appPrimaryKey"`
	ApplyAppConfig string `json:"applyAppConfig,omitempty"`
	CodeInfo ModelCodeInfo `json:"codeInfo,omitempty"`
	CreateTime time.Time `json:"createTime"`
	DeployUser V1NameAlias `json:"deployUser"`
	EnvName string `json:"envName"`
	ImageInfo ModelImageInfo `json:"imageInfo,omitempty"`
	Note string `json:"note"`
	Reason string `json:"reason"`
	RevisionCRName string `json:"revisionCRName"`
	RollbackVersion string `json:"rollbackVersion,omitempty"`
	Status string `json:"status"`
	TriggerType string `json:"triggerType"`
	UpdateTime time.Time `json:"updateTime"`
	Version string `json:"version"`
	WorkflowName string `json:"workflowName"`
}