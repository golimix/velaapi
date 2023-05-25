# \RepositoryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartValues**](RepositoryApi.md#ChartValues) | **Get** /api/v1/repository/charts/{chart}/versions/{version}/values | get chart value
[**GetImageInfo**](RepositoryApi.md#GetImageInfo) | **Get** /api/v1/repository/image/info | get the oci repos
[**GetImageRepos**](RepositoryApi.md#GetImageRepos) | **Get** /api/v1/repository/image/repos | get the oci repos
[**ListCharts**](RepositoryApi.md#ListCharts) | **Get** /api/v1/repository/charts | list charts
[**ListRepo**](RepositoryApi.md#ListRepo) | **Get** /api/v1/repository/chart_repos | list chart repo
[**ListVersions**](RepositoryApi.md#ListVersions) | **Get** /api/v1/repository/charts/{chart}/versions | list versions



## ChartValues

> map[string]interface{} ChartValues(ctx, chart, version, optional)

get chart value

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chart** | **string**|  | 
**version** | **string**|  | 
 **optional** | ***ChartValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChartValuesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **repoUrl** | **optional.String**| helm repository url | 
 **secretName** | **optional.String**| secret of the repo | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageInfo

> V1ImageInfo GetImageInfo(ctx, project, name, optional)

get the oci repos

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string**| the config project | 
**name** | **string**| the image name | 
 **optional** | ***GetImageInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secretName** | **optional.String**| the secret name of the image repository | 

### Return type

[**V1ImageInfo**](V1ImageInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageRepos

> V1ListImageRegistryResponse GetImageRepos(ctx, project)

get the oci repos

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string**| the config project | 

### Return type

[**V1ListImageRegistryResponse**](V1ListImageRegistryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCharts

> []string ListCharts(ctx, optional)

list charts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListChartsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListChartsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoUrl** | **optional.String**| helm repository url | 
 **secretName** | **optional.String**| secret of the repo | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepo

> []string ListRepo(ctx, project)

list chart repo

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string**| the config project | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersions

> V1ChartVersionListResponse ListVersions(ctx, chart, optional)

list versions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chart** | **string**|  | 
 **optional** | ***ListVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repoUrl** | **optional.String**| helm repository url | 
 **secretName** | **optional.String**| secret of the repo | 

### Return type

[**V1ChartVersionListResponse**](V1ChartVersionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

