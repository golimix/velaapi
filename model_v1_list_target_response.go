/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1ListTargetResponse struct for V1ListTargetResponse
type V1ListTargetResponse struct {
	Targets []V1TargetBase `json:"targets"`
	Total int64 `json:"total"`
}
