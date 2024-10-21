# AppInstanceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppInstanceExtended**](AppInstanceExtended.md) |  | 
**Count** | **int32** |  | 
**Next** | **string** |  | 

## Methods

### NewAppInstanceList

`func NewAppInstanceList(items []AppInstanceExtended, count int32, next string, ) *AppInstanceList`

NewAppInstanceList instantiates a new AppInstanceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstanceListWithDefaults

`func NewAppInstanceListWithDefaults() *AppInstanceList`

NewAppInstanceListWithDefaults instantiates a new AppInstanceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AppInstanceList) GetItems() []AppInstanceExtended`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AppInstanceList) GetItemsOk() (*[]AppInstanceExtended, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AppInstanceList) SetItems(v []AppInstanceExtended)`

SetItems sets Items field to given value.


### GetCount

`func (o *AppInstanceList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AppInstanceList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AppInstanceList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *AppInstanceList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AppInstanceList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AppInstanceList) SetNext(v string)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


