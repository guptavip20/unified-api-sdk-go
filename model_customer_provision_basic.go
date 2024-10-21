/*
GLCP Unified API Routing Services API

Allow registration of information needed for routing GLCP APIs. 

API version: v1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unifiedsdkgo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CustomerProvisionBasic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerProvisionBasic{}

// CustomerProvisionBasic struct for CustomerProvisionBasic
type CustomerProvisionBasic struct {
	ApplicationInstanceId string `json:"applicationInstanceId"`
	Region string `json:"region"`
	ApplicationId string `json:"applicationId"`
	WorkspaceId string `json:"workspaceId"`
}

type _CustomerProvisionBasic CustomerProvisionBasic

// NewCustomerProvisionBasic instantiates a new CustomerProvisionBasic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerProvisionBasic(applicationInstanceId string, region string, applicationId string, workspaceId string) *CustomerProvisionBasic {
	this := CustomerProvisionBasic{}
	this.ApplicationInstanceId = applicationInstanceId
	this.Region = region
	this.ApplicationId = applicationId
	this.WorkspaceId = workspaceId
	return &this
}

// NewCustomerProvisionBasicWithDefaults instantiates a new CustomerProvisionBasic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerProvisionBasicWithDefaults() *CustomerProvisionBasic {
	this := CustomerProvisionBasic{}
	return &this
}

// GetApplicationInstanceId returns the ApplicationInstanceId field value
func (o *CustomerProvisionBasic) GetApplicationInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationInstanceId
}

// GetApplicationInstanceIdOk returns a tuple with the ApplicationInstanceId field value
// and a boolean to check if the value has been set.
func (o *CustomerProvisionBasic) GetApplicationInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationInstanceId, true
}

// SetApplicationInstanceId sets field value
func (o *CustomerProvisionBasic) SetApplicationInstanceId(v string) {
	o.ApplicationInstanceId = v
}

// GetRegion returns the Region field value
func (o *CustomerProvisionBasic) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CustomerProvisionBasic) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CustomerProvisionBasic) SetRegion(v string) {
	o.Region = v
}

// GetApplicationId returns the ApplicationId field value
func (o *CustomerProvisionBasic) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *CustomerProvisionBasic) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *CustomerProvisionBasic) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *CustomerProvisionBasic) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *CustomerProvisionBasic) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *CustomerProvisionBasic) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

func (o CustomerProvisionBasic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerProvisionBasic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicationInstanceId"] = o.ApplicationInstanceId
	toSerialize["region"] = o.Region
	toSerialize["applicationId"] = o.ApplicationId
	toSerialize["workspaceId"] = o.WorkspaceId
	return toSerialize, nil
}

func (o *CustomerProvisionBasic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicationInstanceId",
		"region",
		"applicationId",
		"workspaceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCustomerProvisionBasic := _CustomerProvisionBasic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerProvisionBasic)

	if err != nil {
		return err
	}

	*o = CustomerProvisionBasic(varCustomerProvisionBasic)

	return err
}

type NullableCustomerProvisionBasic struct {
	value *CustomerProvisionBasic
	isSet bool
}

func (v NullableCustomerProvisionBasic) Get() *CustomerProvisionBasic {
	return v.value
}

func (v *NullableCustomerProvisionBasic) Set(val *CustomerProvisionBasic) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerProvisionBasic) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerProvisionBasic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerProvisionBasic(val *CustomerProvisionBasic) *NullableCustomerProvisionBasic {
	return &NullableCustomerProvisionBasic{value: val, isSet: true}
}

func (v NullableCustomerProvisionBasic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerProvisionBasic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

