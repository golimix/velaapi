# \CloudshellApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DestroyCloudShell**](CloudshellApi.md#DestroyCloudShell) | **Delete** /api/v1/cloudshell | destroy the user&#39;s cloud shell environment
[**PrepareCloudShell**](CloudshellApi.md#PrepareCloudShell) | **Post** /api/v1/cloudshell | prepare the user&#39;s cloud shell environment
[**Proxy**](CloudshellApi.md#Proxy) | **Get** /view/cloudshell/{subpath} | prepare the user&#39;s cloud shell environment
[**Proxy_0**](CloudshellApi.md#Proxy_0) | **Get** /view/cloudshell | prepare the user&#39;s cloud shell environment



## DestroyCloudShell

> V1SimpleResponse DestroyCloudShell(ctx, )

destroy the user's cloud shell environment

### Required Parameters

This endpoint does not need any parameter.

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


## PrepareCloudShell

> V1SimpleResponse PrepareCloudShell(ctx, )

prepare the user's cloud shell environment

### Required Parameters

This endpoint does not need any parameter.

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


## Proxy

> V1SimpleResponse Proxy(ctx, subpath)

prepare the user's cloud shell environment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subpath** | **string**|  | 

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


## Proxy_0

> V1SimpleResponse Proxy_0(ctx, )

prepare the user's cloud shell environment

### Required Parameters

This endpoint does not need any parameter.

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

