# \ApplicationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplicationTrait**](ApplicationApi.md#AddApplicationTrait) | **Post** /api/v1/applications/{appName}/components/{compName}/traits | add trait for a component
[**ApplicationStatistics**](ApplicationApi.md#ApplicationStatistics) | **Get** /api/v1/applications/{appName}/statistics | detail one application 
[**CompareApp**](ApplicationApi.md#CompareApp) | **Post** /api/v1/applications/{appName}/compare | compare application
[**CreateApplication**](ApplicationApi.md#CreateApplication) | **Post** /api/v1/applications | create one application 
[**CreateApplicationEnv**](ApplicationApi.md#CreateApplicationEnv) | **Post** /api/v1/applications/{appName}/envs | creating an application environment 
[**CreateApplicationPolicy**](ApplicationApi.md#CreateApplicationPolicy) | **Post** /api/v1/applications/{appName}/policies | create policy for application
[**CreateApplicationTrigger**](ApplicationApi.md#CreateApplicationTrigger) | **Post** /api/v1/applications/{appName}/triggers | Create an application trigger
[**CreateComponent**](ApplicationApi.md#CreateComponent) | **Post** /api/v1/applications/{appName}/components | create component  for application 
[**CreateOrUpdateApplicationWorkflow**](ApplicationApi.md#CreateOrUpdateApplicationWorkflow) | **Post** /api/v1/applications/{appName}/workflows | create application workflow
[**DeleteApplication**](ApplicationApi.md#DeleteApplication) | **Delete** /api/v1/applications/{appName} | delete one application
[**DeleteApplicationEnv**](ApplicationApi.md#DeleteApplicationEnv) | **Delete** /api/v1/applications/{appName}/envs/{envName} | delete an application environment 
[**DeleteApplicationPolicy**](ApplicationApi.md#DeleteApplicationPolicy) | **Delete** /api/v1/applications/{appName}/policies/{policyName} | detail policy for application
[**DeleteApplicationTrait**](ApplicationApi.md#DeleteApplicationTrait) | **Delete** /api/v1/applications/{appName}/components/{compName}/traits/{traitType} | delete trait from a component
[**DeleteApplicationTrigger**](ApplicationApi.md#DeleteApplicationTrigger) | **Delete** /api/v1/applications/{appName}/triggers/{token} | Delete an application trigger
[**DeleteComponent**](ApplicationApi.md#DeleteComponent) | **Delete** /api/v1/applications/{appName}/components/{compName} | delete a component
[**DeleteWorkflow**](ApplicationApi.md#DeleteWorkflow) | **Delete** /api/v1/applications/{appName}/workflows/{workflowName} | deletet workflow
[**DeployApplication**](ApplicationApi.md#DeployApplication) | **Post** /api/v1/applications/{appName}/deploy | deploy or upgrade the application
[**DetailApplication**](ApplicationApi.md#DetailApplication) | **Get** /api/v1/applications/{appName} | detail one application 
[**DetailApplicationPolicy**](ApplicationApi.md#DetailApplicationPolicy) | **Get** /api/v1/applications/{appName}/policies/{policyName} | detail policy for application
[**DetailApplicationRevision**](ApplicationApi.md#DetailApplicationRevision) | **Get** /api/v1/applications/{appName}/revisions/{revision} | detail revision for application
[**DetailComponent**](ApplicationApi.md#DetailComponent) | **Get** /api/v1/applications/{appName}/components/{compName} | detail component for application 
[**DetailWorkflow**](ApplicationApi.md#DetailWorkflow) | **Get** /api/v1/applications/{appName}/workflows/{workflowName} | detail application workflow
[**DetailWorkflowRecord**](ApplicationApi.md#DetailWorkflowRecord) | **Get** /api/v1/applications/{appName}/workflows/{workflowName}/records/{record} | query application workflow execution record detail
[**DryRunAppOrRevision**](ApplicationApi.md#DryRunAppOrRevision) | **Post** /api/v1/applications/{appName}/dry-run | dry-run application to latest revision
[**GetApplicationStatus**](ApplicationApi.md#GetApplicationStatus) | **Get** /api/v1/applications/{appName}/envs/{envName}/status | get application status
[**GetWorkflowRecordInputs**](ApplicationApi.md#GetWorkflowRecordInputs) | **Get** /api/v1/applications/{appName}/workflows/{workflowName}/records/{record}/inputs | get the workflow step inputs
[**GetWorkflowRecordLogs**](ApplicationApi.md#GetWorkflowRecordLogs) | **Get** /api/v1/applications/{appName}/workflows/{workflowName}/records/{record}/logs | get the workflow step logs
[**GetWorkflowRecordOutputs**](ApplicationApi.md#GetWorkflowRecordOutputs) | **Get** /api/v1/applications/{appName}/workflows/{workflowName}/records/{record}/outputs | get the workflow step inputs
[**ListApplicationComponents**](ApplicationApi.md#ListApplicationComponents) | **Get** /api/v1/applications/{appName}/components | gets the list of application components
[**ListApplicationEnvs**](ApplicationApi.md#ListApplicationEnvs) | **Get** /api/v1/applications/{appName}/envs | list policy for application
[**ListApplicationPolicies**](ApplicationApi.md#ListApplicationPolicies) | **Get** /api/v1/applications/{appName}/policies | list policy for application
[**ListApplicationRecords**](ApplicationApi.md#ListApplicationRecords) | **Get** /api/v1/applications/{appName}/records | list application records
[**ListApplicationRevisions**](ApplicationApi.md#ListApplicationRevisions) | **Get** /api/v1/applications/{appName}/revisions | list revisions for application
[**ListApplicationTriggers**](ApplicationApi.md#ListApplicationTriggers) | **Get** /api/v1/applications/{appName}/triggers | List the application triggers
[**ListApplicationWorkflows**](ApplicationApi.md#ListApplicationWorkflows) | **Get** /api/v1/applications/{appName}/workflows | list application workflow
[**ListApplications**](ApplicationApi.md#ListApplications) | **Get** /api/v1/applications | list all applications
[**ListWorkflowRecords**](ApplicationApi.md#ListWorkflowRecords) | **Get** /api/v1/applications/{appName}/workflows/{workflowName}/records | query application workflow execution record
[**PublishApplicationTemplate**](ApplicationApi.md#PublishApplicationTemplate) | **Post** /api/v1/applications/{appName}/template | create one application template
[**RecycleApplicationEnv**](ApplicationApi.md#RecycleApplicationEnv) | **Post** /api/v1/applications/{appName}/envs/{envName}/recycle | get application status
[**ResetAppToLatestRevision**](ApplicationApi.md#ResetAppToLatestRevision) | **Post** /api/v1/applications/{appName}/reset | reset application to latest revision
[**ResumeWorkflowRecord**](ApplicationApi.md#ResumeWorkflowRecord) | **Get** /api/v1/applications/{appName}/workflows/{workflowName}/records/{record}/resume | resume suspend workflow record
[**RollbackApplicationWithRevision**](ApplicationApi.md#RollbackApplicationWithRevision) | **Post** /api/v1/applications/{appName}/revisions/{revision}/rollback | detail revision for application
[**RollbackWorkflowRecord**](ApplicationApi.md#RollbackWorkflowRecord) | **Get** /api/v1/applications/{appName}/workflows/{workflowName}/records/{record}/rollback | rollback suspend application record
[**TerminateWorkflowRecord**](ApplicationApi.md#TerminateWorkflowRecord) | **Get** /api/v1/applications/{appName}/workflows/{workflowName}/records/{record}/terminate | terminate suspend workflow record
[**UpdateApplication**](ApplicationApi.md#UpdateApplication) | **Put** /api/v1/applications/{appName} | update one application 
[**UpdateApplicationEnv**](ApplicationApi.md#UpdateApplicationEnv) | **Put** /api/v1/applications/{appName}/envs/{envName} | set application  differences in the specified environment
[**UpdateApplicationPolicy**](ApplicationApi.md#UpdateApplicationPolicy) | **Put** /api/v1/applications/{appName}/policies/{policyName} | update policy for application
[**UpdateApplicationTrait**](ApplicationApi.md#UpdateApplicationTrait) | **Put** /api/v1/applications/{appName}/components/{compName}/traits/{traitType} | update trait from a component
[**UpdateApplicationTrigger**](ApplicationApi.md#UpdateApplicationTrigger) | **Put** /api/v1/applications/{appName}/triggers/{token} | Update an application trigger
[**UpdateComponent**](ApplicationApi.md#UpdateComponent) | **Put** /api/v1/applications/{appName}/components/{compName} | update component config
[**UpdateWorkflow**](ApplicationApi.md#UpdateWorkflow) | **Put** /api/v1/applications/{appName}/workflows/{workflowName} | update application workflow config



## AddApplicationTrait

> string AddApplicationTrait(ctx, appName, compName, optional)

add trait for a component

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**compName** | **string**| identifier of the component | 
 **optional** | ***AddApplicationTraitOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddApplicationTraitOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1CreateApplicationTraitRequest** | [**optional.Interface of V1CreateApplicationTraitRequest**](V1CreateApplicationTraitRequest.md)|  | 

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


## ApplicationStatistics

> V1ApplicationStatisticsResponse ApplicationStatistics(ctx, appName)

detail one application 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 

### Return type

[**V1ApplicationStatisticsResponse**](V1ApplicationStatisticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompareApp

> V1AppCompareResponse CompareApp(ctx, appName, optional)

compare application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***CompareAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CompareAppOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1AppCompareReq** | [**optional.Interface of V1AppCompareReq**](V1AppCompareReq.md)|  | 

### Return type

[**V1AppCompareResponse**](V1AppCompareResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplication

> V1ApplicationBase CreateApplication(ctx, optional)

create one application 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateApplicationRequest** | [**optional.Interface of V1CreateApplicationRequest**](V1CreateApplicationRequest.md)|  | 

### Return type

[**V1ApplicationBase**](V1ApplicationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationEnv

> V1EnvBinding CreateApplicationEnv(ctx, appName, optional)

creating an application environment 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***CreateApplicationEnvOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApplicationEnvOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateApplicationEnvbindingRequest** | [**optional.Interface of V1CreateApplicationEnvbindingRequest**](V1CreateApplicationEnvbindingRequest.md)|  | 

### Return type

[**V1EnvBinding**](V1EnvBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationPolicy

> V1PolicyBase CreateApplicationPolicy(ctx, appName, optional)

create policy for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
 **optional** | ***CreateApplicationPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApplicationPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreatePolicyRequest** | [**optional.Interface of V1CreatePolicyRequest**](V1CreatePolicyRequest.md)|  | 

### Return type

[**V1PolicyBase**](V1PolicyBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationTrigger

> V1ApplicationTriggerBase CreateApplicationTrigger(ctx, appName, optional)

Create an application trigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***CreateApplicationTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApplicationTriggerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateApplicationTriggerRequest** | [**optional.Interface of V1CreateApplicationTriggerRequest**](V1CreateApplicationTriggerRequest.md)|  | 

### Return type

[**V1ApplicationTriggerBase**](V1ApplicationTriggerBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComponent

> V1ComponentBase CreateComponent(ctx, appName, optional)

create component  for application 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***CreateComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateComponentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateComponentRequest** | [**optional.Interface of V1CreateComponentRequest**](V1CreateComponentRequest.md)|  | 

### Return type

[**V1ComponentBase**](V1ComponentBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateApplicationWorkflow

> V1SimpleResponse CreateOrUpdateApplicationWorkflow(ctx, appName, optional)

create application workflow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
 **optional** | ***CreateOrUpdateApplicationWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOrUpdateApplicationWorkflowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateWorkflowRequest** | [**optional.Interface of V1CreateWorkflowRequest**](V1CreateWorkflowRequest.md)|  | 

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


## DeleteApplication

> string DeleteApplication(ctx, appName)

delete one application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 

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


## DeleteApplicationEnv

> string DeleteApplicationEnv(ctx, appName, envName)

delete an application environment 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
**envName** | **string**| identifier of the envBinding  | 

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


## DeleteApplicationPolicy

> string DeleteApplicationPolicy(ctx, appName, policyName, optional)

detail policy for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**policyName** | **string**| identifier of the application policy | 
 **optional** | ***DeleteApplicationPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteApplicationPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.String**| Force delete the policy and all references | 

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


## DeleteApplicationTrait

> V1ApplicationTrait DeleteApplicationTrait(ctx, appName, compName, traitType)

delete trait from a component

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**compName** | **string**| identifier of the component | 
**traitType** | **string**| identifier of the type of trait | 

### Return type

[**V1ApplicationTrait**](V1ApplicationTrait.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationTrigger

> string DeleteApplicationTrigger(ctx, appName, token)

Delete an application trigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
**token** | **string**| identifier of the trigger | 

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


## DeleteComponent

> string DeleteComponent(ctx, appName, compName)

delete a component

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**compName** | **string**| identifier of the component | 

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


## DeleteWorkflow

> V1SimpleResponse DeleteWorkflow(ctx, appName, workflowName)

deletet workflow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 

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


## DeployApplication

> V1ApplicationDeployResponse DeployApplication(ctx, appName, optional)

deploy or upgrade the application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***DeployApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeployApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1ApplicationDeployRequest** | [**optional.Interface of V1ApplicationDeployRequest**](V1ApplicationDeployRequest.md)|  | 

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


## DetailApplication

> V1DetailApplicationResponse DetailApplication(ctx, appName)

detail one application 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 

### Return type

[**V1DetailApplicationResponse**](V1DetailApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetailApplicationPolicy

> V1DetailPolicyResponse DetailApplicationPolicy(ctx, appName, policyName)

detail policy for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**policyName** | **string**| identifier of the application policy | 

### Return type

[**V1DetailPolicyResponse**](V1DetailPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetailApplicationRevision

> V1DetailRevisionResponse DetailApplicationRevision(ctx, appName, revision)

detail revision for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**revision** | **string**| identifier of the application revision | 

### Return type

[**V1DetailRevisionResponse**](V1DetailRevisionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetailComponent

> V1DetailComponentResponse DetailComponent(ctx, appName, compName)

detail component for application 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
**compName** | **string**| identifier of the component | 

### Return type

[**V1DetailComponentResponse**](V1DetailComponentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetailWorkflow

> V1SimpleResponse DetailWorkflow(ctx, appName, workflowName)

detail application workflow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workfloc. | 

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


## DetailWorkflowRecord

> V1SimpleResponse DetailWorkflowRecord(ctx, appName, workflowName, record)

query application workflow execution record detail

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
**record** | **string**| identifier of the workflow record | 

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


## DryRunAppOrRevision

> V1AppDryRunResponse DryRunAppOrRevision(ctx, appName, optional)

dry-run application to latest revision

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***DryRunAppOrRevisionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DryRunAppOrRevisionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1AppDryRunReq** | [**optional.Interface of V1AppDryRunReq**](V1AppDryRunReq.md)|  | 

### Return type

[**V1AppDryRunResponse**](V1AppDryRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationStatus

> V1ApplicationStatusResponse GetApplicationStatus(ctx, appName, envName)

get application status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
**envName** | **string**| identifier of the application envbinding | 

### Return type

[**V1ApplicationStatusResponse**](V1ApplicationStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowRecordInputs

> V1GetPipelineRunInputResponse GetWorkflowRecordInputs(ctx, appName, workflowName, record, step)

get the workflow step inputs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
**record** | **string**| identifier of the workflow record | 
**step** | **string**| Specified the step filter | 

### Return type

[**V1GetPipelineRunInputResponse**](V1GetPipelineRunInputResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowRecordLogs

> map[string]interface{} GetWorkflowRecordLogs(ctx, appName, workflowName, record, step)

get the workflow step logs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
**record** | **string**| identifier of the workflow record | 
**step** | **string**| Specified the step filter | 

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


## GetWorkflowRecordOutputs

> V1GetPipelineRunOutputResponse GetWorkflowRecordOutputs(ctx, appName, workflowName, record, step)

get the workflow step inputs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
**record** | **string**| identifier of the workflow record | 
**step** | **string**| Specified the step filter | 

### Return type

[**V1GetPipelineRunOutputResponse**](V1GetPipelineRunOutputResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationComponents

> V1ComponentListResponse ListApplicationComponents(ctx, appName, optional)

gets the list of application components

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***ListApplicationComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApplicationComponentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envName** | **optional.String**| list components that deployed in define env | 

### Return type

[**V1ComponentListResponse**](V1ComponentListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationEnvs

> V1ListApplicationEnvBinding ListApplicationEnvs(ctx, appName)

list policy for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 

### Return type

[**V1ListApplicationEnvBinding**](V1ListApplicationEnvBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationPolicies

> V1ListApplicationPolicy ListApplicationPolicies(ctx, appName)

list policy for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 

### Return type

[**V1ListApplicationPolicy**](V1ListApplicationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationRecords

> map[string]interface{} ListApplicationRecords(ctx, appName)

list application records

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 

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


## ListApplicationRevisions

> V1ListRevisionsResponse ListApplicationRevisions(ctx, appName, optional)

list revisions for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***ListApplicationRevisionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApplicationRevisionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envName** | **optional.String**| query identifier of the env | 
 **status** | **optional.String**| query identifier of the status | 
 **page** | **optional.Int32**| query the page number | 
 **pageSize** | **optional.Int32**| query the page size number | 

### Return type

[**V1ListRevisionsResponse**](V1ListRevisionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationTriggers

> V1ListApplicationTriggerResponse ListApplicationTriggers(ctx, appName)

List the application triggers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 

### Return type

[**V1ListApplicationTriggerResponse**](V1ListApplicationTriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationWorkflows

> V1SimpleResponse ListApplicationWorkflows(ctx, appName)

list application workflow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 

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


## ListApplications

> V1ListApplicationResponse ListApplications(ctx, optional)

list all applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Fuzzy search based on name or description | 
 **project** | **optional.String**| search base on project name | 
 **env** | **optional.String**| search base on env name | 
 **targetName** | **optional.String**| Name of the application delivery target | 

### Return type

[**V1ListApplicationResponse**](V1ListApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowRecords

> V1SimpleResponse ListWorkflowRecords(ctx, appName, workflowName, optional)

query application workflow execution record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
 **optional** | ***ListWorkflowRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkflowRecordsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| query the page number | 
 **pageSize** | **optional.Int32**| query the page size number | 

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


## PublishApplicationTemplate

> V1ApplicationTemplateBase PublishApplicationTemplate(ctx, appName, optional)

create one application template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***PublishApplicationTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PublishApplicationTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateApplicationTemplateRequest** | [**optional.Interface of V1CreateApplicationTemplateRequest**](V1CreateApplicationTemplateRequest.md)|  | 

### Return type

[**V1ApplicationTemplateBase**](V1ApplicationTemplateBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecycleApplicationEnv

> string RecycleApplicationEnv(ctx, appName, envName)

get application status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
**envName** | **string**| identifier of the application envbinding | 

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


## ResetAppToLatestRevision

> V1AppResetResponse ResetAppToLatestRevision(ctx, appName)

reset application to latest revision

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 

### Return type

[**V1AppResetResponse**](V1AppResetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeWorkflowRecord

> string ResumeWorkflowRecord(ctx, appName, workflowName, record)

resume suspend workflow record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
**record** | **string**| identifier of the  workflow record | 

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


## RollbackApplicationWithRevision

> V1ApplicationRollbackResponse RollbackApplicationWithRevision(ctx, appName, revision)

detail revision for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**revision** | **string**| identifier of the application revision | 

### Return type

[**V1ApplicationRollbackResponse**](V1ApplicationRollbackResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackWorkflowRecord

> V1WorkflowRecordBase RollbackWorkflowRecord(ctx, appName, workflowName, record, optional)

rollback suspend application record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
**record** | **string**| identifier of the workflow record | 
 **optional** | ***RollbackWorkflowRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RollbackWorkflowRecordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rollbackVersion** | **optional.String**| identifier of the rollback revision | 

### Return type

[**V1WorkflowRecordBase**](V1WorkflowRecordBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateWorkflowRecord

> string TerminateWorkflowRecord(ctx, appName, workflowName, record)

terminate suspend workflow record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
**record** | **string**| identifier of the workflow record | 

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


## UpdateApplication

> V1ApplicationBase UpdateApplication(ctx, appName, optional)

update one application 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
 **optional** | ***UpdateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1UpdateApplicationRequest** | [**optional.Interface of V1UpdateApplicationRequest**](V1UpdateApplicationRequest.md)|  | 

### Return type

[**V1ApplicationBase**](V1ApplicationBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationEnv

> V1EnvBinding UpdateApplicationEnv(ctx, appName, envName, optional)

set application  differences in the specified environment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
**envName** | **string**| identifier of the envBinding  | 
 **optional** | ***UpdateApplicationEnvOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApplicationEnvOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **optional.String**|  | 

### Return type

[**V1EnvBinding**](V1EnvBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationPolicy

> V1DetailPolicyResponse UpdateApplicationPolicy(ctx, appName, policyName, optional)

update policy for application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**policyName** | **string**| identifier of the application policy | 
 **optional** | ***UpdateApplicationPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApplicationPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdatePolicyRequest** | [**optional.Interface of V1UpdatePolicyRequest**](V1UpdatePolicyRequest.md)|  | 

### Return type

[**V1DetailPolicyResponse**](V1DetailPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationTrait

> V1ApplicationTrait UpdateApplicationTrait(ctx, appName, compName, traitType, optional)

update trait from a component

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**compName** | **string**| identifier of the component | 
**traitType** | **string**| identifier of the type of trait | 
 **optional** | ***UpdateApplicationTraitOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApplicationTraitOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v1UpdateApplicationTraitRequest** | [**optional.Interface of V1UpdateApplicationTraitRequest**](V1UpdateApplicationTraitRequest.md)|  | 

### Return type

[**V1ApplicationTrait**](V1ApplicationTrait.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationTrigger

> V1ApplicationTriggerBase UpdateApplicationTrigger(ctx, appName, token)

Update an application trigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application  | 
**token** | **string**| identifier of the trigger | 

### Return type

[**V1ApplicationTriggerBase**](V1ApplicationTriggerBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComponent

> V1ComponentBase UpdateComponent(ctx, appName, compName, optional)

update component config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application | 
**compName** | **string**| identifier of the component | 
 **optional** | ***UpdateComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateComponentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateApplicationComponentRequest** | [**optional.Interface of V1UpdateApplicationComponentRequest**](V1UpdateApplicationComponentRequest.md)|  | 

### Return type

[**V1ComponentBase**](V1ComponentBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflow

> V1SimpleResponse UpdateWorkflow(ctx, appName, workflowName, optional)

update application workflow config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string**| identifier of the application. | 
**workflowName** | **string**| identifier of the workflow | 
 **optional** | ***UpdateWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWorkflowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateWorkflowRequest** | [**optional.Interface of V1UpdateWorkflowRequest**](V1UpdateWorkflowRequest.md)|  | 

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

