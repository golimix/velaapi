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
// V1RoleBase struct for V1RoleBase
type V1RoleBase struct {
	Alias string `json:"alias,omitempty"`
	CreateTime time.Time `json:"createTime"`
	Name string `json:"name"`
	Permissions []V1NameAlias `json:"permissions"`
	UpdateTime time.Time `json:"updateTime"`
}
