# \UnifiedapioasAPI

All URIs are relative to *http://localhost:5001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIGroup**](UnifiedapioasAPI.md#CreateAPIGroup) | **Post** /unified-routing/v1alpha1/api-groups | Create API group.
[**CreateInstance**](UnifiedapioasAPI.md#CreateInstance) | **Post** /unified-routing/v1alpha1/application-instances | Create application instance.
[**CreateProvision**](UnifiedapioasAPI.md#CreateProvision) | **Post** /unified-routing/v1alpha1/customer-provisions | Create customer provision.
[**DeleteAPIGroup**](UnifiedapioasAPI.md#DeleteAPIGroup) | **Delete** /unified-routing/v1alpha1/api-groups | Remove API group.
[**DeleteAPIGroupOnID**](UnifiedapioasAPI.md#DeleteAPIGroupOnID) | **Delete** /unified-routing/v1alpha1/api-groups/{id} | Remove API Group by ID.
[**GetAPIGroup**](UnifiedapioasAPI.md#GetAPIGroup) | **Get** /unified-routing/v1alpha1/api-groups | List API groups.
[**GetAPIGroupOnID**](UnifiedapioasAPI.md#GetAPIGroupOnID) | **Get** /unified-routing/v1alpha1/api-groups/{id} | Get API group by ID.
[**GetProvisionOnID**](UnifiedapioasAPI.md#GetProvisionOnID) | **Get** /unified-routing/v1alpha1/customer-provisions/{id} | Get customer provision by ID.
[**ListInstances**](UnifiedapioasAPI.md#ListInstances) | **Get** /unified-routing/v1alpha1/application-instances | List application instances.
[**ListProvisions**](UnifiedapioasAPI.md#ListProvisions) | **Get** /unified-routing/v1alpha1/customer-provisions | List customer provisions.
[**PatchInstance**](UnifiedapioasAPI.md#PatchInstance) | **Patch** /unified-routing/v1alpha1/application-instances | Update fields of an application instance.
[**RemoveInstance**](UnifiedapioasAPI.md#RemoveInstance) | **Delete** /unified-routing/v1alpha1/application-instances | Remove application instances.
[**RemoveProvision**](UnifiedapioasAPI.md#RemoveProvision) | **Delete** /unified-routing/v1alpha1/customer-provisions | Remove customer provisions.
[**RemoveProvisionOnID**](UnifiedapioasAPI.md#RemoveProvisionOnID) | **Delete** /unified-routing/v1alpha1/customer-provisions/{id} | Remove customer provision by ID.



## CreateAPIGroup

> ApiGroupExtended CreateAPIGroup(ctx).ApiGroupBasic(apiGroupBasic).Execute()

Create API group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	apiGroupBasic := *openapiclient.NewApiGroupBasic("ApplicationId_example", "ApiGroup_example") // ApiGroupBasic | Information about the API group record to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.CreateAPIGroup(context.Background()).ApiGroupBasic(apiGroupBasic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.CreateAPIGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAPIGroup`: ApiGroupExtended
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.CreateAPIGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiGroupBasic** | [**ApiGroupBasic**](ApiGroupBasic.md) | Information about the API group record to be created | 

### Return type

[**ApiGroupExtended**](ApiGroupExtended.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstance

> AppInstanceExtended CreateInstance(ctx).AppInstanceBasic(appInstanceBasic).Execute()

Create application instance.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	appInstanceBasic := *openapiclient.NewAppInstanceBasic("ApplicationInstanceId_example", "Region_example", "ApplicationId_example") // AppInstanceBasic | Information about the application instance record to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.CreateInstance(context.Background()).AppInstanceBasic(appInstanceBasic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstance`: AppInstanceExtended
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstanceBasic** | [**AppInstanceBasic**](AppInstanceBasic.md) | Information about the application instance record to be created | 

### Return type

[**AppInstanceExtended**](AppInstanceExtended.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProvision

> CustomerProvisionExtended CreateProvision(ctx).CustomerProvisionBasic(customerProvisionBasic).Execute()

Create customer provision.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	customerProvisionBasic := *openapiclient.NewCustomerProvisionBasic("ApplicationInstanceId_example", "Region_example", "ApplicationId_example", "WorkspaceId_example") // CustomerProvisionBasic | Information about the customer provision record to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.CreateProvision(context.Background()).CustomerProvisionBasic(customerProvisionBasic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.CreateProvision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProvision`: CustomerProvisionExtended
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.CreateProvision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProvisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerProvisionBasic** | [**CustomerProvisionBasic**](CustomerProvisionBasic.md) | Information about the customer provision record to be created | 

### Return type

[**CustomerProvisionExtended**](CustomerProvisionExtended.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIGroup

> DeleteAPIGroup(ctx).Filter(filter).Execute()

Remove API group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	filter := "filter_example" // string | The filter query parameter is used to limit the set of resources to be deleted in a collection-level DELETE.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | applicationId, apiGroup                    | | Logic      | and                                        |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UnifiedapioasAPI.DeleteAPIGroup(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.DeleteAPIGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter query parameter is used to limit the set of resources to be deleted in a collection-level DELETE.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | applicationId, apiGroup                    | | Logic      | and                                        |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIGroupOnID

> DeleteAPIGroupOnID(ctx, id).Execute()

Remove API Group by ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | API Group Record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UnifiedapioasAPI.DeleteAPIGroupOnID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.DeleteAPIGroupOnID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | API Group Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIGroupOnIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIGroup

> ApiGroupList GetAPIGroup(ctx).Filter(filter).Limit(limit).Next(next).Execute()

List API groups.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	filter := "filter_example" // string | The filter query parameter is used to limit the set of resources returned in a collection-level GET. The returned set of resources must match the criteria in the filter.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | applicationId, apiGroup                    | | Logic      | and                                        |  (optional)
	limit := int32(56) // int32 | Limits the number of results returned. (optional) (default to 50)
	next := "next_example" // string | The pagination cursor for the next page of resources. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.GetAPIGroup(context.Background()).Filter(filter).Limit(limit).Next(next).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.GetAPIGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIGroup`: ApiGroupList
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.GetAPIGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter query parameter is used to limit the set of resources returned in a collection-level GET. The returned set of resources must match the criteria in the filter.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | applicationId, apiGroup                    | | Logic      | and                                        |  | 
 **limit** | **int32** | Limits the number of results returned. | [default to 50]
 **next** | **string** | The pagination cursor for the next page of resources. | 

### Return type

[**ApiGroupList**](ApiGroupList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIGroupOnID

> ApiGroupExtended GetAPIGroupOnID(ctx, id).Execute()

Get API group by ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | API Group Record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.GetAPIGroupOnID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.GetAPIGroupOnID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIGroupOnID`: ApiGroupExtended
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.GetAPIGroupOnID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | API Group Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIGroupOnIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiGroupExtended**](ApiGroupExtended.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisionOnID

> CustomerProvisionExtended GetProvisionOnID(ctx, id).Execute()

Get customer provision by ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer Provision Record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.GetProvisionOnID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.GetProvisionOnID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisionOnID`: CustomerProvisionExtended
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.GetProvisionOnID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer Provision Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisionOnIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerProvisionExtended**](CustomerProvisionExtended.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> AppInstanceList ListInstances(ctx).Filter(filter).Limit(limit).Next(next).Execute()

List application instances.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	filter := "filter_example" // string | The filter query parameter is used to limit the set of resources returned in a collection-level GET. The returned set of resources must match the criteria in the filter.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | externalApiEndpoint, applicationInstanceId, applicationId, region | | Logic      | and                                        |  (optional)
	limit := int32(56) // int32 | Limits the number of results returned. Default is 50, max is 200. (optional) (default to 50)
	next := "next_example" // string | The pagination cursor for the next page of resources. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.ListInstances(context.Background()).Filter(filter).Limit(limit).Next(next).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.ListInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstances`: AppInstanceList
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.ListInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter query parameter is used to limit the set of resources returned in a collection-level GET. The returned set of resources must match the criteria in the filter.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | externalApiEndpoint, applicationInstanceId, applicationId, region | | Logic      | and                                        |  | 
 **limit** | **int32** | Limits the number of results returned. Default is 50, max is 200. | [default to 50]
 **next** | **string** | The pagination cursor for the next page of resources. | 

### Return type

[**AppInstanceList**](AppInstanceList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProvisions

> CustomerProvisionList ListProvisions(ctx).Filter(filter).Limit(limit).Next(next).Execute()

List customer provisions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	filter := "filter_example" // string | The filter query parameter is used to limit the set of resources returned in a collection-level GET. The returned set of resources must match the criteria in the filter.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | workspaceId, applicationInstanceId, applicationId, region | | Logic      | and                                        |  (optional)
	limit := int32(56) // int32 | Limits the number of results returned. Default is 50, max is 200. (optional) (default to 50)
	next := "next_example" // string | The pagination cursor for the next page of resources. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.ListProvisions(context.Background()).Filter(filter).Limit(limit).Next(next).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.ListProvisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProvisions`: CustomerProvisionList
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.ListProvisions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter query parameter is used to limit the set of resources returned in a collection-level GET. The returned set of resources must match the criteria in the filter.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | workspaceId, applicationInstanceId, applicationId, region | | Logic      | and                                        |  | 
 **limit** | **int32** | Limits the number of results returned. Default is 50, max is 200. | [default to 50]
 **next** | **string** | The pagination cursor for the next page of resources. | 

### Return type

[**CustomerProvisionList**](CustomerProvisionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchInstance

> AppInstanceExtended PatchInstance(ctx).AppInstancePatch(appInstancePatch).Filter(filter).Execute()

Update fields of an application instance.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	appInstancePatch := *openapiclient.NewAppInstancePatch() // AppInstancePatch | Fields to be updated in the application instance record.
	filter := "filter_example" // string | The filter query parameter is used to limit the set of resources to be updated in PATCH call. The returned set of resources must match the criteria in the filter.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | applicationInstanceId                      |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnifiedapioasAPI.PatchInstance(context.Background()).AppInstancePatch(appInstancePatch).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.PatchInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchInstance`: AppInstanceExtended
	fmt.Fprintf(os.Stdout, "Response from `UnifiedapioasAPI.PatchInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInstancePatch** | [**AppInstancePatch**](AppInstancePatch.md) | Fields to be updated in the application instance record. | 
 **filter** | **string** | The filter query parameter is used to limit the set of resources to be updated in PATCH call. The returned set of resources must match the criteria in the filter.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | applicationInstanceId                      |  | 

### Return type

[**AppInstanceExtended**](AppInstanceExtended.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInstance

> RemoveInstance(ctx).Filter(filter).Execute()

Remove application instances.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	filter := "filter_example" // string | The filter query parameter is used to limit the set of resources to be deleted in a collection-level DELETE.   | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | applicationInstanceId, applicationId, region | | Logic      | and                                        |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UnifiedapioasAPI.RemoveInstance(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.RemoveInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter query parameter is used to limit the set of resources to be deleted in a collection-level DELETE.   | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | applicationInstanceId, applicationId, region | | Logic      | and                                        |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProvision

> RemoveProvision(ctx).Filter(filter).Execute()

Remove customer provisions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	filter := "filter_example" // string | The filter query parameter is used to limit the set of resources to be deleted in a collection-level DELETE.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | workspaceId, applicationInstanceId, applicationId, region | | Logic      | and                                        |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UnifiedapioasAPI.RemoveProvision(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.RemoveProvision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProvisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter query parameter is used to limit the set of resources to be deleted in a collection-level DELETE.  | CLASS      |  EXAMPLES                                  | |------------|--------------------------------------------| | Type       | string                                     | | Operations | eq                                         | | Criteria   | workspaceId, applicationInstanceId, applicationId, region | | Logic      | and                                        |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProvisionOnID

> RemoveProvisionOnID(ctx, id).Execute()

Remove customer provision by ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/guptavip20/https://github.com/guptavip20/unified-api-sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer Provision Record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UnifiedapioasAPI.RemoveProvisionOnID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnifiedapioasAPI.RemoveProvisionOnID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer Provision Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProvisionOnIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

