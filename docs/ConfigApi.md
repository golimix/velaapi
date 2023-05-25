# \ConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig**](ConfigApi.md#CreateConfig) | **Post** /api/v1/configs | create or update a config
[**DeleteConfig**](ConfigApi.md#DeleteConfig) | **Delete** /api/v1/configs/{configName} | delete a config
[**GetConfig**](ConfigApi.md#GetConfig) | **Get** /api/v1/configs/{configName} | detail a config
[**GetConfigTemplate**](ConfigApi.md#GetConfigTemplate) | **Get** /api/v1/config_templates/{templateName} | Detail a template
[**GetConfigs**](ConfigApi.md#GetConfigs) | **Get** /api/v1/configs | list all configs that belong to the system scope
[**ListConfigTemplates**](ConfigApi.md#ListConfigTemplates) | **Get** /api/v1/config_templates | List all config templates from the system namespace
[**UpdateConfig**](ConfigApi.md#UpdateConfig) | **Put** /api/v1/configs/{configName} | update a config



## CreateConfig

> V1Config CreateConfig(ctx, optional)

create or update a config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateConfigRequest** | [**optional.Interface of V1CreateConfigRequest**](V1CreateConfigRequest.md)|  | 

### Return type

[**V1Config**](V1Config.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfig

> string DeleteConfig(ctx, configName)

delete a config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configName** | **string**| identifier of the config | 

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


## GetConfig

> []string GetConfig(ctx, configName)

detail a config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configName** | **string**| identifier of the config | 

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


## GetConfigTemplate

> V1ConfigTemplateDetail GetConfigTemplate(ctx, templateName, optional)

Detail a template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string**| identifier of the config template | 
 **optional** | ***GetConfigTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConfigTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**| the name of the namespace | 

### Return type

[**V1ConfigTemplateDetail**](V1ConfigTemplateDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigs

> V1ListConfigResponse GetConfigs(ctx, optional)

list all configs that belong to the system scope

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | **optional.String**| the name of the template | 

### Return type

[**V1ListConfigResponse**](V1ListConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigTemplates

> V1ListConfigTemplateResponse ListConfigTemplates(ctx, )

List all config templates from the system namespace

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1ListConfigTemplateResponse**](V1ListConfigTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfig

> []string UpdateConfig(ctx, configName)

update a config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configName** | **string**| identifier of the config | 

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

