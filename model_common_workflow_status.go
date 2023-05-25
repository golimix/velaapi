/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// CommonWorkflowStatus struct for CommonWorkflowStatus
type CommonWorkflowStatus struct {
	AppRevision string `json:"appRevision,omitempty"`
	ContextBackend V1ObjectReference `json:"contextBackend,omitempty"`
	EndTime string `json:"endTime,omitempty"`
	Finished bool `json:"finished"`
	Message string `json:"message,omitempty"`
	Mode string `json:"mode"`
	StartTime string `json:"startTime,omitempty"`
	Status string `json:"status,omitempty"`
	Steps []V1alpha1WorkflowStepStatus `json:"steps,omitempty"`
	Suspend bool `json:"suspend"`
	SuspendState string `json:"suspendState,omitempty"`
	Terminated bool `json:"terminated"`
}
