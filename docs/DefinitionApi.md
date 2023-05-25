# \DefinitionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DetailDefinition**](DefinitionApi.md#DetailDefinition) | **Get** /api/v1/definitions/{definitionName} | Detail a definition
[**ListDefinitions**](DefinitionApi.md#ListDefinitions) | **Get** /api/v1/definitions | list all definitions
[**UpdateDefinitionStatus**](DefinitionApi.md#UpdateDefinitionStatus) | **Put** /api/v1/definitions/{definitionName}/status | Update the status for a definition
[**UpdateUISchema**](DefinitionApi.md#UpdateUISchema) | **Put** /api/v1/definitions/{definitionName}/uischema | Update the UI schema for a definition



## DetailDefinition

> V1SimpleResponse DetailDefinition(ctx, definitionName, optional)

Detail a definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionName** | **string**| identifier of the definition | 
 **optional** | ***DetailDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DetailDefinitionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| query the definition type | 

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


## ListDefinitions

> V1SimpleResponse ListDefinitions(ctx, type_, optional)

list all definitions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| query the definition type | 
 **optional** | ***ListDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDefinitionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryAll** | **optional.String**| query all definitions include hidden in UI | 
 **appliedWorkload** | **optional.String**| if specified, query the trait definition applied to the workload | 
 **ownerAddon** | **optional.String**| query by which addon created the definition | 
 **scope** | **optional.String**| query by the specified scope like WorkflowRun or Application | 

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


## UpdateDefinitionStatus

> V1SimpleResponse UpdateDefinitionStatus(ctx, definitionName, optional)

Update the status for a definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionName** | **string**|  | 
 **optional** | ***UpdateDefinitionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDefinitionStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1UpdateDefinitionStatusRequest** | [**optional.Interface of V1UpdateDefinitionStatusRequest**](V1UpdateDefinitionStatusRequest.md)|  | 

### Return type

[**V1SimpleResponse**](V1SimpleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUISchema

> V1SimpleResponse UpdateUISchema(ctx, definitionName, optional)

Update the UI schema for a definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionName** | **string**|  | 
 **optional** | ***UpdateUISchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUISchemaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1UpdateUiSchemaRequest** | [**optional.Interface of V1UpdateUiSchemaRequest**](V1UpdateUiSchemaRequest.md)|  | 

### Return type

[**V1SimpleResponse**](V1SimpleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

