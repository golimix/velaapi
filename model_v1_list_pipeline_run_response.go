/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1ListPipelineRunResponse struct for V1ListPipelineRunResponse
type V1ListPipelineRunResponse struct {
	Runs []V1PipelineRunBriefing `json:"runs"`
	Total int64 `json:"total"`
}
