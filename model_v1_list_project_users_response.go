/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1ListProjectUsersResponse struct for V1ListProjectUsersResponse
type V1ListProjectUsersResponse struct {
	Total int64 `json:"total"`
	Users []V1ProjectUserBase `json:"users"`
}