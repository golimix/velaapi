# \AddonApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DetailAddon**](AddonApi.md#DetailAddon) | **Get** /api/v1/addons/{addonName} | show details of an addon
[**DisableAddon**](AddonApi.md#DisableAddon) | **Post** /api/v1/addons/{addonName}/disable | disable an addon
[**EnableAddon**](AddonApi.md#EnableAddon) | **Post** /api/v1/addons/{addonName}/enable | enable an addon
[**List**](AddonApi.md#List) | **Get** /api/v1/enabled_addon | list all enabled addons
[**ListAddons**](AddonApi.md#ListAddons) | **Get** /api/v1/addons | list all addons
[**StatusAddon**](AddonApi.md#StatusAddon) | **Get** /api/v1/addons/{addonName}/status | show status of an addon
[**UpdateAddon**](AddonApi.md#UpdateAddon) | **Put** /api/v1/addons/{addonName}/update | update an addon



## DetailAddon

> V1DetailAddonResponse DetailAddon(ctx, name, addonName, optional)

show details of an addon

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| addon name to query detail | 
**addonName** | **string**| addon name to query detail | 
 **optional** | ***DetailAddonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DetailAddonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| specify addon version to enable | 
 **registry** | **optional.String**| filter addons from given registry | 

### Return type

[**V1DetailAddonResponse**](V1DetailAddonResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableAddon

> V1AddonStatusResponse DisableAddon(ctx, addonName, optional)

disable an addon

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonName** | **string**| addon name to enable | 
 **optional** | ***DisableAddonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DisableAddonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.String**| force disable an addon | 

### Return type

[**V1AddonStatusResponse**](V1AddonStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableAddon

> V1AddonStatusResponse EnableAddon(ctx, addonName, optional)

enable an addon

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonName** | **string**| addon name to enable | 
 **optional** | ***EnableAddonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnableAddonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1EnableAddonRequest** | [**optional.Interface of V1EnableAddonRequest**](V1EnableAddonRequest.md)|  | 

### Return type

[**V1AddonStatusResponse**](V1AddonStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> V1ListEnabledAddonResponse List(ctx, optional)

list all enabled addons

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registry** | **optional.String**| filter addons from given registry | 
 **query** | **optional.String**| Fuzzy search based on name and description. | 

### Return type

[**V1ListEnabledAddonResponse**](V1ListEnabledAddonResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddons

> V1ListAddonResponse ListAddons(ctx, optional)

list all addons

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAddonsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAddonsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registry** | **optional.String**| filter addons from given registry | 
 **query** | **optional.String**| Fuzzy search based on name and description. | 

### Return type

[**V1ListAddonResponse**](V1ListAddonResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusAddon

> V1AddonStatusResponse StatusAddon(ctx, addonName)

show status of an addon

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonName** | **string**| addon name to query status | 

### Return type

[**V1AddonStatusResponse**](V1AddonStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddon

> V1AddonStatusResponse UpdateAddon(ctx, addonName, optional)

update an addon

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonName** | **string**| addon name to update | 
 **optional** | ***UpdateAddonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAddonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1EnableAddonRequest** | [**optional.Interface of V1EnableAddonRequest**](V1EnableAddonRequest.md)|  | 

### Return type

[**V1AddonStatusResponse**](V1AddonStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

