# \ClusterApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectCloudCluster**](ClusterApi.md#ConnectCloudCluster) | **Post** /api/v1/clusters/cloud_clusters/{provider}/connect | create cluster from cloud cluster
[**CreateCloudCluster**](ClusterApi.md#CreateCloudCluster) | **Post** /api/v1/clusters/cloud_clusters/{provider}/create | create cloud cluster
[**CreateKubeCluster**](ClusterApi.md#CreateKubeCluster) | **Post** /api/v1/clusters | create cluster
[**CreateNamespace**](ClusterApi.md#CreateNamespace) | **Post** /api/v1/clusters/{clusterName}/namespaces | create namespace in cluster
[**DeleteCloudClusterCreation**](ClusterApi.md#DeleteCloudClusterCreation) | **Delete** /api/v1/clusters/cloud_clusters/{provider}/creation/{cloudClusterName} | delete cloud cluster creation
[**DeleteKubeCluster**](ClusterApi.md#DeleteKubeCluster) | **Delete** /api/v1/clusters/{clusterName} | delete cluster
[**GetCloudClusterCreationStatus**](ClusterApi.md#GetCloudClusterCreationStatus) | **Get** /api/v1/clusters/cloud_clusters/{provider}/creation/{cloudClusterName} | check cloud cluster create status
[**GetKubeCluster**](ClusterApi.md#GetKubeCluster) | **Get** /api/v1/clusters/{clusterName} | detail cluster info
[**ListCloudClusterCreation**](ClusterApi.md#ListCloudClusterCreation) | **Get** /api/v1/clusters/cloud_clusters/{provider}/creation | list cloud cluster creation
[**ListCloudClusters**](ClusterApi.md#ListCloudClusters) | **Post** /api/v1/clusters/cloud_clusters/{provider} | list cloud clusters
[**ListKubeClusters**](ClusterApi.md#ListKubeClusters) | **Get** /api/v1/clusters | list all clusters
[**ModifyKubeCluster**](ClusterApi.md#ModifyKubeCluster) | **Put** /api/v1/clusters/{clusterName} | modify cluster



## ConnectCloudCluster

> V1ClusterBase ConnectCloudCluster(ctx, provider, optional)

create cluster from cloud cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string**| identifier of the cloud provider | 
 **optional** | ***ConnectCloudClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectCloudClusterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1ConnectCloudClusterRequest** | [**optional.Interface of V1ConnectCloudClusterRequest**](V1ConnectCloudClusterRequest.md)|  | 

### Return type

[**V1ClusterBase**](V1ClusterBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudCluster

> V1CreateCloudClusterResponse CreateCloudCluster(ctx, provider, optional)

create cloud cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string**| identifier of the cloud provider | 
 **optional** | ***CreateCloudClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCloudClusterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateCloudClusterRequest** | [**optional.Interface of V1CreateCloudClusterRequest**](V1CreateCloudClusterRequest.md)|  | 

### Return type

[**V1CreateCloudClusterResponse**](V1CreateCloudClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKubeCluster

> V1ClusterBase CreateKubeCluster(ctx, optional)

create cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateKubeClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKubeClusterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateClusterRequest** | [**optional.Interface of V1CreateClusterRequest**](V1CreateClusterRequest.md)|  | 

### Return type

[**V1ClusterBase**](V1ClusterBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNamespace

> V1CreateClusterNamespaceResponse CreateNamespace(ctx, clusterName, optional)

create namespace in cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**| name of the target cluster | 
 **optional** | ***CreateNamespaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNamespaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateClusterNamespaceRequest** | [**optional.Interface of V1CreateClusterNamespaceRequest**](V1CreateClusterNamespaceRequest.md)|  | 

### Return type

[**V1CreateClusterNamespaceResponse**](V1CreateClusterNamespaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudClusterCreation

> V1CreateCloudClusterResponse DeleteCloudClusterCreation(ctx, provider, cloudClusterName)

delete cloud cluster creation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string**| identifier of the cloud provider | 
**cloudClusterName** | **string**| identifier for cloud cluster which is creating | 

### Return type

[**V1CreateCloudClusterResponse**](V1CreateCloudClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKubeCluster

> V1ClusterBase DeleteKubeCluster(ctx, clusterName)

delete cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**| identifier of the cluster | 

### Return type

[**V1ClusterBase**](V1ClusterBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudClusterCreationStatus

> V1CreateCloudClusterResponse GetCloudClusterCreationStatus(ctx, provider, cloudClusterName)

check cloud cluster create status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string**| identifier of the cloud provider | 
**cloudClusterName** | **string**| identifier for cloud cluster which is creating | 

### Return type

[**V1CreateCloudClusterResponse**](V1CreateCloudClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubeCluster

> V1DetailClusterResponse GetKubeCluster(ctx, clusterName)

detail cluster info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**| identifier of the cluster | 

### Return type

[**V1DetailClusterResponse**](V1DetailClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudClusterCreation

> V1ListCloudClusterCreationResponse ListCloudClusterCreation(ctx, provider)

list cloud cluster creation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string**| identifier of the cloud provider | 

### Return type

[**V1ListCloudClusterCreationResponse**](V1ListCloudClusterCreationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudClusters

> V1ListCloudClusterResponse ListCloudClusters(ctx, provider, optional)

list cloud clusters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string**| identifier of the cloud provider | 
 **optional** | ***ListCloudClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCloudClustersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page for paging | 
 **pageSize** | **optional.Int32**| PageSize for paging | 
 **v1AccessKeyRequest** | [**optional.Interface of V1AccessKeyRequest**](V1AccessKeyRequest.md)|  | 

### Return type

[**V1ListCloudClusterResponse**](V1ListCloudClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubeClusters

> V1SimpleResponse ListKubeClusters(ctx, optional)

list all clusters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListKubeClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKubeClustersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Fuzzy search based on name or description | 
 **page** | **optional.Int32**| Page for paging | 
 **pageSize** | **optional.Int32**| PageSize for paging | 

### Return type

[**V1SimpleResponse**](V1SimpleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyKubeCluster

> V1ClusterBase ModifyKubeCluster(ctx, clusterName, optional)

modify cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**| identifier of the cluster | 
 **optional** | ***ModifyKubeClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyKubeClusterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateClusterRequest** | [**optional.Interface of V1CreateClusterRequest**](V1CreateClusterRequest.md)|  | 

### Return type

[**V1ClusterBase**](V1ClusterBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

