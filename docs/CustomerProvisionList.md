# CustomerProvisionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CustomerProvisionExtended**](CustomerProvisionExtended.md) |  | 
**Count** | **int32** |  | 
**Next** | **string** |  | 

## Methods

### NewCustomerProvisionList

`func NewCustomerProvisionList(items []CustomerProvisionExtended, count int32, next string, ) *CustomerProvisionList`

NewCustomerProvisionList instantiates a new CustomerProvisionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerProvisionListWithDefaults

`func NewCustomerProvisionListWithDefaults() *CustomerProvisionList`

NewCustomerProvisionListWithDefaults instantiates a new CustomerProvisionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CustomerProvisionList) GetItems() []CustomerProvisionExtended`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomerProvisionList) GetItemsOk() (*[]CustomerProvisionExtended, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomerProvisionList) SetItems(v []CustomerProvisionExtended)`

SetItems sets Items field to given value.


### GetCount

`func (o *CustomerProvisionList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CustomerProvisionList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CustomerProvisionList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *CustomerProvisionList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CustomerProvisionList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CustomerProvisionList) SetNext(v string)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


