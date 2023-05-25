# \OamApplicationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateApplication**](OamApplicationApi.md#CreateOrUpdateApplication) | **Post** /v1/namespaces/{namespace}/applications/{appname} | create or update oam application in the specified namespace
[**DeleteOAMApplication**](OamApplicationApi.md#DeleteOAMApplication) | **Delete** /v1/namespaces/{namespace}/applications/{appname} | create or update oam application in the specified namespace
[**GetApplication**](OamApplicationApi.md#GetApplication) | **Get** /v1/namespaces/{namespace}/applications/{appname} | get the specified oam application in the specified namespace



## CreateOrUpdateApplication

> map[string]interface{} CreateOrUpdateApplication(ctx, namespace, appname, optional)

create or update oam application in the specified namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string**| identifier of the namespace | 
**appname** | **string**| identifier of the oam application | 
 **optional** | ***CreateOrUpdateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dryRun** | **optional.String**| When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | 
 **v1ApplicationRequest** | [**optional.Interface of V1ApplicationRequest**](V1ApplicationRequest.md)|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAMApplication

> map[string]interface{} DeleteOAMApplication(ctx, namespace, appname)

create or update oam application in the specified namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string**| identifier of the namespace | 
**appname** | **string**| identifier of the oam application | 

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


## GetApplication

> V1ApplicationResponse GetApplication(ctx, namespace, appname)

get the specified oam application in the specified namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string**| identifier of the namespace | 
**appname** | **string**| identifier of the oam application | 

### Return type

[**V1ApplicationResponse**](V1ApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

