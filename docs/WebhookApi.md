# \WebhookApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleApplicationWebhook**](WebhookApi.md#HandleApplicationWebhook) | **Post** /api/v1/webhook/{token} | handle application webhook request



## HandleApplicationWebhook

> V1ApplicationDeployResponse HandleApplicationWebhook(ctx, token, optional)

handle application webhook request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**| webhook token | 
 **optional** | ***HandleApplicationWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HandleApplicationWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1HandleApplicationTriggerWebhookRequest** | [**optional.Interface of V1HandleApplicationTriggerWebhookRequest**](V1HandleApplicationTriggerWebhookRequest.md)|  | 

### Return type

[**V1ApplicationDeployResponse**](V1ApplicationDeployResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

