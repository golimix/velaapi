/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1PipelineInfo struct for V1PipelineInfo
type V1PipelineInfo struct {
	LastRun V1PipelineRun `json:"lastRun"`
	RunStat V1RunStat `json:"runStat"`
}
