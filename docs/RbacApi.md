# \RbacApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlatformPermission**](RbacApi.md#CreatePlatformPermission) | **Post** /api/v1/permissions | create the platform perm policy
[**CreatePlatformRole**](RbacApi.md#CreatePlatformRole) | **Post** /api/v1/roles | create platform level role
[**DeletePlatformPermission**](RbacApi.md#DeletePlatformPermission) | **Delete** /api/v1/permissions/{permissionName} | delete a platform perm policy
[**DeletePlatformRole**](RbacApi.md#DeletePlatformRole) | **Delete** /api/v1/roles/{roleName} | update platform level role
[**ListPlatformPermissions**](RbacApi.md#ListPlatformPermissions) | **Get** /api/v1/permissions | list all platform level perm policies
[**ListPlatformRoles**](RbacApi.md#ListPlatformRoles) | **Get** /api/v1/roles | list all platform level roles
[**UpdatePlatformRole**](RbacApi.md#UpdatePlatformRole) | **Put** /api/v1/roles/{roleName} | update platform level role



## CreatePlatformPermission

> V1PermissionBase CreatePlatformPermission(ctx, optional)

create the platform perm policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreatePlatformPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePlatformPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreatePermissionRequest** | [**optional.Interface of V1CreatePermissionRequest**](V1CreatePermissionRequest.md)|  | 

### Return type

[**V1PermissionBase**](V1PermissionBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlatformRole

> V1RoleBase CreatePlatformRole(ctx, optional)

create platform level role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreatePlatformRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePlatformRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateRoleRequest** | [**optional.Interface of V1CreateRoleRequest**](V1CreateRoleRequest.md)|  | 

### Return type

[**V1RoleBase**](V1RoleBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlatformPermission

> string DeletePlatformPermission(ctx, permissionName)

delete a platform perm policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionName** | **string**| identifier of the permission | 

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


## DeletePlatformRole

> string DeletePlatformRole(ctx, roleName)

update platform level role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string**| identifier of the role | 

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


## ListPlatformPermissions

> []V1PermissionBase ListPlatformPermissions(ctx, )

list all platform level perm policies

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]V1PermissionBase**](V1PermissionBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlatformRoles

> V1ListRolesResponse ListPlatformRoles(ctx, )

list all platform level roles

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1ListRolesResponse**](V1ListRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlatformRole

> V1RoleBase UpdatePlatformRole(ctx, roleName, optional)

update platform level role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string**| identifier of the role | 
 **optional** | ***UpdatePlatformRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePlatformRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1UpdateRoleRequest** | [**optional.Interface of V1UpdateRoleRequest**](V1UpdateRoleRequest.md)|  | 

### Return type

[**V1RoleBase**](V1RoleBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

