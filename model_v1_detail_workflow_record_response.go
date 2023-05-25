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
// V1DetailWorkflowRecordResponse struct for V1DetailWorkflowRecordResponse
type V1DetailWorkflowRecordResponse struct {
	ApplicationRevision string `json:"applicationRevision"`
	DeployTime time.Time `json:"deployTime"`
	DeployUser string `json:"deployUser"`
	EndTime time.Time `json:"endTime,omitempty"`
	Message string `json:"message"`
	Mode string `json:"mode"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Note string `json:"note"`
	StartTime time.Time `json:"startTime,omitempty"`
	Status string `json:"status"`
	Steps []ModelWorkflowStepStatus `json:"steps,omitempty"`
	TriggerType string `json:"triggerType"`
	WorkflowAlias string `json:"workflowAlias"`
	WorkflowName string `json:"workflowName"`
}