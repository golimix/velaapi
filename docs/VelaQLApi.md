# \VelaQLApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryView**](VelaQLApi.md#QueryView) | **Get** /api/v1/query | use velaQL to query resource status



## QueryView

> map[string]interface{} QueryView(ctx, optional)

use velaQL to query resource status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryViewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryViewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **velaql** | **optional.String**| velaql query statement | 

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

