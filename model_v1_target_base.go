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
// V1TargetBase struct for V1TargetBase
type V1TargetBase struct {
	Alias string `json:"alias,omitempty"`
	AppNum int64 `json:"appNum,omitempty"`
	Cluster V1ClusterTarget `json:"cluster,omitempty"`
	ClusterAlias string `json:"clusterAlias,omitempty"`
	CreateTime time.Time `json:"createTime"`
	Description string `json:"description,omitempty"`
	Name string `json:"name"`
	Project V1NameAlias `json:"project"`
	UpdateTime time.Time `json:"updateTime"`
	Variable map[string]interface{} `json:"variable,omitempty"`
}