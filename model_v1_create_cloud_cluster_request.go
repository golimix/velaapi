/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1CreateCloudClusterRequest struct for V1CreateCloudClusterRequest
type V1CreateCloudClusterRequest struct {
	AccessKeyID string `json:"accessKeyID"`
	AccessKeySecret string `json:"accessKeySecret"`
	CpuCoresPerWorker int64 `json:"cpuCoresPerWorker"`
	MemoryPerWorker int64 `json:"memoryPerWorker"`
	Name string `json:"name"`
	WorkerNumber int32 `json:"workerNumber"`
	Zone string `json:"zone"`
}
