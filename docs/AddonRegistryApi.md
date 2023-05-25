# \AddonRegistryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddonRegistry**](AddonRegistryApi.md#CreateAddonRegistry) | **Post** /api/v1/addon_registries | create an addon registry
[**DeleteAddonRegistry**](AddonRegistryApi.md#DeleteAddonRegistry) | **Delete** /api/v1/addon_registries/{addonRegName} | delete an addon registry
[**ListAddonRegistry**](AddonRegistryApi.md#ListAddonRegistry) | **Get** /api/v1/addon_registries | list all addon registry
[**UpdateAddonRegistry**](AddonRegistryApi.md#UpdateAddonRegistry) | **Put** /api/v1/addon_registries/{addonRegName} | update an addon registry



## CreateAddonRegistry

> V1AddonRegistry CreateAddonRegistry(ctx, optional)

create an addon registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAddonRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAddonRegistryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateAddonRegistryRequest** | [**optional.Interface of V1CreateAddonRegistryRequest**](V1CreateAddonRegistryRequest.md)|  | 

### Return type

[**V1AddonRegistry**](V1AddonRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddonRegistry

> V1AddonRegistry DeleteAddonRegistry(ctx, addonRegName)

delete an addon registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonRegName** | **string**| identifier of the addon registry | 

### Return type

[**V1AddonRegistry**](V1AddonRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddonRegistry

> V1ListAddonRegistryResponse ListAddonRegistry(ctx, )

list all addon registry

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1ListAddonRegistryResponse**](V1ListAddonRegistryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddonRegistry

> V1AddonRegistry UpdateAddonRegistry(ctx, addonRegName, optional)

update an addon registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonRegName** | **string**| identifier of the addon registry | 
 **optional** | ***UpdateAddonRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAddonRegistryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1UpdateAddonRegistryRequest** | [**optional.Interface of V1UpdateAddonRegistryRequest**](V1UpdateAddonRegistryRequest.md)|  | 

### Return type

[**V1AddonRegistry**](V1AddonRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

