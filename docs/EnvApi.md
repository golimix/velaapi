# \EnvApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Envcreate**](EnvApi.md#Envcreate) | **Post** /api/v1/envs | create an env
[**Envdelete**](EnvApi.md#Envdelete) | **Delete** /api/v1/envs/{envName} | delete one env
[**Envlist**](EnvApi.md#Envlist) | **Get** /api/v1/envs | list all envs
[**Envupdate**](EnvApi.md#Envupdate) | **Put** /api/v1/envs/{envName} | update an env



## Envcreate

> V1Env Envcreate(ctx, optional)

create an env

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnvcreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnvcreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateEnvRequest** | [**optional.Interface of V1CreateEnvRequest**](V1CreateEnvRequest.md)|  | 

### Return type

[**V1Env**](V1Env.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Envdelete

> string Envdelete(ctx, envName)

delete one env

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| identifier of the environment | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Envlist

> V1ListEnvResponse Envlist(ctx, )

list all envs

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1ListEnvResponse**](V1ListEnvResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Envupdate

> V1Env Envupdate(ctx, envName, optional)

update an env

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| identifier of the environment | 
 **optional** | ***EnvupdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnvupdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateEnvRequest** | [**optional.Interface of V1CreateEnvRequest**](V1CreateEnvRequest.md)|  | 

### Return type

[**V1Env**](V1Env.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

