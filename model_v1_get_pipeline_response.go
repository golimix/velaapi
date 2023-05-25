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
// V1GetPipelineResponse struct for V1GetPipelineResponse
type V1GetPipelineResponse struct {
	Alias string `json:"alias"`
	CreateTime time.Time `json:"createTime"`
	Description string `json:"description"`
	Info V1PipelineInfo `json:"info"`
	Name string `json:"name"`
	Project V1NameAlias `json:"project"`
	Spec ModelWorkflowSpec `json:"spec"`
}
