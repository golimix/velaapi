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
// V1ProjectBase struct for V1ProjectBase
type V1ProjectBase struct {
	Alias string `json:"alias"`
	CreateTime time.Time `json:"createTime"`
	Description string `json:"description"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Owner V1NameAlias `json:"owner,omitempty"`
	UpdateTime time.Time `json:"updateTime"`
}
