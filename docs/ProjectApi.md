# \ProjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyDistribution**](ProjectApi.md#ApplyDistribution) | **Post** /api/v1/projects/{projectName}/distributions | apply the distribution job of the config
[**CreateConfig**](ProjectApi.md#CreateConfig) | **Post** /api/v1/projects/{projectName}/configs | create a config in a project
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /api/v1/projects | create a project
[**CreateProjectPermission**](ProjectApi.md#CreateProjectPermission) | **Post** /api/v1/projects/{projectName}/permissions | create a project level perm policy
[**CreateProjectRole**](ProjectApi.md#CreateProjectRole) | **Post** /api/v1/projects/{projectName}/roles | create project level role
[**CreateProjectUser**](ProjectApi.md#CreateProjectUser) | **Post** /api/v1/projects/{projectName}/users | add a user to a project
[**DeleteConfig**](ProjectApi.md#DeleteConfig) | **Delete** /api/v1/projects/{projectName}/configs/{configName} | delete a config from a project
[**DeleteDistribution**](ProjectApi.md#DeleteDistribution) | **Delete** /api/v1/projects/{projectName}/distributions/{distributionName} | delete a distribution job of the config
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /api/v1/projects/{projectName} | delete a project
[**DeleteProjectPermission**](ProjectApi.md#DeleteProjectPermission) | **Delete** /api/v1/projects/{projectName}/permissions/{permissionName} | delete a project level perm policy
[**DeleteProjectRole**](ProjectApi.md#DeleteProjectRole) | **Delete** /api/v1/projects/{projectName}/roles/{roleName} | delete project level role
[**DeleteProjectUser**](ProjectApi.md#DeleteProjectUser) | **Delete** /api/v1/projects/{projectName}/users/{userName} | delete a user from a project
[**DetailConfig**](ProjectApi.md#DetailConfig) | **Get** /api/v1/projects/{projectName}/configs/{configName} | detail a config in a project
[**DetailProject**](ProjectApi.md#DetailProject) | **Get** /api/v1/projects/{projectName} | detail a project
[**GetConfigTemplate**](ProjectApi.md#GetConfigTemplate) | **Get** /api/v1/projects/{projectName}/config_templates/{templateName} | Detail a template
[**GetConfigTemplates**](ProjectApi.md#GetConfigTemplates) | **Get** /api/v1/projects/{projectName}/config_templates | get the templates which are in a project
[**GetConfigs**](ProjectApi.md#GetConfigs) | **Get** /api/v1/projects/{projectName}/configs | get configs which are in a project
[**GetProviders**](ProjectApi.md#GetProviders) | **Get** /api/v1/projects/{projectName}/providers | get providers which are in a project
[**ListDistributions**](ProjectApi.md#ListDistributions) | **Get** /api/v1/projects/{projectName}/distributions | list the distribution jobs of the config
[**ListProjectPermissions**](ProjectApi.md#ListProjectPermissions) | **Get** /api/v1/projects/{projectName}/permissions | list all project level perm policies
[**ListProjectRoles**](ProjectApi.md#ListProjectRoles) | **Get** /api/v1/projects/{projectName}/roles | list all project level roles
[**ListProjectTargets**](ProjectApi.md#ListProjectTargets) | **Get** /api/v1/projects/{projectName}/targets | get targets list belong to a project
[**ListProjectUser**](ProjectApi.md#ListProjectUser) | **Get** /api/v1/projects/{projectName}/users | list all users belong to a project
[**ListProjects**](ProjectApi.md#ListProjects) | **Get** /api/v1/projects | list all projects
[**UpdateConfig**](ProjectApi.md#UpdateConfig) | **Put** /api/v1/projects/{projectName}/configs/{configName} | update a config in a project
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Put** /api/v1/projects/{projectName} | update a project
[**UpdateProjectRole**](ProjectApi.md#UpdateProjectRole) | **Put** /api/v1/projects/{projectName}/roles/{roleName} | update project level role
[**UpdateProjectUser**](ProjectApi.md#UpdateProjectUser) | **Put** /api/v1/projects/{projectName}/users/{userName} | update a user from a project



## ApplyDistribution

> string ApplyDistribution(ctx, projectName, optional)

apply the distribution job of the config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
 **optional** | ***ApplyDistributionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplyDistributionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateConfigDistributionRequest** | [**optional.Interface of V1CreateConfigDistributionRequest**](V1CreateConfigDistributionRequest.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfig

> V1Config CreateConfig(ctx, projectName, optional)

create a config in a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
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


## CreateProject

> V1ProjectBase CreateProject(ctx, optional)

create a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateProjectRequest** | [**optional.Interface of V1CreateProjectRequest**](V1CreateProjectRequest.md)|  | 

### Return type

[**V1ProjectBase**](V1ProjectBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectPermission

> []V1PermissionBase CreateProjectPermission(ctx, projectName)

create a project level perm policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

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


## CreateProjectRole

> V1RoleBase CreateProjectRole(ctx, projectName, optional)

create project level role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
 **optional** | ***CreateProjectRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateProjectRoleOpts struct


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


## CreateProjectUser

> V1ProjectUserBase CreateProjectUser(ctx, projectName, optional)

add a user to a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
 **optional** | ***CreateProjectUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateProjectUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1AddProjectUserRequest** | [**optional.Interface of V1AddProjectUserRequest**](V1AddProjectUserRequest.md)|  | 

### Return type

[**V1ProjectUserBase**](V1ProjectUserBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfig

> string DeleteConfig(ctx, projectName, configName)

delete a config from a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
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


## DeleteDistribution

> string DeleteDistribution(ctx, projectName, distributionName)

delete a distribution job of the config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**distributionName** | **string**| identifier of the distribution | 

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


## DeleteProject

> string DeleteProject(ctx, projectName)

delete a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

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


## DeleteProjectPermission

> []V1PermissionBase DeleteProjectPermission(ctx, projectName, permissionName)

delete a project level perm policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**permissionName** | **string**| identifier of the permission | 

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


## DeleteProjectRole

> string DeleteProjectRole(ctx, projectName, roleName)

delete project level role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**roleName** | **string**| identifier of the project role | 

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


## DeleteProjectUser

> string DeleteProjectUser(ctx, projectName, userName, optional)

delete a user from a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**userName** | **string**| identifier of the project user | 
 **optional** | ***DeleteProjectUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteProjectUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateProjectUserRequest** | [**optional.Interface of V1UpdateProjectUserRequest**](V1UpdateProjectUserRequest.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetailConfig

> V1Config DetailConfig(ctx, projectName, configName, optional)

detail a config in a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**configName** | **string**| identifier of the config | 
 **optional** | ***DetailConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DetailConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateConfigRequest** | [**optional.Interface of V1UpdateConfigRequest**](V1UpdateConfigRequest.md)|  | 

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


## DetailProject

> V1ProjectBase DetailProject(ctx, projectName)

detail a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

### Return type

[**V1ProjectBase**](V1ProjectBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigTemplate

> V1ConfigTemplateDetail GetConfigTemplate(ctx, templateName, projectName, optional)

Detail a template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string**| identifier of the config template | 
**projectName** | **string**|  | 
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


## GetConfigTemplates

> V1ListConfigTemplateResponse GetConfigTemplates(ctx, projectName, namespace)

get the templates which are in a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**namespace** | **string**| the namespace of the template | 

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


## GetConfigs

> V1ListConfigResponse GetConfigs(ctx, projectName, optional)

get configs which are in a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
 **optional** | ***GetConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **template** | **optional.String**| the template name | 

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


## GetProviders

> V1ListTerraformProviderResponse GetProviders(ctx, projectName)

get providers which are in a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

### Return type

[**V1ListTerraformProviderResponse**](V1ListTerraformProviderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDistributions

> V1ListConfigDistributionResponse ListDistributions(ctx, projectName)

list the distribution jobs of the config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

### Return type

[**V1ListConfigDistributionResponse**](V1ListConfigDistributionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectPermissions

> []V1PermissionBase ListProjectPermissions(ctx, projectName)

list all project level perm policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

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


## ListProjectRoles

> V1ListRolesResponse ListProjectRoles(ctx, projectName)

list all project level roles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

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


## ListProjectTargets

> string ListProjectTargets(ctx, projectName)

get targets list belong to a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

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


## ListProjectUser

> V1ListProjectUsersResponse ListProjectUser(ctx, projectName)

list all users belong to a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 

### Return type

[**V1ListProjectUsersResponse**](V1ListProjectUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> V1ListProjectResponse ListProjects(ctx, )

list all projects

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1ListProjectResponse**](V1ListProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfig

> V1Config UpdateConfig(ctx, projectName, configName, optional)

update a config in a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**configName** | **string**| identifier of the config | 
 **optional** | ***UpdateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateConfigRequest** | [**optional.Interface of V1UpdateConfigRequest**](V1UpdateConfigRequest.md)|  | 

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


## UpdateProject

> V1ProjectBase UpdateProject(ctx, projectName, optional)

update a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
 **optional** | ***UpdateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1UpdateProjectRequest** | [**optional.Interface of V1UpdateProjectRequest**](V1UpdateProjectRequest.md)|  | 

### Return type

[**V1ProjectBase**](V1ProjectBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectRole

> V1RoleBase UpdateProjectRole(ctx, projectName, roleName, optional)

update project level role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**roleName** | **string**| identifier of the project role | 
 **optional** | ***UpdateProjectRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProjectRoleOpts struct


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


## UpdateProjectUser

> V1ProjectUserBase UpdateProjectUser(ctx, projectName, userName, optional)

update a user from a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| identifier of the project | 
**userName** | **string**| identifier of the project user | 
 **optional** | ***UpdateProjectUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProjectUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateProjectUserRequest** | [**optional.Interface of V1UpdateProjectUserRequest**](V1UpdateProjectUserRequest.md)|  | 

### Return type

[**V1ProjectUserBase**](V1ProjectUserBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

