/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1PipelineRunMeta struct for V1PipelineRunMeta
type V1PipelineRunMeta struct {
	PipelineName string `json:"pipelineName"`
	PipelineRunName string `json:"pipelineRunName"`
	Project V1NameAlias `json:"project"`
}
