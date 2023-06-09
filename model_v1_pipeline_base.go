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
// V1PipelineBase struct for V1PipelineBase
type V1PipelineBase struct {
	Alias string `json:"alias"`
	CreateTime time.Time `json:"createTime"`
	Description string `json:"description"`
	Name string `json:"name"`
	Project V1NameAlias `json:"project"`
	Spec ModelWorkflowSpec `json:"spec"`
}
