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
// V1DetailUserResponse struct for V1DetailUserResponse
type V1DetailUserResponse struct {
	Alias string `json:"alias,omitempty"`
	CreateTime time.Time `json:"createTime"`
	Disabled bool `json:"disabled"`
	Email string `json:"email"`
	LastLoginTime time.Time `json:"lastLoginTime"`
	Name string `json:"name"`
	Projects []V1ProjectBase `json:"projects"`
	Roles []V1NameAlias `json:"roles"`
}
