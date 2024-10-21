# ApiGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ApiGroupExtended**](ApiGroupExtended.md) |  | 
**Count** | **int32** |  | 
**Next** | **string** |  | 

## Methods

### NewApiGroupList

`func NewApiGroupList(items []ApiGroupExtended, count int32, next string, ) *ApiGroupList`

NewApiGroupList instantiates a new ApiGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiGroupListWithDefaults

`func NewApiGroupListWithDefaults() *ApiGroupList`

NewApiGroupListWithDefaults instantiates a new ApiGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ApiGroupList) GetItems() []ApiGroupExtended`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApiGroupList) GetItemsOk() (*[]ApiGroupExtended, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApiGroupList) SetItems(v []ApiGroupExtended)`

SetItems sets Items field to given value.


### GetCount

`func (o *ApiGroupList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApiGroupList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApiGroupList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *ApiGroupList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ApiGroupList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ApiGroupList) SetNext(v string)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


