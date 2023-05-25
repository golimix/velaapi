# \PipelineApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContextValue**](PipelineApi.md#CreateContextValue) | **Post** /api/v1/projects/{projectName}/pipelines/{pipelineName}/contexts | create pipeline context values
[**CreatePipeline**](PipelineApi.md#CreatePipeline) | **Post** /api/v1/projects/{projectName}/pipelines | create pipeline
[**DeleteContextValue**](PipelineApi.md#DeleteContextValue) | **Delete** /api/v1/projects/{projectName}/pipelines/{pipelineName}/contexts/{contextName} | delete pipeline context value
[**DeletePipeline**](PipelineApi.md#DeletePipeline) | **Delete** /api/v1/projects/{projectName}/pipelines/{pipelineName} | delete pipeline
[**DeletePipelineRun**](PipelineApi.md#DeletePipelineRun) | **Delete** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName} | delete pipeline run
[**GetPipeline**](PipelineApi.md#GetPipeline) | **Get** /api/v1/projects/{projectName}/pipelines/{pipelineName} | get pipeline
[**GetPipelineRun**](PipelineApi.md#GetPipelineRun) | **Get** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName} | get pipeline run
[**GetPipelineRunInput**](PipelineApi.md#GetPipelineRunInput) | **Get** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName}/input | get pipeline run input
[**GetPipelineRunLog**](PipelineApi.md#GetPipelineRunLog) | **Get** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName}/log | get pipeline run log
[**GetPipelineRunOutput**](PipelineApi.md#GetPipelineRunOutput) | **Get** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName}/output | get pipeline run output
[**GetPipelineRunStatus**](PipelineApi.md#GetPipelineRunStatus) | **Get** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName}/status | get pipeline run status
[**ListContextValues**](PipelineApi.md#ListContextValues) | **Get** /api/v1/projects/{projectName}/pipelines/{pipelineName}/contexts | list pipeline context values
[**ListPipelineRuns**](PipelineApi.md#ListPipelineRuns) | **Get** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs | list pipeline runs
[**ListPipelines**](PipelineApi.md#ListPipelines) | **Get** /api/v1/pipelines | list pipelines
[**ResumePipelineRun**](PipelineApi.md#ResumePipelineRun) | **Post** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName}/resume | resume suspend pipeline run
[**RunPipeline**](PipelineApi.md#RunPipeline) | **Post** /api/v1/projects/{projectName}/pipelines/{pipelineName}/run | run pipeline
[**StopPipeline**](PipelineApi.md#StopPipeline) | **Post** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName}/stop | stop pipeline run
[**TerminatePipelineRun**](PipelineApi.md#TerminatePipelineRun) | **Post** /api/v1/projects/{projectName}/pipelines/{pipelineName}/runs/{runName}/terminate | resume suspend pipeline run
[**UpdateContextValue**](PipelineApi.md#UpdateContextValue) | **Put** /api/v1/projects/{projectName}/pipelines/{pipelineName}/contexts/{contextName} | update pipeline context value
[**UpdatePipeline**](PipelineApi.md#UpdatePipeline) | **Put** /api/v1/projects/{projectName}/pipelines/{pipelineName} | update pipeline



## CreateContextValue

> V1Context CreateContextValue(ctx, projectName, pipelineName, optional)

create pipeline context values

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
 **optional** | ***CreateContextValueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateContextValueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1CreateContextValuesRequest** | [**optional.Interface of V1CreateContextValuesRequest**](V1CreateContextValuesRequest.md)|  | 

### Return type

[**V1Context**](V1Context.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipeline

> V1PipelineBase CreatePipeline(ctx, projectName, optional)

create pipeline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
 **optional** | ***CreatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePipelineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreatePipelineRequest** | [**optional.Interface of V1CreatePipelineRequest**](V1CreatePipelineRequest.md)|  | 

### Return type

[**V1PipelineBase**](V1PipelineBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContextValue

> V1ContextNameResponse DeleteContextValue(ctx, projectName, pipelineName, contextName)

delete pipeline context value

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**contextName** | **string**| pipeline context name | 

### Return type

[**V1ContextNameResponse**](V1ContextNameResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipeline

> V1PipelineMetaResponse DeletePipeline(ctx, projectName, pipelineName)

delete pipeline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 

### Return type

[**V1PipelineMetaResponse**](V1PipelineMetaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipelineRun

> V1PipelineRunMeta DeletePipelineRun(ctx, projectName, pipelineName, runName)

delete pipeline run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 

### Return type

[**V1PipelineRunMeta**](V1PipelineRunMeta.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipeline

> V1GetPipelineResponse GetPipeline(ctx, pipelineName, projectName)

get pipeline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineName** | **string**| pipeline name | 
**projectName** | **string**| project name | 

### Return type

[**V1GetPipelineResponse**](V1GetPipelineResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineRun

> V1PipelineRunBase GetPipelineRun(ctx, projectName, pipelineName, runName)

get pipeline run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 

### Return type

[**V1PipelineRunBase**](V1PipelineRunBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineRunInput

> V1GetPipelineRunInputResponse GetPipelineRunInput(ctx, projectName, pipelineName, runName, step)

get pipeline run input

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 
**step** | **string**| query by specific step name | 

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


## GetPipelineRunLog

> V1GetPipelineRunLogResponse GetPipelineRunLog(ctx, projectName, pipelineName, runName, optional)

get pipeline run log

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 
 **optional** | ***GetPipelineRunLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPipelineRunLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **step** | **optional.String**| query by specific step name | 

### Return type

[**V1GetPipelineRunLogResponse**](V1GetPipelineRunLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineRunOutput

> V1GetPipelineRunOutputResponse GetPipelineRunOutput(ctx, projectName, pipelineName, runName, step)

get pipeline run output

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 
**step** | **string**| query by specific step name | 

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


## GetPipelineRunStatus

> V1alpha1WorkflowRunStatus GetPipelineRunStatus(ctx, projectName, pipelineName, runName)

get pipeline run status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 

### Return type

[**V1alpha1WorkflowRunStatus**](V1alpha1WorkflowRunStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContextValues

> V1ListContextValueResponse ListContextValues(ctx, projectName, pipelineName)

list pipeline context values

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 

### Return type

[**V1ListContextValueResponse**](V1ListContextValueResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelineRuns

> V1ListPipelineRunResponse ListPipelineRuns(ctx, projectName, pipelineName, optional)

list pipeline runs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
 **optional** | ***ListPipelineRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPipelineRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **optional.String**| query identifier of the status | 

### Return type

[**V1ListPipelineRunResponse**](V1ListPipelineRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelines

> V1ListPipelineResponse ListPipelines(ctx, optional)

list pipelines

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPipelinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Fuzzy search based on name or description | 
 **projectName** | **optional.String**| query pipelines within a project | 
 **detailed** | **optional.String**| query pipelines with detail | 

### Return type

[**V1ListPipelineResponse**](V1ListPipelineResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumePipelineRun

> string ResumePipelineRun(ctx, projectName, pipelineName, runName)

resume suspend pipeline run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 

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


## RunPipeline

> V1PipelineRun RunPipeline(ctx, projectName, pipelineName, optional)

run pipeline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
 **optional** | ***RunPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunPipelineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1RunPipelineRequest** | [**optional.Interface of V1RunPipelineRequest**](V1RunPipelineRequest.md)|  | 

### Return type

[**V1PipelineRun**](V1PipelineRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopPipeline

> V1PipelineRunMeta StopPipeline(ctx, projectName, pipelineName, runName)

stop pipeline run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 

### Return type

[**V1PipelineRunMeta**](V1PipelineRunMeta.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminatePipelineRun

> string TerminatePipelineRun(ctx, projectName, pipelineName, runName)

resume suspend pipeline run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**runName** | **string**| pipeline run name | 

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


## UpdateContextValue

> V1Context UpdateContextValue(ctx, projectName, pipelineName, contextName, optional)

update pipeline context value

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
**contextName** | **string**| pipeline context name | 
 **optional** | ***UpdateContextValueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateContextValueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v1UpdateContextValuesRequest** | [**optional.Interface of V1UpdateContextValuesRequest**](V1UpdateContextValuesRequest.md)|  | 

### Return type

[**V1Context**](V1Context.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipeline

> V1PipelineBase UpdatePipeline(ctx, projectName, pipelineName, optional)

update pipeline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string**| project name | 
**pipelineName** | **string**| pipeline name | 
 **optional** | ***UpdatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePipelineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdatePipelineRequest** | [**optional.Interface of V1UpdatePipelineRequest**](V1UpdatePipelineRequest.md)|  | 

### Return type

[**V1PipelineBase**](V1PipelineBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

