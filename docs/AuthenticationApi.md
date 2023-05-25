# \AuthenticationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDexConfig**](AuthenticationApi.md#GetDexConfig) | **Get** /api/v1/auth/dex_config | get Dex config
[**GetLoginType**](AuthenticationApi.md#GetLoginType) | **Get** /api/v1/auth/login_type | get login type
[**GetLoginUserInfo**](AuthenticationApi.md#GetLoginUserInfo) | **Get** /api/v1/auth/user_info | get login user detail info
[**Login**](AuthenticationApi.md#Login) | **Post** /api/v1/auth/login | handle login request
[**RefreshToken**](AuthenticationApi.md#RefreshToken) | **Get** /api/v1/auth/refresh_token | refresh token



## GetDexConfig

> V1DexConfigResponse GetDexConfig(ctx, )

get Dex config

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1DexConfigResponse**](V1DexConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginType

> V1GetLoginTypeResponse GetLoginType(ctx, )

get login type

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1GetLoginTypeResponse**](V1GetLoginTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginUserInfo

> V1LoginUserInfoResponse GetLoginUserInfo(ctx, )

get login user detail info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1LoginUserInfoResponse**](V1LoginUserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> V1LoginResponse Login(ctx, optional)

handle login request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoginOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1LoginRequest** | [**optional.Interface of V1LoginRequest**](V1LoginRequest.md)|  | 

### Return type

[**V1LoginResponse**](V1LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshToken

> V1RefreshTokenResponse RefreshToken(ctx, )

refresh token

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1RefreshTokenResponse**](V1RefreshTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

