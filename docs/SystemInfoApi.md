# \SystemInfoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemInfo**](SystemInfoApi.md#GetSystemInfo) | **Get** /api/v1/system_info | /api/v1/system_info
[**UpdateSystemInfo**](SystemInfoApi.md#UpdateSystemInfo) | **Put** /api/v1/system_info | /api/v1/system_info



## GetSystemInfo

> V1SystemInfoResponse GetSystemInfo(ctx, )

/api/v1/system_info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1SystemInfoResponse**](V1SystemInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSystemInfo

> V1SystemInfoResponse UpdateSystemInfo(ctx, optional)

/api/v1/system_info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateSystemInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSystemInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemInfoRequest** | [**optional.Interface of V1SystemInfoRequest**](V1SystemInfoRequest.md)|  | 

### Return type

[**V1SystemInfoResponse**](V1SystemInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

